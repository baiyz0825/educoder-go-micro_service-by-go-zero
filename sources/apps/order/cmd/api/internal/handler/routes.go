// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	aliCallback "github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/handler/aliCallback"
	order "github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/handler/order"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/order/info",
				Handler: order.GetOrderInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/order/all",
				Handler: order.GetUserAllOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/order/do",
				Handler: order.DoOrderHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/trade/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/order/apy/ali/callback",
				Handler: aliCallback.AlipayNoticeCallBackHandler(serverCtx),
			},
		},
		rest.WithPrefix("/trade/v1"),
	)
}