package comment

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentDetailByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentDetailByIdLogic {
	return &GetCommentDetailByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetCommentDetailById
//
//	@Description:
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GetCommentDetailByIdLogic) GetCommentDetailById(req *types.GetCommentByIdReq) (resp *types.GetCommentByIdResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// rpc
	idReq := &pb.GetResCommentByIdReq{ID: req.ID}
	detail, err := l.svcCtx.ResourcesRpc.GetResCommentById(deadlineCtx, idReq)
	if err != nil {
		return nil, err
	}
	// copier
	resp = &types.GetCommentByIdResp{}
	err = copier.Copy(&resp.Comment, detail.GetResComment())
	if err != nil {
		l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
		return nil, xerr.NewErrMsg("数据处理异常")
	}
	return resp, nil
}
