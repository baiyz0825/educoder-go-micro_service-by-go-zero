package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOrderLogic {
	return &DelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOrderLogic) DelOrder(in *pb.DelOrderReq) (*pb.DelOrderResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check id
	if in.GetId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	order := l.svcCtx.Query.Order
	_, err := order.WithContext(l.ctx).Where(order.ID.Eq(in.GetId())).Delete()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
	}
	return &pb.DelOrderResp{}, nil
}
