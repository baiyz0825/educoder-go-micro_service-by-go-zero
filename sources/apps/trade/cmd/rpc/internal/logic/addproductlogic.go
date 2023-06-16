package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProduct
// @Description: 增加商品信息
// @receiver l
// @param in
// @return *pb.AddProductResp
// @return error
func (l *AddProductLogic) AddProduct(in *pb.AddProductReq) (*pb.AddProductResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check pb data
	if len(in.GetName()) == 0 || in.GetType() == 0 || in.GetOwner() == 0 || in.GetPrice() == 0 || in.GetProductBind() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// gen uuid && db data
	id, err := utils.GenSnowFlakeId()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("生成系统id错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// set db data
	product := &model.Product{
		UUID:          id,
		Name:          in.GetName(),
		Type:          in.GetType(),
		ProductBind:   in.GetProductBind(),
		Owner:         in.GetOwner(),
		Price:         in.GetPrice(),
		Saled:         0,
		ProductPoster: &in.ProductPoster,
	}
	p := l.svcCtx.Query.Product
	// insert db
	err = p.WithContext(l.ctx).Create(product)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	return &pb.AddProductResp{}, nil
}
