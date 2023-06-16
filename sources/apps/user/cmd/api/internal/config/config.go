package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type AliCloud struct {
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	EndPoint        string
	UserCachePath   string
}

type Config struct {
	rest.RestConf
	UserRpcConfig zrpc.RpcClientConf
	Auth          struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis    redis.RedisConf
	AliCloud AliCloud
}
