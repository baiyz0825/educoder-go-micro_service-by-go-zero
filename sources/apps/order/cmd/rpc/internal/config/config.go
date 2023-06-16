package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	AliPay     struct {
		AppId                string
		IsProduction         bool
		PayCallBackUrl       string
		NoticeUrl            string
		ALiPayRootCertPath   string
		ALiPayPublicCertPath string
		AppPublicCertPath    string
		AppPrivateKeyPath    string
		ContentAesKey        string
		PayNoticeCallBackUrl string
	}
	TradeRpcConfig zrpc.RpcClientConf
}
