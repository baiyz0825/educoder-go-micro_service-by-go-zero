package logic

import (
	"context"

	"gorm.io/gen"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *pb.UpdateProductReq) (*pb.UpdateProductResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check params 默认sale为0
	if (in.GetID() == 0 && in.GetUUID() == 0) || len(in.GetName()) == 0 || in.GetType() == 0 || in.GetPrice() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// make db struct
	product := model.Product{
		Name:  in.GetName(),
		Type:  in.GetType(),
		Price: in.GetPrice(),
		Saled: in.GetSaled(),
	}
	// update
	p := l.svcCtx.Query.Product
	_, err := p.WithContext(l.ctx).Where(utils.If(
		in.GetID() == 0,
		p.UUID.Eq(in.GetUUID()),
		p.ID.Eq(in.GetID())).(gen.Condition),
	).Updates(product)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}
	return &pb.UpdateProductResp{}, nil
}
