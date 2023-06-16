package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strconv"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/gabriel-vasile/mimetype"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserAvatarLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	AvatarFileHeader multipart.FileHeader
	AvatarFile       multipart.File
}

func NewUploadUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UploadUserAvatar
//
//	@Description: 上传用户头像
//	@receiver l
//	@return error
func (l *UploadUserAvatarLogic) UploadUserAvatar() error {
	userId, err := strconv.ParseInt(l.ctx.Value(xconst.JWT_USER_ID).(json.Number).String(), 10, 64)
	if err != nil {
		return err
	}
	userUniqueId, err := strconv.ParseInt(l.ctx.Value(xconst.JWT_USER_USERUNIQUEID).(json.Number).String(), 10, 64)
	if err != nil {
		return err
	}
	// 上传用户头像
	fileContent, err := io.ReadAll(l.AvatarFile)
	if err != nil {
		return xerr.NewErrMsg("文件内容读取失败")
	}
	defer func(AvatarFile multipart.File) {
		err := AvatarFile.Close()
		if err != nil {

		}
	}(l.AvatarFile)
	// 不支持图片类型
	filePath := ""
	if utils.JudgeIsSupportImage(mimetype.Detect(fileContent)) {
		// 上传文件到cos
		filename := utils.RandString(5) + "-" + l.ctx.Value(xconst.JWT_USER_USERUNIQUEID).(json.Number).String() + "-" + strconv.FormatInt(userId, 10) + "-" + utils.RandString(5) + path.Ext(l.AvatarFileHeader.Filename)
		err, filePath = l.svcCtx.OSSClient.UploadFile(filename, l.svcCtx.Config.AliCloud.UserCachePath, fileContent)
		if err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).Error(fmt.Sprintf("上传文件到cos失败，文件名称:%v,文件路径:%v",
				filename, filePath))
			return xerr.NewErrMsg("系统错误：用户头像上传失败")
		}
	}
	userInfo := &pb.UpdateUserReq{
		UID:      userId,
		UniqueID: userUniqueId,
		Avatar:   utils.If(len(filePath) != 0, l.svcCtx.OSSClient.GetOssFileFullAccessPath(filePath), "").(string),
	}
	ctx, cancelFunc := context.WithDeadline(l.ctx, utils.GetContextDefaultTime())
	defer cancelFunc()
	_, err = l.svcCtx.UserRpc.UpdateUser(ctx, userInfo)
	if err != nil {
		return err
	}
	return nil
}
