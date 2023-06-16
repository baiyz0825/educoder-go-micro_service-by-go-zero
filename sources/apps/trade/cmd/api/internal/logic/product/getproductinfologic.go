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

type GetProductInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductInfoLogic {
	return &GetProductInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductInfoLogic) GetProductInfo(req *types.GetProductInfoReq) (resp *types.GetProductInfoResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	// rpc
	product, err := l.svcCtx.TradeRpc.GetProductById(l.ctx, &pb.GetProductByIdReq{
		ID:   req.ID,
		UUID: req.UUID,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetProductInfoResp{}
	err = copier.Copy(&resp.Product, product.GetProduct())
	if err != nil {
		l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
		return nil, xerr.NewErrMsg("数据处理异常")
	}
	resp.Product.FileType = product.GetProduct().Type
	return resp, nil
}
