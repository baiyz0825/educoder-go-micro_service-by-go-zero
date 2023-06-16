package logic

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByIdLogic {
	return &GetProductByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductById
// @Description: 通过uuid或者id获取商品
// @receiver l
// @param in
// @return *pb.GetProductByIdResp
// @return error
func (l *GetProductByIdLogic) GetProductById(in *pb.GetProductByIdReq) (*pb.GetProductByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check data
	if in.GetUUID() == 0 && in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	product := l.svcCtx.Query.Product
	// find db
	productData, err := product.WithContext(l.ctx).Where(utils.If(
		in.GetID() == 0,
		product.UUID.Eq(in.GetUUID()),
		product.ID.Eq(in.GetID()),
	).(gen.Condition)).First()
	if err == gorm.ErrRecordNotFound {
		// return
		return &pb.GetProductByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_DELETE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
	}
	// copy data
	p := &pb.Product{}
	err = copier.Copy(p, productData)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("复制db -> pb错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	p.CreateTime = productData.CreateTime.UnixMilli()
	p.UpdateTime = productData.UpdateTime.UnixMilli()
	// return
	return &pb.GetProductByIdResp{
		Product: p,
	}, nil
}
