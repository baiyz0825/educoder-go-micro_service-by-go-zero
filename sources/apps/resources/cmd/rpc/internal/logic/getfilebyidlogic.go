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

type GetFileByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileByIdLogic {
	return &GetFileByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileByIdLogic) GetFileById(in *pb.GetFileByIdReq) (*pb.GetFileByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// delete
	file := l.svcCtx.Query.File
	data, err := file.WithContext(l.ctx).Where(file.ID.Eq(in.GetID())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetFileByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 复制数据
	fileData := &pb.File{}
	err = copier.Copy(fileData, data)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	fileData.UpdateTime = data.UpdateTime.UnixMilli()
	fileData.CreateTime = data.CreateTime.UnixMilli()

	return &pb.GetFileByIdResp{File: fileData}, nil
}
