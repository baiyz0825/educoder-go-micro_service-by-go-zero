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

type GetClassificationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassificationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassificationByIdLogic {
	return &GetClassificationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClassificationByIdLogic) GetClassificationById(in *pb.GetClassificationByIdReq) (*pb.GetClassificationByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// search
	classification := l.svcCtx.Query.Classification
	data, err := classification.WithContext(l.ctx).Where(classification.ClassID.Eq(in.GetID())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetClassificationByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 复制数据
	classificationData := &pb.Classification{}
	err = copier.Copy(classificationData, data)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	classificationData.UpdateTime = data.UpdateTime.UnixMilli()
	classificationData.CreateTime = data.CreateTime.UnixMilli()
	return &pb.GetClassificationByIdResp{Classification: classificationData}, nil
}
