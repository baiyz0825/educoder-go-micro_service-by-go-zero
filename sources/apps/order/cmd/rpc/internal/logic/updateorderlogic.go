package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrder
//
//	@Description: 更新订单状态
//	@receiver l
//	@param in
//	@return *pb.UpdateOrderResp
//	@return error
func (l *UpdateOrderLogic) UpdateOrder(in *pb.UpdateOrderReq) (*pb.UpdateOrderResp, error) {
	if in.GetId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 查询
	order := l.svcCtx.Query.Order
	orderData := &model.Order{
		Status:          in.GetStatus(),
		PayPathOrderNum: in.GetPayPathOrderNum(),
	}
	_, err := order.WithContext(l.ctx).Where(order.ID.Eq(in.GetId())).Updates(orderData)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}
	data, err := order.WithContext(l.ctx).Select(order.ID, order.UUID).Where(order.ID.Eq(in.GetId())).First()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}
	return &pb.UpdateOrderResp{
		Id:   data.ID,
		Uuid: data.UUID,
	}, nil
}
