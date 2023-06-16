package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	AliPay struct {
		AppId                string
		IsProduction         bool
		PayCallBackUrl       string
		NoticeUrl            string
		ALiPayRootCertPath   string
		ALiPayPublicCertPath string
		AppPublicCertPath    string
		AppPrivateKeyPath    string
		ContentAesKey        string
	}
	OrderRpcConfig     zrpc.RpcClientConf
	ResourcesRpcConfig zrpc.RpcClientConf
	UserRpcConfig      zrpc.RpcClientConf
	TradeRpcConfig     zrpc.RpcClientConf
}
