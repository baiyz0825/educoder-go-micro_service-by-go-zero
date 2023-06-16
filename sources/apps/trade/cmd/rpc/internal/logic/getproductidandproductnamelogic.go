package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductIdAndProductNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductIdAndProductNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductIdAndProductNameLogic {
	return &GetProductIdAndProductNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductIdAndProductName 通过商品id获取商品名称 productId集合
func (l *GetProductIdAndProductNameLogic) GetProductIdAndProductName(in *pb.GetProductIdAndProductNameReq) (*pb.GetProductIdAndProductNameResp, error) {
	// 参数检查
	if len(in.GetProductId()) == 0 {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	Q := l.svcCtx.Query.Product
	withContextQ := Q.WithContext(context.Background())
	// 检索删除部分
	if in.GetIsDelete() {
		withContextQ.Unscoped()
	}
	// 查询
	find, err := withContextQ.Where(Q.ID.In(in.ProductId...)).Find()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 组床数据
	var respProductBind []*pb.ProductNameAndIdBindId
	for _, product := range find {
		temp := &pb.ProductNameAndIdBindId{
			ProductId:   product.ID,
			ProductName: product.Name,
			ResourceId:  product.ProductBind,
		}
		respProductBind = append(respProductBind, temp)
	}
	// 返回数据
	return &pb.GetProductIdAndProductNameResp{
		ProductInfo: respProductBind,
	}, nil
}
