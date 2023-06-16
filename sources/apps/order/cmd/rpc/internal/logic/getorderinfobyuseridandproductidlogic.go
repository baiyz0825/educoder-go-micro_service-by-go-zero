package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoByUserIdAndProductIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoByUserIdAndProductIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoByUserIdAndProductIdLogic {
	return &GetOrderInfoByUserIdAndProductIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderInfoByUserIdAndProductId
//
//	@Description: 使用用户id和商品id查询订单信息
//	@receiver l
//	@param in
//	@return *pb.GetOrderInfoByUserIdAndProductIdResp
//	@return error
func (l *GetOrderInfoByUserIdAndProductIdLogic) GetOrderInfoByUserIdAndProductId(in *pb.GetOrderInfoByUserIdAndProductIdReq) (*pb.GetOrderInfoByUserIdAndProductIdResp, error) {
	if in.GetProductId() == 0 || in.GetUserId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 查询数据
	order := l.svcCtx.Query.Order
	first, err := order.WithContext(l.ctx).Where(order.ProductID.Eq(in.GetProductId()), order.UserID.Eq(in.GetUserId()), order.Status.In(2, 3, 4)).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetOrderInfoByUserIdAndProductIdResp{}, nil
	}
	if err != nil {
		return nil, err
	}
	respData := &pb.Order{}
	err = copier.Copy(respData, first)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}

	return &pb.GetOrderInfoByUserIdAndProductIdResp{
		Order: respData,
	}, nil
}
