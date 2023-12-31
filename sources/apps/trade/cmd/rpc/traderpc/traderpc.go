// Code generated by goctl. DO NOT EDIT.
// Source: traderpc.proto

package traderpc

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddProductReq                    = pb.AddProductReq
	AddProductResp                   = pb.AddProductResp
	DelProductReq                    = pb.DelProductReq
	DelProductResp                   = pb.DelProductResp
	GetProductBindByProductIdReq     = pb.GetProductBindByProductIdReq
	GetProductBindByProductIdResp    = pb.GetProductBindByProductIdResp
	GetProductByBindIdAndOwnerReq    = pb.GetProductByBindIdAndOwnerReq
	GetProductByBindIdAndOwnerResp   = pb.GetProductByBindIdAndOwnerResp
	GetProductByIdReq                = pb.GetProductByIdReq
	GetProductByIdResp               = pb.GetProductByIdResp
	GetProductIdAndProductNameReq    = pb.GetProductIdAndProductNameReq
	GetProductIdAndProductNameResp   = pb.GetProductIdAndProductNameResp
	Product                          = pb.Product
	ProductNameAndIdBindId           = pb.ProductNameAndIdBindId
	SearchProductByConditionReq      = pb.SearchProductByConditionReq
	SearchProductByConditionResp     = pb.SearchProductByConditionResp
	SearchProductByResourcesBindReq  = pb.SearchProductByResourcesBindReq
	SearchProductByResourcesBindResp = pb.SearchProductByResourcesBindResp
	UpdateProductReq                 = pb.UpdateProductReq
	UpdateProductResp                = pb.UpdateProductResp

	Traderpc interface {
		// -----------------------产品信息-----------------------
		AddProduct(ctx context.Context, in *AddProductReq, opts ...grpc.CallOption) (*AddProductResp, error)
		UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error)
		DelProduct(ctx context.Context, in *DelProductReq, opts ...grpc.CallOption) (*DelProductResp, error)
		GetProductById(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResp, error)
		SearchProduct(ctx context.Context, in *SearchProductByConditionReq, opts ...grpc.CallOption) (*SearchProductByConditionResp, error)
		SearchProductByResourcesBind(ctx context.Context, in *SearchProductByResourcesBindReq, opts ...grpc.CallOption) (*SearchProductByResourcesBindResp, error)
		GetProductBindByProductId(ctx context.Context, in *GetProductBindByProductIdReq, opts ...grpc.CallOption) (*GetProductBindByProductIdResp, error)
		// 通过商品id获取商品名称
		GetProductIdAndProductName(ctx context.Context, in *GetProductIdAndProductNameReq, opts ...grpc.CallOption) (*GetProductIdAndProductNameResp, error)
		// 通过商品绑定用户和绑定商品查询商品详情
		GetProductByBindIdAndOwner(ctx context.Context, in *GetProductByBindIdAndOwnerReq, opts ...grpc.CallOption) (*GetProductByBindIdAndOwnerResp, error)
	}

	defaultTraderpc struct {
		cli zrpc.Client
	}
)

func NewTraderpc(cli zrpc.Client) Traderpc {
	return &defaultTraderpc{
		cli: cli,
	}
}

// -----------------------产品信息-----------------------
func (m *defaultTraderpc) AddProduct(ctx context.Context, in *AddProductReq, opts ...grpc.CallOption) (*AddProductResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.AddProduct(ctx, in, opts...)
}

func (m *defaultTraderpc) UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.UpdateProduct(ctx, in, opts...)
}

func (m *defaultTraderpc) DelProduct(ctx context.Context, in *DelProductReq, opts ...grpc.CallOption) (*DelProductResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.DelProduct(ctx, in, opts...)
}

func (m *defaultTraderpc) GetProductById(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.GetProductById(ctx, in, opts...)
}

func (m *defaultTraderpc) SearchProduct(ctx context.Context, in *SearchProductByConditionReq, opts ...grpc.CallOption) (*SearchProductByConditionResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.SearchProduct(ctx, in, opts...)
}

func (m *defaultTraderpc) SearchProductByResourcesBind(ctx context.Context, in *SearchProductByResourcesBindReq, opts ...grpc.CallOption) (*SearchProductByResourcesBindResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.SearchProductByResourcesBind(ctx, in, opts...)
}

func (m *defaultTraderpc) GetProductBindByProductId(ctx context.Context, in *GetProductBindByProductIdReq, opts ...grpc.CallOption) (*GetProductBindByProductIdResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.GetProductBindByProductId(ctx, in, opts...)
}

// 通过商品id获取商品名称
func (m *defaultTraderpc) GetProductIdAndProductName(ctx context.Context, in *GetProductIdAndProductNameReq, opts ...grpc.CallOption) (*GetProductIdAndProductNameResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.GetProductIdAndProductName(ctx, in, opts...)
}

// 通过商品绑定用户和绑定商品查询商品详情
func (m *defaultTraderpc) GetProductByBindIdAndOwner(ctx context.Context, in *GetProductByBindIdAndOwnerReq, opts ...grpc.CallOption) (*GetProductByBindIdAndOwnerResp, error) {
	client := pb.NewTraderpcClient(m.cli.Conn())
	return client.GetProductByBindIdAndOwner(ctx, in, opts...)
}
