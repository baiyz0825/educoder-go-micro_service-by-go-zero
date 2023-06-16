package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserEarnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserEarnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserEarnLogic {
	return &AddUserEarnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户收入支出统计-----------------------

func (l *AddUserEarnLogic) AddUserEarn(in *pb.AddUserEarnReq) (*pb.AddUserEarnResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetUserId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// create model
	earn := &model.UserEarn{
		UserID:  in.GetUserId(),
		EarnNum: in.GetEarnNum(),
		PayNum:  in.GetPayNum(),
	}
	err := l.svcCtx.Query.UserEarn.WithContext(l.ctx).Create(earn)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	// return
	return &pb.AddUserEarnResp{}, nil
}
