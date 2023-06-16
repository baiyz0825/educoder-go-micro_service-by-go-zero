package user

import (
	"mime/multipart"
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/logic/user"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/common/respresult"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
)

func UploadUserAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUploadUserAvatarLogic(r.Context(), svcCtx)
		// avatar不为空，处理头像
		err := r.ParseMultipartForm(64 << 20) // 32MB
		if err != nil {
			// 处理错误信息
			respresult.ApiResult(r, w, nil, xerr.NewErrMsg("文件超过限制，请上传64M以下文件，格式为常见图片格式"))
		}
		// 处理用户头像
		file, header, err := r.FormFile("avatar")
		if err != nil && err != http.ErrMissingFile {
			respresult.ApiResult(r, w, nil, xerr.NewErrMsg("服务器异常，请稍后重新上传用户头图"))
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		l.AvatarFile = file
		l.AvatarFileHeader = *header
		err = l.UploadUserAvatar()
		respresult.ApiResult(r, w, nil, err)
	}
}
