package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductByResourcesBindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductByResourcesBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductByResourcesBindLogic {
	return &SearchProductByResourcesBindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchProductByResourcesBind
//
//	@Description: 使用商品绑定的资源id获取相应的商品id
//	@receiver l
//	@param in
//	@return *pb.SearchProductByResourcesBindResp
//	@return error
func (l *SearchProductByResourcesBindLogic) SearchProductByResourcesBind(in *pb.SearchProductByResourcesBindReq) (*pb.SearchProductByResourcesBindResp, error) {
	if in.GetResourceId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	product := l.svcCtx.Query.Product
	first, err := product.WithContext(l.ctx).Select(product.ID, product.Name).Where(product.ProductBind.Eq(in.GetResourceId())).First()
	if err == gorm.ErrRecordNotFound {
		// return
		return &pb.SearchProductByResourcesBindResp{ProductId: 0}, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	return &pb.SearchProductByResourcesBindResp{
		ProductId:   first.ID,
		ProductName: first.Name,
	}, nil
}
