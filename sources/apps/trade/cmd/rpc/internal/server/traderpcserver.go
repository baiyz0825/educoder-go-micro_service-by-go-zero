// Code generated by goctl. DO NOT EDIT.
// Source: traderpc.proto

package server

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/logic"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
)

type TraderpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTraderpcServer
}

func NewTraderpcServer(svcCtx *svc.ServiceContext) *TraderpcServer {
	return &TraderpcServer{
		svcCtx: svcCtx,
	}
}

// -----------------------产品信息-----------------------
func (s *TraderpcServer) AddProduct(ctx context.Context, in *pb.AddProductReq) (*pb.AddProductResp, error) {
	l := logic.NewAddProductLogic(ctx, s.svcCtx)
	return l.AddProduct(in)
}

func (s *TraderpcServer) UpdateProduct(ctx context.Context, in *pb.UpdateProductReq) (*pb.UpdateProductResp, error) {
	l := logic.NewUpdateProductLogic(ctx, s.svcCtx)
	return l.UpdateProduct(in)
}

func (s *TraderpcServer) DelProduct(ctx context.Context, in *pb.DelProductReq) (*pb.DelProductResp, error) {
	l := logic.NewDelProductLogic(ctx, s.svcCtx)
	return l.DelProduct(in)
}

func (s *TraderpcServer) GetProductById(ctx context.Context, in *pb.GetProductByIdReq) (*pb.GetProductByIdResp, error) {
	l := logic.NewGetProductByIdLogic(ctx, s.svcCtx)
	return l.GetProductById(in)
}

func (s *TraderpcServer) SearchProduct(ctx context.Context, in *pb.SearchProductByConditionReq) (*pb.SearchProductByConditionResp, error) {
	l := logic.NewSearchProductLogic(ctx, s.svcCtx)
	return l.SearchProduct(in)
}

func (s *TraderpcServer) SearchProductByResourcesBind(ctx context.Context, in *pb.SearchProductByResourcesBindReq) (*pb.SearchProductByResourcesBindResp, error) {
	l := logic.NewSearchProductByResourcesBindLogic(ctx, s.svcCtx)
	return l.SearchProductByResourcesBind(in)
}

func (s *TraderpcServer) GetProductBindByProductId(ctx context.Context, in *pb.GetProductBindByProductIdReq) (*pb.GetProductBindByProductIdResp, error) {
	l := logic.NewGetProductBindByProductIdLogic(ctx, s.svcCtx)
	return l.GetProductBindByProductId(in)
}

// 通过商品id获取商品名称
func (s *TraderpcServer) GetProductIdAndProductName(ctx context.Context, in *pb.GetProductIdAndProductNameReq) (*pb.GetProductIdAndProductNameResp, error) {
	l := logic.NewGetProductIdAndProductNameLogic(ctx, s.svcCtx)
	return l.GetProductIdAndProductName(in)
}

// 通过商品绑定用户和绑定商品查询商品详情
func (s *TraderpcServer) GetProductByBindIdAndOwner(ctx context.Context, in *pb.GetProductByBindIdAndOwnerReq) (*pb.GetProductByBindIdAndOwnerResp, error) {
	l := logic.NewGetProductByBindIdAndOwnerLogic(ctx, s.svcCtx)
	return l.GetProductByBindIdAndOwner(in)
}
