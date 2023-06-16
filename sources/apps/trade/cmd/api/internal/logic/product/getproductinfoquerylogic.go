package product

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductInfoQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductInfoQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductInfoQueryLogic {
	return &GetProductInfoQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductInfoQueryLogic) GetProductInfoQuery(req *types.SearchProductByConditionReq) (resp *types.SearchProductByConditionResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	// rpc
	pbProducts, err := l.svcCtx.TradeRpc.SearchProduct(l.ctx, &pb.SearchProductByConditionReq{
		Page:        req.Page,
		Limit:       req.Limit,
		Type:        req.ProductType,
		Name:        req.Name,
		BottomPrice: req.BottonPrice,
		TopPrice:    req.TopPrice,
		Desc:        req.Desc,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.SearchProductByConditionResp{}
	for _, product := range pbProducts.GetProduct() {
		temp := types.Product{}
		err := copier.Copy(&temp, product)
		if err != nil {
			l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
			return nil, xerr.NewErrMsg("数据处理异常")
		}
		temp.FileType = product.Type
		resp.Products = append(resp.Products, temp)
	}
	return resp, nil
}
