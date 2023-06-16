package textRes

import (
	"mime/multipart"
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/logic/textRes"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadTextResHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadTextReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 最大文件大小限制
		err := r.ParseMultipartForm(64 << 20) // 64MB
		if err != nil {
			// 处理错误信息
			respresult.ApiResult(r, w, nil, xerr.NewFileErrMsg("文件超过限制，请上传64M以下文件，并且格式符合常见图片格式"))
		}
		file, header, err := r.FormFile("textPoster")
		if err != nil {
			respresult.ApiResult(r, w, nil, xerr.NewFileErrMsg("请上传文本头图！"))
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		l := textRes.NewUploadTextResLogic(r.Context(), svcCtx)
		l.File = file
		l.FileHeader = *header
		err = l.UploadTextRes(&req)
		respresult.ApiResult(r, w, nil, err)
	}
}
