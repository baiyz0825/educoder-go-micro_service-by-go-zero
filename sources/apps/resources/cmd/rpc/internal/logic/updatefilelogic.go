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

type UpdateFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileLogic {
	return &UpdateFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFileLogic) UpdateFile(in *pb.UpdateFileReq) (*pb.UpdateFileResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check data
	if len(in.GetName()) == 0 || in.GetOwner() == 0 || in.GetClass() == 0 || len(in.GetSuffix()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// insert db
	file := &model.File{
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
	q := l.svcCtx.Query.File
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.GetID())).Updates(file)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}

	return &pb.UpdateFileResp{}, nil
}
