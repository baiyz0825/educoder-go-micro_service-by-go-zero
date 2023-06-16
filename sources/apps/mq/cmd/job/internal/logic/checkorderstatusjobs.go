package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/bo"
	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckOrderStatusJob struct {
	logx.Logger
	svcCtx *svc.ServiceContext
}

func NewCheckOrderStatusJobLogic(svcCtx *svc.ServiceContext) *CheckOrderStatusJob {
	return &CheckOrderStatusJob{
		Logger: logx.WithContext(context.Background()),
		svcCtx: svcCtx,
	}
}

// ProcessTask
//
//	@Description: 定时轮询支付状态
//	@receiver l
//	@param ctx
//	@param t
//	@return error
func (l *CheckOrderStatusJob) ProcessTask(ctx context.Context, t *asynq.Task) error {
	l.Logger.Info("开始处理轮询支付状态任务:")
	var orderJobDec bo.OrderMqStruct
	l.Logger.Info(fmt.Sprintf("[1]轮询支付状态:开始处理 %v", t.Payload()))
	err := json.Unmarshal(t.Payload(), &orderJobDec)
	deadline, cancelFunc := context.WithDeadline(ctx, utils.GetContextDefaultTime())
	defer cancelFunc()
	// 查询订单状态
	orderStatusData, err := l.svcCtx.OrderRpc.GetOrderStatusByUUID(deadline, &pb.GetOrderStatusByUUIDReq{OrderUUid: orderJobDec.Uuid})
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).Error("[2]轮询支付状态:定时更新订单状态失败")
		return err
	}
	// 订单不存在
	if orderStatusData.IsDelete {
		l.Logger.WithFields(logx.Field("orderId:", orderJobDec.Uuid)).Info("[2]轮询支付状态:订单已被过期删除！不需要再次查询订单状态")
		return nil
	}
	// 循环更新订单状态
	for {
		// 查询支付宝订单状态
		status, err := l.svcCtx.OrderRpc.CheckAilPayStatus(context.Background(), &pb.CheckAilPayStatusReq{OrderUuid: orderJobDec.Uuid})
		if err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).Error("[3]轮询支付状态:支付宝订单状态检查失败！")
			return err
		}
		if status.AliPayStatus == xconst.ALIPAY_TRADE_NOT_EXIST {
			l.Logger.WithFields(logx.Field("order:", orderJobDec.Uuid)).Info("[3]轮询支付状态:支付宝订单未创建，稍后重试！")
			// 休眠重试
			time.Sleep(20 * time.Second)
			continue
		}
		// 需要更新情况
		orderUpdateData := &pb.UpdateOrderStatusReq{
			Uuid: orderJobDec.Uuid,
		}
		if status.AliPayStatus == xconst.AlIPAY_TRADE_CLOSED {
			l.Logger.WithFields(logx.Field("order:", orderJobDec.Uuid)).Info("[3]轮询支付状态:支付宝订单已取消或者删除，正在更新订单状态！")
			// 删除订单
			orderUpdateData.NeedDelete = true
		}
		if status.AliPayStatus == xconst.ALIPAY_TRADE_SUCCESS || status.AliPayStatus == xconst.ALIPAY_TRADE_FINISHED {
			l.Logger.WithFields(logx.Field("order:", orderJobDec.Uuid)).Info("[3]轮询支付状态:支付宝订单完成，正在更新订单状态！")
			// 更新订单状态
			orderUpdateData.NeedDelete = false
			orderUpdateData.Status = xconst.ORDER_STATUS_PAYED
		}
		// 更新
		_, err = l.svcCtx.OrderRpc.UpdateOrderStatus(context.Background(), orderUpdateData)
		if err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).Error("[3.1]轮询支付状态:本次更新订单数据失败，稍后重试")
			return err
		}
		break
	}
	l.Logger.WithFields(logx.Field("订单uuid:", orderJobDec.Uuid)).Info("[4]轮询支付状态:更新订单状态成功！")
	return nil
}
