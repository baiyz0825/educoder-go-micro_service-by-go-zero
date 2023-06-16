package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductBindByProductIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductBindByProductIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductBindByProductIdLogic {
	return &GetProductBindByProductIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductBindByProductId
//
//	@Description: 使用商品id查询商品对应绑定的资源id
//	@receiver l
//	@param in
//	@return *pb.GetProductBindByProductIdResp
//	@return error
func (l *GetProductBindByProductIdLogic) GetProductBindByProductId(in *pb.GetProductBindByProductIdReq) (*pb.GetProductBindByProductIdResp, error) {
	product := l.svcCtx.Query.Product
	var productBindId int64
	err := product.WithContext(context.Background()).
		Select(product.ProductBind).
		Where(product.ID.Eq(in.GetProductId())).
		Scan(&productBindId)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductBindByProductIdResp{
		ResourcesBind: productBindId,
	}, nil
}
