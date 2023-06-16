package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/pkg/errors"
	"github.com/smartwalle/alipay/v3"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderAliAndDbLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderAliAndDbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderAliAndDbLogic {
	return &DeleteOrderAliAndDbLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrderAliAndDb
//
//	@Description: 删除支付宝和数据库订单信息
//	@receiver l
//	@param in
//	@return *pb.DeleteOrderAliAndDbResp
//	@return error
func (l *DeleteOrderAliAndDbLogic) DeleteOrderAliAndDb(in *pb.DeleteOrderAliAndDbReq) (*pb.DeleteOrderAliAndDbResp, error) {
	// 检查参数
	if in.GetUuid() == 0 || in.GetUserId() == 0 || len(in.GetPayPathOrderNum()) <= 0 {
		return &pb.DeleteOrderAliAndDbResp{Status: false}, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 删除支付宝订单
	rsp := alipay.TradeClose{}
	rsp.OutTradeNo = in.PayPathOrderNum
	l.Logger.Debugf("待取消的支付宝订单是：%v", rsp)
	tradeClose, err := l.svcCtx.AliPayClient.TradeClose(rsp)
	if err != nil {
		return nil, errors.Wrap(err, "支付宝撤销订单失败")
	}
	// 如果失败，但是不是交易不存在（订单未创建 https://opendocs.alipay.com/support/01rax8 ），直接删除数据库侧
	if !tradeClose.Content.Code.IsSuccess() && tradeClose.Content.SubCode != xconst.ALIPAY_TRADE_NOT_EXIST {
		l.Logger.WithFields(logx.Field("msg:", tradeClose.Content)).Error("删除支付宝订单失败")
		return nil, errors.Wrap(err, "支付宝撤销订单失败:"+tradeClose.Content.Msg)
	}
	// 删除数据库订单
	orderQ := l.svcCtx.Query.Order
	_, err = orderQ.WithContext(context.Background()).Where(orderQ.UUID.Eq(in.GetUuid()), orderQ.PayPathOrderNum.Eq(in.PayPathOrderNum)).Delete()
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).Error("数据库撤销订单失败:" + tradeClose.Content.Msg)
		return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
	}
	return &pb.DeleteOrderAliAndDbResp{Status: true}, nil
}
