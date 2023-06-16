package utils

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type OSSClient struct {
	AccessKeyId     string
	AccessKeySecret string
	Client          *oss.Client
	BucketName      string
	EndPoint        string
}

// InitOssClient
// @Description: 初始化阿里云工具包配置
// @param endpoint
// @param accessKeyId
// @param accessKeySecret
// @param bucketName
func InitOssClient(accessKeyId, accessKeySecret, endpoint, bucketName string) *OSSClient {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(fmt.Sprintf("初始化OSS客户端失败：%v", err))
	}
	return &OSSClient{
		Client:     client,
		BucketName: bucketName,
		EndPoint:   endpoint,
	}
}

// UploadFile
// @Description: 上传文件到OSS
// @receiver o
// @param fileName
// @param filePath
// @param dataBytes
// @return error
func (o *OSSClient) UploadFile(fileName, filePath string, dataBytes []byte) (error, string) {
	if len(fileName) <= 0 || len(filePath) <= 0 || dataBytes == nil {
		return errors.New("上传必须指定文件名称、文件路径、文件内容"), ""
	}
	bucket, err := o.Client.Bucket(o.BucketName)
	if err != nil {
		return errors.Wrap(err, "初始化oss链接失败"), ""
	}
	// 转化为reader
	reader := bytes.NewReader(dataBytes)
	// 拼接上传路径
	fileUploadPathStr := filePath + "/" + fileName
	err = bucket.PutObject(fileUploadPathStr, reader)
	if err != nil {
		return errors.Wrap(err, "上传文件失败"), ""
	}
	return nil, fileUploadPathStr
}

// GetOssFileFullAccessPath
//
//	@Description: 返回OSS公共读写路径
//	@receiver o
//	@param fileFullPath 文件完整上传路径
//	@return string
func (o *OSSClient) GetOssFileFullAccessPath(fileFullPath string) string {
	return "https://" + o.BucketName + "." + o.EndPoint + "/" + fileFullPath
}

// DownloadFile
//
//	@Description: 从oss下载文件
//	@receiver o
//	@param fileUploadPathStr
//	@return []byte
//	@return error
func (o *OSSClient) DownloadFile(fileUploadPathStr string) ([]byte, error) {
	bucket, err := o.Client.Bucket(o.BucketName)
	if err != nil {
		return nil, errors.Wrap(err, "初始化oss链接失败")
	}
	object, err := bucket.GetObject(fileUploadPathStr)
	if err != nil {
		logx.Errorf("下载OSS文件失败：%v", err)
		return nil, errors.Wrap(err, "下载文件失败")
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, object); err != nil {
		logx.Errorf("下载OSS文件失败：%v", err)
		return nil, err
	}
	defer func(object io.ReadCloser) {
		err := object.Close()
		if err != nil {

		}
	}(object)
	return buf.Bytes(), nil
}

// DeleteFile
// @Description: 从OSS删除文件
// @receiver o
// @param fileFullPath
// @return error
func (o *OSSClient) DeleteFile(fileFullPath string) error {
	if len(fileFullPath) <= 0 {
		return errors.New("删除文件必须指定文件完整路径")
	}
	bucket, err := o.Client.Bucket(o.BucketName)
	if err != nil {
		return errors.Wrap(err, "删除文件失败")
	}
	err = bucket.DeleteObject(fileFullPath)
	if err != nil {
		return errors.Wrap(err, "删除文件失败")
	}
	return nil
}
