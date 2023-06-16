package logic

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	resPb "github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	tradePb "github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/hibiken/asynq"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserEarnJob struct {
	logx.Logger
	svcCtx *svc.ServiceContext
}

func NewUpdateUserEarnJobLogic(svcCtx *svc.ServiceContext) *UpdateUserEarnJob {
	return &UpdateUserEarnJob{
		Logger: logx.WithContext(context.Background()),
		svcCtx: svcCtx,
	}
}

// ProcessTask
//
//	@Description: 统计用户earn任务
//	@receiver l
//	@param ctx
//	@param t
//	@return error
func (l *UpdateUserEarnJob) ProcessTask(ctx context.Context, t *asynq.Task) error {
	l.Logger.Info(fmt.Sprintf("[1] 更新用户所得任务 :开始处理 %v", t.Payload()))
	// 过滤所有订单数据，按照商品id进行分组
	data, err := l.svcCtx.OrderRpc.GetProductBindAndPrices(context.Background(), &pb.GetProductBindAndPricesReq{})
	if err != nil {
		l.Logger.Error(fmt.Sprintf("[2] 更新用户所得任务 获取订单价格以及商品绑定失败 %v", t.Payload()))
		return err
	}
	if len(data.ProductBindPrice) == 0 {
		l.Logger.Info(fmt.Sprintf("[2] 更新用户所得任务 暂时没有需要处理的订单数据 %v", t.Payload()))

	}
	// 查询分组商品Id的用户id
	l.Logger.Info(fmt.Sprintf("[2] 更新用户所得任务 分组处理订单数据 %v", t.Payload()))
	// 创建用户金额缓存
	currentMap := cmap.New[float64]()
	waitG := sync.WaitGroup{}
	routinePool := make(chan bool, 100)
	for i, bind := range data.ProductBindPrice {
		routinePool <- true
		waitG.Add(1)
		i := i
		go func(productId int64, total float64) {
			timeout, cancelFunc := context.WithTimeout(context.Background(), utils.GetContextDuration())
			defer cancelFunc()
			// 查询产品对应的资源id
			data, err := l.svcCtx.TradeRpc.GetProductBindByProductId(timeout, &tradePb.GetProductBindByProductIdReq{ProductId: productId})
			if err != nil {
				l.Logger.WithFields(logx.Field("err:", err)).Info(fmt.Sprintf("[2] 更新用户所得任务 分组处理订单数据 第%v组 执行失败", i))
				return
			}
			// 查询资源id对应的用户id
			uidData, err := l.svcCtx.ResourcesRpc.GetFilResourcesUSerId(timeout, &resPb.GetFilResourcesUSerIdReq{ResourcesId: data.ResourcesBind})
			if err != nil {
				return
			}
			// 更新map中数据
			// 并发读写 https://github.com/orcaman/concurrent-map
			currentMap.Upsert(strconv.FormatInt(uidData.UserId, 10), total, func(exist bool, old float64, newValue float64) float64 {
				if exist {
					return old + newValue
				} else {
					return newValue
				}
			})
			l.Logger.Info(fmt.Sprintf("[2] 更新用户所得任务 分组处理订单数据 临时保存处理成功！"))
			_ = <-routinePool
			waitG.Done()
		}(bind.GetProductID(), bind.GetTotal())
	}
	// 阻塞子协程处理结束
	waitG.Wait()
	// 进行更新用户数据
	l.Logger.Info(fmt.Sprintf("[3] 更新用户所得任务 开始更新用户Earn"))
	updateUserGroup := sync.WaitGroup{}
	updateUserRoutinePool := make(chan bool, 100)
	missionSize := currentMap.Count()
	var failureSize uint32 = 0
	for item := range currentMap.IterBuffered() {
		updateUserRoutinePool <- true
		uid, err := strconv.ParseInt(item.Key, 10, 64)
		if err != nil {
			return err
		}
		updateUserGroup.Add(1)
		go func(uid int64, total float64) {
			// 远程更新用户数据
			_, err := l.svcCtx.OrderRpc.UpsertUserEarn(context.Background(), &pb.AddUserEarnReq{UserId: uid, EarnNum: total})
			if err != nil {
				l.Logger.WithFields(logx.Field("err:", err)).
					Info(fmt.Sprintf("[3] 更新用户所得任务 更新用户Earn 执行失败,用户id：%v,数据是:%v", uid, total))
				atomic.AddUint32(&failureSize, 1)
				return
			}
			l.Logger.Info(fmt.Sprintf("[3] 更新用户所得任务 更新用户Earn 成功！！用户id：%v", uid))
			_ = <-updateUserRoutinePool
			updateUserGroup.Done()
		}(uid, item.Val)
	}
	updateUserGroup.Wait()
	l.Logger.Info(fmt.Sprintf("[4] 更新用户所得任务 本次任务执行成功！任务总数:%v 失败个数:%v", missionSize, failureSize))
	return nil
}
