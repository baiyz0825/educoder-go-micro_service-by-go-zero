package svc

import (
	"errors"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/common"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/userrpc"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	Captcha     *base64Captcha.Captcha
	UserRpc     userrpc.Userrpc
	Validator   *utils.Validator
	OSSClient   *utils.OSSClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	// init
	rd := initRedisClient(c)
	captcha := common.InitCaptcha(rd)
	return &ServiceContext{
		Config:      c,
		RedisClient: rd,
		Captcha:     captcha,
		UserRpc:     userrpc.NewUserrpc(zrpc.MustNewClient(c.UserRpcConfig)),
		Validator:   utils.GetValidator(),
		OSSClient:   utils.InitOssClient(c.AliCloud.AccessKeyId, c.AliCloud.AccessKeySecret, c.AliCloud.EndPoint, c.AliCloud.BucketName),
	}
}

// 初始化redis
func initRedisClient(c config.Config) *redis.Redis {
	// 创建 Redis 配置
	r := redis.MustNewRedis(c.Redis)
	if r != nil {
		return r
	}
	panic(errors.New("初始化Redis失败"))
}
