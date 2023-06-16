package textRes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strconv"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/gabriel-vasile/mimetype"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadTextResLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	FileHeader multipart.FileHeader
	File       multipart.File
}

func NewUploadTextResLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadTextResLogic {
	return &UploadTextResLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UploadTextRes
//
//	@Description: 上传在线文本资源
//	@receiver l
//	@param req
//	@return error
func (l *UploadTextResLogic) UploadTextRes(req *types.UploadTextReq) error {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	if err != nil {
		return xerr.NewErrCode(xerr.SERVER_ERROR)
	}

	// 获取文件内容
	fileContent, err := io.ReadAll(l.File)
	defer func(File multipart.File) {
		err := File.Close()
		if err != nil {

		}
	}(l.File)
	if err != nil {
		return xerr.NewFileErrMsg("文件内容读取失败")
	}
	mimeType := mimetype.Detect(fileContent)

	if !utils.JudgeIsSupportImage(mimeType) {
		return xerr.NewFileErrMsg("不支持的头图图片格式")
	}
	// 上传文件到cos
	filename := req.TextName + "-" + strconv.FormatInt(uid, 10) + "-" + utils.RandString(5) + path.Ext(l.FileHeader.Filename)
	err, fileOssSubPath := l.svcCtx.OSSClient.UploadFile(filename, l.svcCtx.Config.AliCloud.CachePath, fileContent)
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).Error(fmt.Sprintf("上传头图到cos失败，文件名称:%v,文件路径:%v",
			filename, l.svcCtx.Config.AliCloud.CommonPath))
		return xerr.NewErrMsg("系统错误：头图上传失败")
	}

	// rpc
	addData := &pb.AddOnlineTextReq{
		TypeSuffix: req.TypeSuffix,
		Owner:      uid,
		Content:    req.Content,
		ClassID:    req.ClassID,
		Permission: req.Permission,
		TextName:   req.TextName,
		TextPoster: l.svcCtx.OSSClient.GetOssFileFullAccessPath(fileOssSubPath),
	}
	_, err = l.svcCtx.ResourcesRpc.AddOnlineText(l.ctx, addData)
	if err != nil {
		return err
	}
	return nil
}
