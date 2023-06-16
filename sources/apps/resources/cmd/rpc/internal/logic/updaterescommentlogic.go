package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateResCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateResCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateResCommentLogic {
	return &UpdateResCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateResCommentLogic) UpdateResComment(in *pb.UpdateResCommentReq) (*pb.UpdateResCommentResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check param
	if in.GetID() == 0 || in.GetOwner() == 0 || in.GetResourceID() == 0 || len(in.GetContent()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// dirty work check
	// update db
	comment := &model.ResComment{
		Content: &in.Content,
	}
	resComment := l.svcCtx.Query.ResComment
	_, err := resComment.WithContext(l.ctx).Where(resComment.ID.Eq(in.GetID())).Updates(comment)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}
	return &pb.UpdateResCommentResp{}, nil
}
