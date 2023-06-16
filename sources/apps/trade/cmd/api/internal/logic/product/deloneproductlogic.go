package product

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOneProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOneProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOneProductLogic {
	return &DelOneProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DelOneProduct
//
//	@Description: 删除商品
//	@receiver l
//	@param req
//	@return error
func (l *DelOneProductLogic) DelOneProduct(req *types.DelOneReq) error {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	// rpc
	_, err := l.svcCtx.TradeRpc.DelProduct(l.ctx, &pb.DelProductReq{
		ID:   req.ID,
		UUID: req.UUID,
	})
	if err != nil {
		return err
	}
	return nil
}
