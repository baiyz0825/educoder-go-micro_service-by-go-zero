package logic

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchProduct
// @Description: 条件检索商品信息
// @receiver l
// @param in
// @return *pb.SearchProductByConditionResp
// @return error
func (l *SearchProductLogic) SearchProduct(in *pb.SearchProductByConditionReq) (*pb.SearchProductByConditionResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	q := l.svcCtx.Query.Product
	query := q.WithContext(l.ctx)
	var condition []gen.Condition

	// 按照条件进行排序
	// name
	if len(in.GetName()) != 0 {
		condition = append(condition, q.Name.Eq(in.GetName()))
	}
	// type
	if in.GetType() != 0 {
		condition = append(condition, q.Type.Eq(in.GetType()))
	}
	// price
	if (in.GetBottomPrice() >= 0 && in.GetTopPrice() > 0) && in.GetTopPrice() > in.GetBottomPrice() {
		condition = append(condition, q.Price.Between(in.GetBottomPrice(), in.GetTopPrice()))
	}
	// pages
	// 负分页
	if in.GetPage() <= 0 || in.GetLimit() <= 0 {
		return nil, xerr.NewErrCode(xerr.RPC_PAGES_PARAM_ERR)
	}
	where := query.Where(condition...)
	// desc
	if in.GetDesc() {
		where.Order(q.ID.Desc())
	}
	results, _, err := where.FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
	if err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	if err == gorm.ErrRecordNotFound {
		return &pb.SearchProductByConditionResp{}, nil
	}
	if results == nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// copy 返回值
	var respProducts []*pb.Product
	for _, result := range results {
		product := &pb.Product{}
		err := copier.Copy(product, result)
		if err != nil {
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}
		product.CreateTime = result.CreateTime.UnixMilli()
		product.UpdateTime = result.UpdateTime.UnixMilli()
		respProducts = append(respProducts, product)
	}
	return &pb.SearchProductByConditionResp{
		Product: respProducts,
	}, nil
}
