package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderStatus
//
//	@Description: 更新或者删除订单状态
//	@receiver l
//	@param in
//	@return *pb.UpdateOrderStatusResp
//	@return error
func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	order := l.svcCtx.Query.Order
	// 订单需要删除
	if in.NeedDelete {
		_, err := order.WithContext(context.Background()).Where(order.UUID.Eq(in.GetUuid())).Delete()
		if err != nil && err != gorm.ErrRecordNotFound {
			l.Logger.WithFields(logx.Field("err:", err), logx.Field("订单uuid:", in.GetUuid())).Error("更新订单状态失败")
			return nil, errors.Wrap(err, "数据库错误，删除失败！")
		}
	}
	// 不需要删除，更新订单状态
	_, err := order.WithContext(context.Background()).Where(order.UUID.Eq(in.GetUuid())).Update(order.Status, in.Status)
	if err != nil && err != gorm.ErrRecordNotFound {
		l.Logger.WithFields(logx.Field("err:", err), logx.Field("订单uuid:", in.GetUuid())).Error("更新订单状态失败")
		return nil, errors.Wrap(err, "数据库错误，更新订单失败！")
	}
	return &pb.UpdateOrderStatusResp{}, nil
}
