package deplete

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/logic/deplete"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
)

func GetFileAndSpaceInsightHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := deplete.NewGetFileAndSpaceInsightLogic(r.Context(), svcCtx)
		resp, err := l.GetFileAndSpaceInsight()
		respresult.ApiResult(r, w, resp, err)
	}
}
