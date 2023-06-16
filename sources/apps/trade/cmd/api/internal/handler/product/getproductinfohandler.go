package product

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/logic/product"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewGetProductInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetProductInfo(&req)
		respresult.ApiResult(r, w, resp, err)
	}
}
