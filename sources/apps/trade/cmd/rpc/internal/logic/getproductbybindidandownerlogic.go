package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByBindIdAndOwnerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductByBindIdAndOwnerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByBindIdAndOwnerLogic {
	return &GetProductByBindIdAndOwnerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductByBindIdAndOwner 通过商品绑定用户和绑定商品查询商品详情
func (l *GetProductByBindIdAndOwnerLogic) GetProductByBindIdAndOwner(in *pb.GetProductByBindIdAndOwnerReq) (*pb.GetProductByBindIdAndOwnerResp, error) {
	productQ := l.svcCtx.Query.Product
	first, err := productQ.WithContext(context.Background()).Where(productQ.ProductBind.Eq(in.GetProductBind()), productQ.Owner.Eq(in.GetUid())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetProductByBindIdAndOwnerResp{Product: nil}, nil
	}
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	respP := &pb.Product{
		ID:          first.ID,
		UUID:        first.UUID,
		Name:        first.Name,
		Type:        first.Type,
		Owner:       first.Owner,
		ProductBind: first.ProductBind,
		Price:       first.Price,
		Saled:       first.Saled,
		CreateTime:  first.CreateTime.UnixMilli(),
		UpdateTime:  first.UpdateTime.UnixMilli(),
	}
	if first.ProductPoster != nil {
		respP.ProductPoster = *first.ProductPoster
	}
	return &pb.GetProductByBindIdAndOwnerResp{Product: respP}, nil
}
