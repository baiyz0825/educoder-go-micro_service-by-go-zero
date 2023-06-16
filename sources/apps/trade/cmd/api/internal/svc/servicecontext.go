package svc

import (
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/resourcesrpc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/traderpc"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	TradeRpc     traderpc.Traderpc
	ResourcesRpc resourcesrpc.Resourcesrpc
	Validator    *utils.Validator
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		TradeRpc:     traderpc.NewTraderpc(zrpc.MustNewClient(c.TradeRpc)),
		ResourcesRpc: resourcesrpc.NewResourcesrpc(zrpc.MustNewClient(c.ResourcesRpc)),
		Validator:    utils.GetValidator(),
	}
}
