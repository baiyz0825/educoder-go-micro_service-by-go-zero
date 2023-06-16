package textRes

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTextResLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTextResLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTextResLogic {
	return &DeleteTextResLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteTextRes
//
//	@Description: 删除文本类型资源
//	@receiver l
//	@param req
//	@return error
func (l *DeleteTextResLogic) DeleteTextRes(req *types.DelTextReq) error {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	// 删除
	_, err := l.svcCtx.ResourcesRpc.DelOnlineText(l.ctx, &pb.DelOnlineTextReq{ID: req.Id})
	if err != nil {
		return err
	}
	return nil
}
