package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCountLogic {
	return &UpdateCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCountLogic) UpdateCount(in *pb.UpdateCountReq) (*pb.UpdateCountResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetFileNum() < 0 || in.GetVideoNum() < 0 || in.GetPicNum() < 0 || in.GetStorageSize() < 0 || in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// update db
	count := model.Count{
		FileNum:     in.GetFileNum(),
		VideoNum:    in.GetVideoNum(),
		PicNum:      in.GetPicNum(),
		StorageSize: in.GetStorageSize(),
	}
	q := l.svcCtx.Query.Count
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.GetID())).Updates(count)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}

	return &pb.UpdateCountResp{}, nil
}
