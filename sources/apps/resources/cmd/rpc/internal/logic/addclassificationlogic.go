package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClassificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassificationLogic {
	return &AddClassificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------资源分类信息-----------------------

func (l *AddClassificationLogic) AddClassification(in *pb.AddClassificationReq) (*pb.AddClassificationResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if len(in.GetClassName()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetClassParentID() == 0 {
		in.ClassParentID = 0
	}
	// insert db
	classification := &model.Classification{
		ClassParentID:    in.GetClassParentID(),
		ClassName:        &in.ClassName,
		ClassResourceNum: &in.ClassResourceNum,
	}
	err := l.svcCtx.Query.Classification.WithContext(l.ctx).Create(classification)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	return &pb.AddClassificationResp{}, nil
}
