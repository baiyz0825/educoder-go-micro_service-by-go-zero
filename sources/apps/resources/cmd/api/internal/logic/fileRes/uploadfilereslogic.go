package fileRes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strconv"

	"golang.org/x/sync/errgroup"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/gabriel-vasile/mimetype"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileResLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	FileHeader       multipart.FileHeader
	File             multipart.File
	FilePoster       multipart.File
	FilePosterHeader multipart.FileHeader
}

func NewUploadFileResLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileResLogic {
	return &UploadFileResLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UploadFileRes
//
//	@Description: 上传文件类型资源
//	@receiver l
//	@param req
//	@return error
func (l *UploadFileResLogic) UploadFileRes(req *types.UploadFileReq) error {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	// 获取文件内容
	fileCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// 获取token userId
	uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	if err != nil {
		return xerr.NewErrCode(xerr.FILE_UPLOAD_ERR)
	}
	errorGroup, _ := errgroup.WithContext(context.Background())
	// 资源文件本体
	var resFileType int64
	var fileOssSubPath string
	errorGroup.Go(func() error {
		fileContent, err := io.ReadAll(l.File)
		if err != nil {
			return xerr.NewFileErrMsg("文件内容读取失败")
		}
		defer func(File multipart.File) {
			err := File.Close()
			if err != nil {

			}
		}(l.File)
		mimeType := mimetype.Detect(fileContent)
		// 检查支持文件类型
		mimeName, fileType := utils.JudgeIsSupportedFileType(mimeType)
		if mimeName == "" {
			return xerr.NewFileErrMsg("不支持的文件内容")
		}
		resFileType = fileType
		// 上传文件到cos
		filename := req.Name + "-" + strconv.FormatInt(uid, 10) + "-" + utils.RandString(5) + path.Ext(l.FileHeader.Filename)
		err, filePath := l.svcCtx.OSSClient.UploadFile(filename, l.svcCtx.Config.AliCloud.CommonPath, fileContent)
		if err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).Error(fmt.Sprintf("上传文件到cos失败，文件名称:%v,文件路径:%v",
				filename, filePath))
			return xerr.NewErrMsg("系统错误：文件上传失败")
		}
		fileOssSubPath = filePath
		return nil
	})
	// 文件头图
	var filePosterOssSubPath string
	errorGroup.Go(func() error {
		fileContent, err := io.ReadAll(l.FilePoster)
		if err != nil {
			return xerr.NewFileErrMsg("文件内容读取失败")
		}
		defer func(FilePoster multipart.File) {
			err := FilePoster.Close()
			if err != nil {

			}
		}(l.FilePoster)
		// 不支持图片类型
		if !utils.JudgeIsSupportImage(mimetype.Detect(fileContent)) {
			return xerr.NewFileErrMsg("不支持的图片类型")
		}
		// 上传文件到cos
		filename := req.Name + "-" + strconv.FormatInt(uid, 10) + "-" + utils.RandString(5) + path.Ext(l.FilePosterHeader.Filename)
		err, filePath := l.svcCtx.OSSClient.UploadFile(filename, l.svcCtx.Config.AliCloud.CachePath, fileContent)
		if err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).Error(fmt.Sprintf("上传文件到cos失败，文件名称:%v,文件路径:%v",
				filename, filePath))
			return xerr.NewErrMsg("系统错误：文件上传失败")
		}
		filePosterOssSubPath = filePath
		return nil
	})

	err = errorGroup.Wait()
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).Error("头图和文件资源上传失败！")
		return err
	}
	// 远程保存文件路径
	deadlineCtx, cancelFunc := context.WithDeadline(fileCtx, utils.GetContextDefaultTime())
	defer cancelFunc()

	fileResPb := &pb.AddFileReq{
		Name:          req.Name,
		ObfuscateName: l.ctx.Value(xconst.JWT_USER_ID).(json.Number).String() + "-" + utils.RandString(10) + "-" + l.FileHeader.Filename,
		Size:          utils.ByteToKB(l.FileHeader.Size),
		Owner:         uid,
		Status:        2,
		Type:          resFileType,
		Class:         req.Class,
		Suffix:        path.Ext(l.FileHeader.Filename),
		FilePoster:    l.svcCtx.OSSClient.GetOssFileFullAccessPath(filePosterOssSubPath),
		DownloadAllow: req.DownloadAllow,
		Link:          fileOssSubPath,
	}
	_, err = l.svcCtx.ResourcesRpc.AddFile(deadlineCtx, fileResPb)
	if err != nil {
		return err
	}
	return nil
}
