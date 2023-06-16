package order

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/logic/order"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DoOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DoOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewDoOrderLogic(r.Context(), svcCtx)
		resp, err := l.DoOrder(&req)
		respresult.ApiResult(r, w, resp, err)
	}
}
