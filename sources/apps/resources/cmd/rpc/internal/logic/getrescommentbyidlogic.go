package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetResCommentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetResCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetResCommentByIdLogic {
	return &GetResCommentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetResCommentByIdLogic) GetResCommentById(in *pb.GetResCommentByIdReq) (*pb.GetResCommentByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// delete
	resComment := l.svcCtx.Query.ResComment
	data, err := resComment.WithContext(l.ctx).Where(resComment.ID.Eq(in.GetID())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetResCommentByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 复制数据
	resCommentData := &pb.ResComment{}
	err = copier.Copy(resCommentData, data)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	resCommentData.CreateTime = data.CreateTime.UnixMilli()
	resCommentData.UpdateTime = data.UpdateTime.UnixMilli()
	return &pb.GetResCommentByIdResp{
		ResComment: resCommentData,
	}, nil
}
