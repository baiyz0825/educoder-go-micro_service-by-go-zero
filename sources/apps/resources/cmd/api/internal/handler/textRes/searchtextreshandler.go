package textRes

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/logic/textRes"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchTextResHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchOnlineConditionTextReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := textRes.NewSearchTextResLogic(r.Context(), svcCtx)
		resp, err := l.SearchTextRes(&req)
		respresult.ApiResult(r, w, resp, err)
	}
}
