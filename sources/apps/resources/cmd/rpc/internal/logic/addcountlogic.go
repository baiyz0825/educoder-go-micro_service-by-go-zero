package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCountLogic {
	return &AddCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户上传资源量统计信息-----------------------

func (l *AddCountLogic) AddCount(in *pb.AddCountReq) (*pb.AddCountResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetFileNum() < 0 || in.GetVideoNum() < 0 || in.GetPicNum() < 0 || in.GetStorageSize() < 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// insert db
	count := &model.Count{
		UID:         in.GetUID(),
		FileNum:     in.GetFileNum(),
		VideoNum:    in.GetVideoNum(),
		PicNum:      in.GetPicNum(),
		StorageSize: in.GetStorageSize(),
	}
	err := l.svcCtx.Query.Count.WithContext(l.ctx).Create(count)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	return &pb.AddCountResp{}, nil
}
