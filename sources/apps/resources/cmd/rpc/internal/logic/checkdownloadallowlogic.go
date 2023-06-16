package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckDownloadAllowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckDownloadAllowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckDownloadAllowLogic {
	return &CheckDownloadAllowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckDownloadAllow 检查文件是否允许下载
func (l *CheckDownloadAllowLogic) CheckDownloadAllow(in *pb.CheckDownloadAllowReq) (*pb.CheckDownloadAllowResp, error) {
	file := l.svcCtx.Query.File
	deadline, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	find, err := file.WithContext(deadline).Select(file.ID, file.DownloadAllow).Where(file.ID.Eq(in.FileId)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return &pb.CheckDownloadAllowResp{}, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// rpc return
	return &pb.CheckDownloadAllowResp{
		IsAllow: utils.If(*(find.DownloadAllow) == 1, true, false).(bool),
	}, nil
}
