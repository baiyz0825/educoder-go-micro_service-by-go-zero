package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserEarnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserEarnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserEarnLogic {
	return &DelUserEarnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserEarnLogic) DelUserEarn(in *pb.DelUserEarnReq) (*pb.DelUserEarnResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check id
	if in.GetId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	earn := l.svcCtx.Query.UserEarn
	_, err := earn.WithContext(l.ctx).Where(earn.ID.Eq(in.GetId())).Delete()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
	}
	return &pb.DelUserEarnResp{}, nil
}
