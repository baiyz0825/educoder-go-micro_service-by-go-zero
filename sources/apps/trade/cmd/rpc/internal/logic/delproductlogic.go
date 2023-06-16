package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProductLogic {
	return &DelProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DelProduct
// @Description: 删除商品
// @receiver l
// @param in
// @return *pb.DelProductResp
// @return error
func (l *DelProductLogic) DelProduct(in *pb.DelProductReq) (*pb.DelProductResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check data
	if in.GetUUID() == 0 && in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	product := l.svcCtx.Query.Product
	// delete db
	if in.GetID() != 0 {
		_, err := product.WithContext(l.ctx).Where(product.ID.Eq(in.GetID())).Delete()
		if err != nil {
			l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_DELETE_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
		}
	} else {
		_, err := product.WithContext(l.ctx).Where(product.UUID.Eq(in.GetUUID())).Delete()
		if err != nil {
			l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_DELETE_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
		}
	}
	return &pb.DelProductResp{}, nil
}
