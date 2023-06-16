package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFileLogic {
	return &AddFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------文件资源存储表（非文本类型）-----------------------

func (l *AddFileLogic) AddFile(in *pb.AddFileReq) (*pb.AddFileResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check data
	if len(in.GetName()) == 0 || in.GetOwner() == 0 || in.GetClass() == 0 || len(in.GetSuffix()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// gen uuid && db data
	id, err := utils.GenSnowFlakeId()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("生成系统id错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// insert db
	file := &model.File{
		UUID:          id,
		Name:          in.GetName(),
		ObfuscateName: utils.RandString(0),
		Size:          in.GetSize(),
		Owner:         in.GetOwner(),
		Status:        &in.Status,
		Type:          in.GetType(),
		Class:         in.GetClass(),
		Suffix:        in.GetSuffix(),
		DownloadAllow: &in.DownloadAllow,
		Link:          &in.Link,
		FilePoster:    &in.FilePoster,
	}
	err = l.svcCtx.Query.File.WithContext(l.ctx).Create(file)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	return &pb.AddFileResp{}, nil
}
