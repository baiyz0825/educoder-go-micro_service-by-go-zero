package user

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/logic/user"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha()
		respresult.ApiResult(r, w, resp, err)
	}
}
