package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoByUUIDAndUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoByUUIDAndUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoByUUIDAndUserIdLogic {
	return &GetOrderInfoByUUIDAndUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderInfoByUUIDAndUserId
//
//	@Description: 使用订单id和用户id查询订单
//	@receiver l
//	@param in
//	@return *pb.GetOrderInfoByUUIDAndUserIDResp
//	@return error
func (l *GetOrderInfoByUUIDAndUserIdLogic) GetOrderInfoByUUIDAndUserId(in *pb.GetOrderInfoByUUIDAndUserIDReq) (*pb.GetOrderInfoByUUIDAndUserIDResp, error) {
	if in.GetUuid() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	timeout, cancelFunc := context.WithTimeout(context.Background(), utils.GetContextDuration())
	defer cancelFunc()
	order := l.svcCtx.Query.Order
	data, err := order.WithContext(timeout).Where(order.UUID.Eq(in.GetUuid()), order.UserID.Eq(in.GetUserId())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetOrderInfoByUUIDAndUserIDResp{}, nil
	}
	if err != nil {
		return nil, err
	}
	return &pb.GetOrderInfoByUUIDAndUserIDResp{
		Order: &pb.Order{
			Id:              data.ID,
			Uuid:            data.UUID,
			ProductId:       data.ProductID,
			SysModel:        data.SysModel,
			Status:          data.Status,
			UserId:          data.UserID,
			PayPrice:        *data.PayPrice,
			PayPath:         data.PayPath,
			PayPathOrderNum: data.PayPathOrderNum,
			CreateTime:      data.CreateTime.Unix(),
			UpdateTime:      data.UpdateTime.Unix(),
		},
	}, nil
}
