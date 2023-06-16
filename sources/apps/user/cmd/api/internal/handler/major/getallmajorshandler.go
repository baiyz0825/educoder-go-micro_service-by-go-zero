package major

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/logic/major"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
)

func GetAllMajorsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := major.NewGetAllMajorsLogic(r.Context(), svcCtx)
		resp, err := l.GetAllMajors()
		respresult.ApiResult(r, w, resp, err)
	}
}
