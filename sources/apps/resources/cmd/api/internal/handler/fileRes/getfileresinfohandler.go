package fileRes

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/logic/fileRes"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFileResInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileResInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := fileRes.NewGetFileResInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetFileResInfo(&req)
		respresult.ApiResult(r, w, resp, err)
	}
}
