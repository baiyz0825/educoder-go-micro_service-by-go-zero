package comment

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/logic/comment"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddUserCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddResCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewAddUserCommentLogic(r.Context(), svcCtx)
		err := l.AddUserComment(&req)
		respresult.ApiResult(r, w, nil, err)
	}
}
