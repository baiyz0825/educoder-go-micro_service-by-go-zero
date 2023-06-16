package svc

import (
	"fmt"
	"os"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/orderrpc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/resourcesrpc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/traderpc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/userrpc"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	OrderRpc     orderrpc.Orderrpc
	ResourcesRpc resourcesrpc.Resourcesrpc
	UserRpc      userrpc.Userrpc
	Validator    *utils.Validator
	AliPayClient *alipay.Client
	TradeRpc     traderpc.Traderpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		OrderRpc:     orderrpc.NewOrderrpc(zrpc.MustNewClient(c.OrderRpcConfig)),
		ResourcesRpc: resourcesrpc.NewResourcesrpc(zrpc.MustNewClient(c.ResourcesRpcConfig)),
		UserRpc:      userrpc.NewUserrpc(zrpc.MustNewClient(c.UserRpcConfig)),
		TradeRpc:     traderpc.NewTraderpc(zrpc.MustNewClient(c.TradeRpcConfig)),
		Validator:    utils.GetValidator(),
		AliPayClient: initAlipayClient(c),
	}
}

// initAlipayClient
//
//	@Description: 初始化支付宝客户端
//	@param c
//	@return *alipay.Client
func initAlipayClient(c config.Config) *alipay.Client {
	privateKey, err := os.ReadFile(c.AliPay.AppPrivateKeyPath)
	if err != nil {
		panic(fmt.Sprintf("初始化支付渠道: 支付宝应用私钥失败:%v", err))
	}
	payClient, err := alipay.New(c.AliPay.AppId, string(privateKey), c.AliPay.IsProduction)
	// 加载应用公钥证书
	err = payClient.LoadAppPublicCertFromFile(c.AliPay.AppPublicCertPath)
	if err != nil {
		panic(fmt.Sprintf("初始化支付渠道: 加载应用公钥证书失败:%v", err))
	}
	// 加载支付宝根证书
	err = payClient.LoadAliPayRootCertFromFile(c.AliPay.ALiPayRootCertPath)
	if err != nil {
		panic(fmt.Sprintf("初始化支付渠道: 加载支付宝根证书失败:%v", err))
	}
	// 加载支付宝证书
	err = payClient.LoadAliPayPublicCertFromFile(c.AliPay.ALiPayPublicCertPath)
	if err != nil {
		panic(fmt.Sprintf("初始化支付渠道: 加载支付宝证书失败:%v", err))
	}
	return payClient
}
