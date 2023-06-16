package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *pb.DelUserReq) (*pb.DelUserResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// delete db
	user := l.svcCtx.Query.User
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	if _, err := user.WithContext(ctx).Where(user.UID.Eq(in.ID)).Delete(); err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_DELETE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
	}
	return &pb.DelUserResp{}, nil
}
