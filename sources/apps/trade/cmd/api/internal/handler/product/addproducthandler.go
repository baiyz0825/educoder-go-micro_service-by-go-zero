package product

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/logic/product"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewAddProductLogic(r.Context(), svcCtx)
		err := l.AddProduct(&req)
		respresult.ApiResult(r, w, nil, err)
	}
}
