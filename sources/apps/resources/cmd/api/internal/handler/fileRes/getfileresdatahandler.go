package fileRes

import (
	"fmt"
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/logic/fileRes"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFileResDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownLoadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := fileRes.NewGetFileResDataLogic(r.Context(), svcCtx)
		data, fileName, err := l.GetFileResData(&req)
		if err != nil {
			respresult.ApiResult(r, w, nil, err)
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v", fileName))
		if _, err = w.Write(data); err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).Error("文件回写客户端错误")
			respresult.ApiResult(r, w, nil, xerr.NewErrMsg("文件传输失败"))
		}
	}
}
