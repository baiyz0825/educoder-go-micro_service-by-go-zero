package classification

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/logic/classification"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
)

func GetAllClassificationsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := classification.NewGetAllClassificationsLogic(r.Context(), svcCtx)
		resp, err := l.GetAllClassifications()
		respresult.ApiResult(r, w, resp, err)
	}
}
