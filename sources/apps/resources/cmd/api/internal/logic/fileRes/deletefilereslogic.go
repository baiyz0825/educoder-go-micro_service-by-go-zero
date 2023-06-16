package fileRes

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileResLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileResLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileResLogic {
	return &DeleteFileResLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteFileRes
//
//	@Description: 删除评论信息
//	@receiver l
//	@param req
//	@return error
func (l *DeleteFileResLogic) DeleteFileRes(req *types.DelFileReq) error {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	_, err := l.svcCtx.ResourcesRpc.DelFile(deadlineCtx, &pb.DelFileReq{ID: req.Id})
	if err != nil {
		return err
	}
	return nil
}
