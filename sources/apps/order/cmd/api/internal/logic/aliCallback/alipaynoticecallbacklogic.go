package aliCallback

import (
	"context"
	"strconv"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

type AlipayNoticeCallBackLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	Notification *alipay.TradeNotification
}

func NewAlipayNoticeCallBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlipayNoticeCallBackLogic {
	return &AlipayNoticeCallBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AlipayNoticeCallBack
//
//	@Description: 支付宝订单消息回调接口
//	@receiver l
//	@return error
func (l *AlipayNoticeCallBackLogic) AlipayNoticeCallBack() error {
	if l.Notification == nil {
		return xerr.NewErrMsg("未收到支付宝回调")
	}
	// 分发消息类别
	switch l.Notification.TradeStatus {
	case xconst.ALIPAY_WAIT_BUYER_PAY:
		{
			l.Logger.WithFields(logx.Field("订单号:", l.Notification.OutTradeNo),
				logx.Field("订单价格:", l.Notification.TotalAmount)).
				Info("等待买家支付!")
		}
	case xconst.AlIPAY_TRADE_CLOSED:
		{
			l.Logger.WithFields(logx.Field("订单号:", l.Notification.OutTradeNo),
				logx.Field("订单价格:", l.Notification.TotalAmount)).
				Info("交易关闭!")
			// 删除订单信息
			uuid, err := strconv.ParseInt(l.Notification.OutTradeNo, 10, 64)
			if err != nil {
				l.Logger.WithFields(logx.Field("err:", err)).Error("转化消息类型失败")
				return xerr.NewErrCode(xerr.SERVER_ERROR)
			}
			_, err = l.svcCtx.OrderRpc.UpdateOrderStatus(context.Background(), &pb.UpdateOrderStatusReq{
				NeedDelete: true,
				Uuid:       uuid,
			})
			if err != nil {
				l.Logger.WithFields(logx.Field("err:", err)).Error("远程删除订单失败")
				return xerr.NewErrCode(xerr.SERVER_ERROR)
			}
		}
	case xconst.ALIPAY_TRADE_SUCCESS:
		{
			// 订单交易成功，更新订单状态
			uuid, err := strconv.ParseInt(l.Notification.OutTradeNo, 10, 64)
			if err != nil {
				l.Logger.WithFields(logx.Field("err:", err)).Error("转化消息类型失败")
				return xerr.NewErrCode(xerr.SERVER_ERROR)
			}
			_, err = l.svcCtx.OrderRpc.UpdateOrderStatus(context.Background(), &pb.UpdateOrderStatusReq{
				Status:     xconst.ORDER_STATUS_PAYED,
				NeedDelete: false,
				Uuid:       uuid,
			})
			if err != nil {
				l.Logger.WithFields(logx.Field("err:", err)).Error("远程删除订单失败")
				return xerr.NewErrCode(xerr.SERVER_ERROR)
			}
		}
	// 不可退款
	case xconst.ALIPAY_TRADE_FINISHED:
		{
			// 订单交易成功，不可退款
			uuid, err := strconv.ParseInt(l.Notification.OutTradeNo, 10, 64)
			if err != nil {
				l.Logger.WithFields(logx.Field("err:", err)).Error("转化消息类型失败")
				return xerr.NewErrCode(xerr.SERVER_ERROR)
			}
			_, err = l.svcCtx.OrderRpc.UpdateOrderStatus(context.Background(), &pb.UpdateOrderStatusReq{
				Status:     xconst.ORDER_STATUS_PAYED,
				NeedDelete: false,
				Uuid:       uuid,
			})
			if err != nil {
				l.Logger.WithFields(logx.Field("err:", err)).Error("远程删除订单失败")
				return xerr.NewErrCode(xerr.SERVER_ERROR)
			}
		}
	default:
		l.Logger.WithFields(logx.Field("消息内容:", l.Notification)).Info("未知回调消息类别")
		return nil

	}

	return nil
}
