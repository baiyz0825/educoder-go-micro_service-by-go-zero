package svc

import (
	"fmt"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/orderrpc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/resourcesrpc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/traderpc"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	AsynqServer  *asynq.Server
	OrderRpc     orderrpc.Orderrpc
	ResourcesRpc resourcesrpc.Resourcesrpc
	TradeRpc     traderpc.Traderpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AsynqServer:  initAsynqServer(c),
		OrderRpc:     orderrpc.NewOrderrpc(zrpc.MustNewClient(c.OrderRpcConfig)),
		ResourcesRpc: resourcesrpc.NewResourcesrpc(zrpc.MustNewClient(c.ResourcesConfig)),
		TradeRpc:     traderpc.NewTraderpc(zrpc.MustNewClient(c.TradeConfig)),
	}
}

// initAsynqServer
//
//	@Description: 初始化异步任务
//	@param c
//	@return *asynq.Server
func initAsynqServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			// 任务失败回调
			IsFailure: func(err error) bool {
				fmt.Printf("执行任务失败: %+v \n", err)
				return true
			},
			// 最大执行任务并发数量
			Concurrency: 20,
		},
	)
}
