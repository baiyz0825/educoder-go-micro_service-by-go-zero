package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderStatusByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderStatusByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderStatusByUUIDLogic {
	return &GetOrderStatusByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderStatusByUUIDLogic) GetOrderStatusByUUID(in *pb.GetOrderStatusByUUIDReq) (*pb.GetOrderStatusByUUIDResp, error) {
	order := l.svcCtx.Query.Order
	orderData, err := order.WithContext(context.Background()).Select(order.Status, order.UUID).Where(order.UUID.Eq(in.GetOrderUUid())).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 订单不存在
	if err == gorm.ErrRecordNotFound {
		return &pb.GetOrderStatusByUUIDResp{
			IsDelete: true,
		}, nil
	}
	// 返回订单状态
	return &pb.GetOrderStatusByUUIDResp{
		Status:   orderData.Status,
		IsDelete: false,
	}, nil
}
