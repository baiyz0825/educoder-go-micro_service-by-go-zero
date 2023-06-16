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
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsConditionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentsConditionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsConditionLogic {
	return &GetCommentsConditionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetCommentsCondition
//
//	@Description: 条件查询评论（用户Id或者资源id）
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GetCommentsConditionLogic) GetCommentsCondition(req *types.ResCommentByUserOrResIdReq) (resp *types.ResCommentByUserOrResIdResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.SERVER_ERROR)
	}
	// rpc
	pbReq := &pb.SearchResCommentByUserOrResIdReq{
		Page:       req.Page,
		Limit:      req.Limit,
		Owner:      uid,
		ResourceID: req.ResourceID,
	}
	data, err := l.svcCtx.ResourcesRpc.SearchResCommentByUserOrResId(deadlineCtx, pbReq)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.SERVER_ERROR)
	}
	resp = &types.ResCommentByUserOrResIdResp{}
	err = copier.Copy(&resp.Comments, data.GetResComment())
	if err != nil {
		l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
		return nil, xerr.NewErrMsg("数据处理异常")
	}
	return
}
