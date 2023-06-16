package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelResCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelResCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelResCommentLogic {
	return &DelResCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelResCommentLogic) DelResComment(in *pb.DelResCommentReq) (*pb.DelResCommentResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// delete
	comment := l.svcCtx.Query.ResComment
	_, err := comment.WithContext(l.ctx).Where(comment.ID.Eq(in.GetID())).Delete()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_DELETE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
	}

	return &pb.DelResCommentResp{}, nil
}
