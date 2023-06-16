package order

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	resPb "github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	tradePb "github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	userPb "github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAllOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserAllOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAllOrderLogic {
	return &GetUserAllOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUserAllOrder
//
//	@Description: 条件获取用户订单数据
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GetUserAllOrderLogic) GetUserAllOrder(req *types.GetUserAllOrder) (resp *types.GetUserAllOrderResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	timeout, cancelFunc := context.WithTimeout(context.Background(), utils.GetContextDuration())
	defer cancelFunc()
	uid, _ := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	conditionReq := &pb.SearchOrderByConditionReq{
		Page:     req.Page,
		Limit:    req.Limit,
		SysModel: xconst.ORDER_SYSTEM_MODE_TRADE, // 默认
		Status:   req.Status,
		UserId:   uid,
		PayPath:  req.PayPath,
	}
	ordersData, err := l.svcCtx.OrderRpc.SearchOrderByCondition(timeout, conditionReq)
	if err != nil {
		return nil, err
	}
	resp = &types.GetUserAllOrderResp{Total: 0}
	if len(ordersData.GetOrder()) == 0 {
		return resp, nil
	}
	// 过滤订单的绑定的商品id
	// 使用product bind id查询商品对应的资源绑定的用户信息（owner name 、 id）（包含软删除 isDelete = true）
	// 重新map 绑定资源id
	// 多线查询，设置任务组
	waitG := sync.WaitGroup{}
	// 线程队列
	// 初始化返回数据
	var respBindData []types.OrderInfo
	// 带缓冲数据的通道
	cacheRespDataChan := make(chan types.OrderInfo, 100)
	goPool := make(chan struct{}, 50)
	// 循环所有订单
	// TODO 对于处理订单数据过程中出现错误，未进行处理（可以采用channel进行错误监听）
	for _, order := range ordersData.Order {
		// 获得 goroutine。
		goPool <- struct{}{}
		waitG.Add(1)
		go func(data *pb.Order) {
			// 处理业务
			// 查询 订单绑定的商品
			product, err := l.svcCtx.TradeRpc.GetProductById(l.ctx, &tradePb.GetProductByIdReq{ID: data.ProductId})
			if err != nil || product.Product == nil {
				l.Logger.WithFields(logx.Field("err:", err)).Errorf("处理获取订单匹配的商品数据出错，该获取的数据是:%v", data)
				return
			}
			// 获取资源数据
			udiData, err := l.svcCtx.ResourcesRpc.GetFilResourcesUSerId(context.Background(), &resPb.GetFilResourcesUSerIdReq{ResourcesId: product.Product.ProductBind})
			if err != nil || udiData == nil {
				l.Logger.WithFields(logx.Field("err:", err)).Errorf("处理获取匹配资源数据出错，该获取的数据是:%v", product)
				return
			}
			// 获取用户名称
			userData, err := l.svcCtx.UserRpc.GetUserById(context.Background(), &userPb.GetUserByIdReq{ID: udiData.UserId})
			if err != nil || userData.User == nil {
				l.Logger.WithFields(logx.Field("err:", err)).Errorf("处理获取匹配资源的用户数据出错，带查询的用户id是:%v", udiData.UserId)
				return
			}
			// 拼接数据
			dataTemp := types.OrderInfo{
				Id:   data.Id,
				Uuid: data.Uuid,
				Product: types.ProductBaseInfo{
					ID:               product.Product.ID,
					Name:             product.Product.Name,
					ProductOwnerName: userData.User.Name,
					ProductOwnerId:   userData.User.UID,
				},
				Status:           data.Status,
				UserId:           data.UserId,
				PayPrice:         data.PayPrice,
				PayPath:          data.PayPath,
				CreateTime:       data.CreateTime,
				StatusUpdateTime: data.UpdateTime,
			}
			cacheRespDataChan <- dataTemp // 将数据写入通道
			// 在使用完通道后，返回以供其他 goroutine 使用。
			defer func() {
				_ = <-goPool
			}()
			// 去除任务队列
			waitG.Done()
		}(order)
	}

	// 等待所有 goroutine 完成并关闭通道
	go func() {
		waitG.Wait()
		close(cacheRespDataChan)
		// 关闭协程任务队列
		defer close(goPool)
	}()

	// 读取通道数据
	// 阻塞，读取通道中的数据
	for orderInfo := range cacheRespDataChan {
		// 处理 cacheRespDataChan 中的数据
		respBindData = append(respBindData, orderInfo)
	}
	return &types.GetUserAllOrderResp{
		Orders: respBindData,
		Total:  ordersData.Total,
	}, nil
}
