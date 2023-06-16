package comment

import (
	"context"
	"encoding/json"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserCommentLogic {
	return &AddUserCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddUserComment
//
//	@Description: 增加文件一个用户资源评论
//	@receiver l
//	@param req
//	@return error
func (l *AddUserCommentLogic) AddUserComment(req *types.AddResCommentReq) error {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	if err != nil {
		return xerr.NewErrMsg("获取用户信息失败")
	}
	// rpc
	commentReq := &pb.AddResCommentReq{
		Owner:      uid,
		ResourceID: req.ResourceID,
		Content:    req.Content,
	}
	_, err = l.svcCtx.ResourcesRpc.AddResComment(deadlineCtx, commentReq)
	if err != nil {
		return err
	}
	return nil
}
