package logic

import (
	"context"
	"strconv"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/pkg/errors"
	"github.com/smartwalle/alipay/v3"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAilPayStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAilPayStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAilPayStatusLogic {
	return &CheckAilPayStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAilPayStatusLogic) CheckAilPayStatus(in *pb.CheckAilPayStatusReq) (*pb.CheckAilPayStatusResp, error) {
	var p = alipay.TradeQuery{}
	p.OutTradeNo = strconv.FormatInt(in.OrderUuid, 10)
	alipayStatus, err := l.svcCtx.AliPayClient.TradeQuery(p)
	if err != nil {
		return nil, errors.Wrap(err, "支付宝接口请求失败！")
	}
	return &pb.CheckAilPayStatusResp{
		AliPayStatus: string(alipayStatus.Content.TradeStatus),
		Status:       1,
	}, nil
}
