package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateClassificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassificationLogic {
	return &UpdateClassificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateClassificationLogic) UpdateClassification(in *pb.UpdateClassificationReq) (*pb.UpdateClassificationResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetClassID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// insert db
	classification := model.Classification{
		ClassParentID:    in.GetClassParentID(),
		ClassName:        &in.ClassName,
		ClassResourceNum: &in.ClassResourceNum,
	}
	q := l.svcCtx.Query.Classification
	_, err := q.WithContext(l.ctx).Where(q.ClassID.Eq(in.GetClassID())).Updates(classification)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}

	return &pb.UpdateClassificationResp{}, nil
}
