package textRes

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTextInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTextInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTextInfoLogic {
	return &GetTextInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetTextInfo
//
//	@Description: 使用onlineText id查询onlineText
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GetTextInfoLogic) GetTextInfo(req *types.TextResInfoReq) (resp *types.OnlineText, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}

	// 查询
	data, err := l.svcCtx.ResourcesRpc.GetOnlineTextById(l.ctx, &pb.GetOnlineTextByIdReq{ID: req.TextResId})
	if err != nil || data.OnlineText == nil {
		return nil, err
	}
	resp = &types.OnlineText{}
	err = copier.Copy(resp, data.OnlineText)
	if err != nil {
		l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
		return nil, xerr.NewErrMsg("数据处理异常")
	}
	return resp, nil
}
