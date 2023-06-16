package svc

import (
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/orderrpc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/resourcesrpc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/traderpc"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	ResourcesRpc resourcesrpc.Resourcesrpc
	TradeRpc     traderpc.Traderpc
	OrderRpc     orderrpc.Orderrpc
	Validator    *utils.Validator
	OSSClient    *utils.OSSClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ResourcesRpc: resourcesrpc.NewResourcesrpc(zrpc.MustNewClient(c.ResourcesRpc)),
		TradeRpc:     traderpc.NewTraderpc(zrpc.MustNewClient(c.TradeRpc)),
		OrderRpc:     orderrpc.NewOrderrpc(zrpc.MustNewClient(c.OrderRpc)),
		Validator:    utils.GetValidator(),
		OSSClient:    utils.InitOssClient(c.AliCloud.AccessKeyId, c.AliCloud.AccessKeySecret, c.AliCloud.EndPoint, c.AliCloud.BucketName),
	}
}
