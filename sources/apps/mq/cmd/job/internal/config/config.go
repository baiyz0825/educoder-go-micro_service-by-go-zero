package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	// 基本服务配置
	service.ServiceConf
	Redis           redis.RedisConf
	OrderRpcConfig  zrpc.RpcClientConf
	ResourcesConfig zrpc.RpcClientConf
	TradeConfig     zrpc.RpcClientConf
}
