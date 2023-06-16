package order

import (
	"context"
	"fmt"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	resPb "github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	tradePb "github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	userPb "github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoLogic {
	return &GetOrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetOrderInfo
//
//	@Description: 使用订单uuid和用户id查询订单详情信息
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GetOrderInfoLogic) GetOrderInfo(req *types.OrderInfoReq) (resp *types.OrderInfoResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	// 查询订单
	deadline, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	orderData, err := l.svcCtx.OrderRpc.GetOrderInfoByUUIDAndUserId(deadline, &pb.GetOrderInfoByUUIDAndUserIDReq{
		UserId: req.UserId,
		Uuid:   req.Uuid,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("订单查询失败，请稍后再试")
	}
	resp = &types.OrderInfoResp{}
	err = copier.Copy(resp, orderData.Order)
	resp.StatusUpdateTime = orderData.Order.UpdateTime
	if err != nil {
		l.Logger.WithFields(
			logx.Field("err:", err),
			logx.Field("详情：", fmt.Sprintf("用户id：%v,订单编号uuid：%v", req.UserId, req.Uuid))).
			Error("用户订单查询失败")
		return nil, xerr.NewErrMsg("订单数据查询失败，系统错误")
	}
	// 查询商品详情
	productData, err := l.svcCtx.TradeRpc.GetProductById(deadline, &tradePb.GetProductByIdReq{ID: orderData.Order.ProductId})
	if err != nil {
		l.Logger.WithFields(
			logx.Field("err:", err),
			logx.Field("详情：", fmt.Sprintf("用户id：%v,订单编号uuid：%v,商品id:%v", req.UserId, req.Uuid, orderData.Order.ProductId))).
			Error("用户订单中商品信息查询失败")
		return nil, xerr.NewErrMsg("订单数据查询失败，系统错误")
	}
	// 查询商品所属资源信息
	fileProductData, err := l.svcCtx.ResourcesRpc.GetFileById(deadline, &resPb.GetFileByIdReq{ID: productData.Product.ProductBind})
	if err != nil {
		l.Logger.WithFields(
			logx.Field("err:", err),
			logx.Field("详情：", fmt.Sprintf("用户id：%v,订单编号uuid：%v,商品id:%v", req.UserId, req.Uuid, orderData.Order.ProductId))).
			Error("用户订单中商品绑定的资源信息详细查询失败")
		return nil, xerr.NewErrMsg("订单数据查询失败，系统错误")
	}
	resp.Product.ID = productData.Product.ID
	resp.Product.Name = productData.Product.Name
	// 查询商品所属用户信息
	userData, err := l.svcCtx.UserRpc.GetUserById(deadline, &userPb.GetUserByIdReq{ID: fileProductData.File.Owner})
	if err != nil {
		l.Logger.WithFields(
			logx.Field("err:", err),
			logx.Field("详情：", fmt.Sprintf(
				"用户id：%v,订单编号uuid：%v,商品id:%v,资源id:%v",
				req.UserId, req.Uuid, orderData.Order.ProductId, fileProductData.File.Owner))).
			Error("用户订单中商品对应资源的用户所属人信息详细查询失败")
		return nil, xerr.NewErrMsg("订单数据查询失败，系统错误")
	}
	if userData != nil {
		resp.Product.ProductOwnerName = userData.User.Name
		resp.Product.ProductOwnerId = userData.User.UID
	}
	return resp, nil
}
