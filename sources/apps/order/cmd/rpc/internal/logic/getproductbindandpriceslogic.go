package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductBindAndPricesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductBindAndPricesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductBindAndPricesLogic {
	return &GetProductBindAndPricesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductBindAndPrices
//
//	@Description: 获取所有订单产品id下价格消费数据
//	@receiver l
//	@param in
//	@return *pb.GetProductBindAndPricesResp
//	@return error
func (l *GetProductBindAndPricesLogic) GetProductBindAndPrices(in *pb.GetProductBindAndPricesReq) (*pb.GetProductBindAndPricesResp, error) {
	// 订单中同一个商品的价格一致
	order := l.svcCtx.Query.Order
	var productPriceBind []struct {
		ProductID int64
		Total     float64
	}
	// 统计每个订单的数量
	err := order.WithContext(context.Background()).Select(
		order.ProductID,
		order.PayPrice.Sum().As("Total")).
		Group(order.ProductID).
		Scan(&productPriceBind)
	if err != nil {
		return nil, err
	}
	resp := &pb.GetProductBindAndPricesResp{}
	err = copier.Copy(&resp.ProductBindPrice, productPriceBind)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
