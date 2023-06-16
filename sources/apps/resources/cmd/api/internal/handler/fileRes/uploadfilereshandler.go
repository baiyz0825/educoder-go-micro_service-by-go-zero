package fileRes

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/logic/fileRes"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadFileResHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			respresult.ApiResult(r, w, nil, xerr.NewFileErrMsg("请求参数错误，请检查"))
		}
		// 最大文件大小限制
		err := r.ParseMultipartForm(64 << 20) // 32MB
		if err != nil {
			// 处理错误信息
			respresult.ApiResult(r, w, nil, xerr.NewFileErrMsg("文件超过限制，请上传64M以下文件，并且格式符合xls、pdf、world、xlsx等常见格式"))
		}
		l := fileRes.NewUploadFileResLogic(r.Context(), svcCtx)
		file, header, err := r.FormFile("file")
		if err != nil {
			respresult.ApiResult(r, w, nil, xerr.NewFileErrMsg("服务器异常，请稍后重新上传文件资源"))
		}
		l.File = file
		l.FileHeader = *header

		// 处理头图
		file, header, err = r.FormFile("filePoster")
		if err != nil {
			respresult.ApiResult(r, w, nil, xerr.NewErrMsg("服务器异常，请稍后重新上传文件头图"))
		}
		l.FilePoster = file
		l.FilePosterHeader = *header
		// 处理逻辑
		err = l.UploadFileRes(&req)
		respresult.ApiResult(r, w, nil, err)
	}
}
