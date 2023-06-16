package user

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Captcha
// @Description: 获取图片验证码digit
// @receiver l
// @return resp
// @return err
func (l *CaptchaLogic) Captcha() (resp *types.ChaptchaResp, err error) {
	captchaId, captcha, err := l.svcCtx.Captcha.Generate()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.CAPTCHA_GEN_ERR)
	}

	return &types.ChaptchaResp{
		CaptchaB64: captcha,
		CaptchaId:  captchaId,
	}, nil
}
