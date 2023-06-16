package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type AliCloud struct {
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	EndPoint        string
	CommonPath      string
}

type Config struct {
	rest.RestConf
	TradeRpc zrpc.RpcClientConf
	Auth     struct {
		AccessSecret string
		AccessExpire int64
	}
	AliCloud     AliCloud
	ResourcesRpc zrpc.RpcClientConf
}
