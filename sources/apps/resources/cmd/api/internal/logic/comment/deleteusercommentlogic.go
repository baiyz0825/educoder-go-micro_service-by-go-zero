package comment

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserCommentLogic {
	return &DeleteUserCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteUserComment
//
//	@Description: 删除用户评论，使用评论id删除
//	@receiver l
//	@param req
//	@return error
func (l *DeleteUserCommentLogic) DeleteUserComment(req *types.DelCommentReq) error {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// rpc
	commentReq := &pb.DelResCommentReq{ID: req.ID}
	_, err := l.svcCtx.ResourcesRpc.DelResComment(deadlineCtx, commentReq)
	if err != nil {
		return err
	}
	return nil
}
