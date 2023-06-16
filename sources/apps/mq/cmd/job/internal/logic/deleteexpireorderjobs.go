package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/bo"
	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExpireOrderJob struct {
	logx.Logger
	svcCtx *svc.ServiceContext
}

func NewDoOrderLogic(svcCtx *svc.ServiceContext) *DeleteExpireOrderJob {
	return &DeleteExpireOrderJob{
		Logger: logx.WithContext(context.Background()),
		svcCtx: svcCtx,
	}
}

// ProcessTask
//
//	@Description: 处理删除过期订单任务
//	@receiver l
//	@param ctx
//	@param t
//	@return error
func (l *DeleteExpireOrderJob) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var orderJobDec bo.OrderMqStruct
	l.Logger.Info(fmt.Sprintf("[1]<删除过期订单任务>:开始处理 %v", t.Payload()))
	err := json.Unmarshal(t.Payload(), &orderJobDec)
	if err != nil {
		return err
	}
	// 查询数据库订单状态是否未支付
	orderInfoReq := &pb.GetOrderInfoByUUIDAndUserIDReq{
		UserId: orderJobDec.UserId,
		Uuid:   orderJobDec.Uuid,
	}
	deadline, cancelFunc := context.WithDeadline(ctx, utils.GetContextDefaultTime())
	defer cancelFunc()
	l.Logger.Info(fmt.Sprintf("[2]<删除过期订单任务>:查询订单状态"))
	orderInfo, err := l.svcCtx.OrderRpc.GetOrderInfoByUUIDAndUserId(deadline, orderInfoReq)
	if err != nil || orderInfo.Order == nil || orderInfo.Order.Uuid != orderJobDec.Uuid {
		l.Logger.WithFields(logx.Field("err:", err)).Error("查询订单状态失败")
		return err
	}
	// 订单已经完成
	if orderInfo.Order.Status == xconst.ORDER_STATUS_PAYED {
		l.Logger.Info(fmt.Sprintf("[3]<删除过期订单任务>:订单已完成"))
		return nil
	}
	l.Logger.Info(fmt.Sprintf("[3]<删除过期订单任务>:订单未完成，开始删除订单"))
	// 未完成 删除订单
	deleteOrderAliAndDbReq := &pb.DeleteOrderAliAndDbReq{
		Uuid:            orderJobDec.Uuid,
		UserId:          orderJobDec.UserId,
		PayPathOrderNum: orderJobDec.PayPathOrderNum,
	}
	status, err := l.svcCtx.OrderRpc.DeleteOrderAliAndDb(deadline, deleteOrderAliAndDbReq)
	if err != nil || !status.GetStatus() {
		l.Logger.WithFields(logx.Field("err:", err)).Error("[3]<删除过期订单任务>:删除数据库以及支付宝订单失败！！")
		return err
	}
	l.Logger.Info(fmt.Sprintf("[4]<删除过期订单任务>:任务完成，成功删除订单(%v)", orderInfo.Order.Uuid))
	// 任务执行成功
	return nil
}
