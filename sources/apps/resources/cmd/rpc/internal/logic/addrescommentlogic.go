package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddResCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddResCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddResCommentLogic {
	return &AddResCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------资源评论信息-----------------------

func (l *AddResCommentLogic) AddResComment(in *pb.AddResCommentReq) (*pb.AddResCommentResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check param
	if in.GetOwner() == 0 || in.GetResourceID() == 0 || len(in.GetContent()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// dirty work check
	// create db
	comment := &model.ResComment{
		Owner:      in.GetOwner(),
		ResourceID: in.GetResourceID(),
		Content:    &in.Content,
	}
	err := l.svcCtx.Query.ResComment.WithContext(l.ctx).Create(comment)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}

	return &pb.AddResCommentResp{}, nil
}
