package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCountLogic {
	return &DelCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCountLogic) DelCount(in *pb.DelCountReq) (*pb.DelCountResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// delete
	count := l.svcCtx.Query.Count
	_, err := count.WithContext(l.ctx).Where(count.ID.Eq(in.GetId())).Delete()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_DELETE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
	}
	return &pb.DelCountResp{}, nil
}
