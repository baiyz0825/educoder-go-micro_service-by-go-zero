package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIdLogic) GetOrderById(in *pb.GetOrderByIdReq) (*pb.GetOrderByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check id
	if in.GetId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	order := l.svcCtx.Query.Order
	data, err := order.WithContext(l.ctx).Where(order.ID.Eq(in.GetId())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetOrderByIdResp{}, nil
	}
	if err != nil {
		return nil, err
	}
	p := &pb.Order{
		Id:              data.ID,
		Uuid:            data.UUID,
		ProductId:       data.ProductID,
		SysModel:        data.SysModel,
		Status:          data.Status,
		UserId:          data.UserID,
		PayPrice:        *data.PayPrice,
		PayPath:         data.PayPath,
		PayPathOrderNum: data.PayPathOrderNum,
		CreateTime:      data.CreateTime.UnixMilli(),
		UpdateTime:      data.UpdateTime.UnixMilli(),
	}
	if data.PayCodeURL != nil {
		p.PayCodeURL = *data.PayCodeURL
	}
	return &pb.GetOrderByIdResp{
		Order: p,
	}, nil
}
