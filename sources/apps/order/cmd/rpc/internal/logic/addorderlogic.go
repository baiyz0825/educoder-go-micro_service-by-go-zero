package logic

import (
	"context"
	"fmt"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------订单表-----------------------

// AddOrder
//
//	@Description: 插入一个订单
//	@receiver l
//	@param in
//	@return *pb.AddOrderResp
//	@return error
func (l *AddOrderLogic) AddOrder(in *pb.AddOrderReq) (*pb.AddOrderResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// param check
	if in.GetProductId() == 0 || in.GetUserId() == 0 || in.GetPayPrice() == 0 || len(in.GetPayPathOrderNum()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// gen uuid && db data
	id, err := utils.GenSnowFlakeId()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("生成订单流水号失败")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	orderQ := l.svcCtx.Query.Order
	// create data
	order := &model.Order{
		UUID:            id,
		ProductID:       in.GetProductId(),
		SysModel:        in.GetSysModel(),
		Status:          0,
		UserID:          in.GetUserId(),
		PayPrice:        &in.PayPrice,
		PayPath:         in.GetPayPath(),
		PayPathOrderNum: in.GetPayPathOrderNum(),
	}
	err = orderQ.WithContext(l.ctx).Create(order)
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).Error(fmt.Sprintf("生成新订单失败，数据%v", order))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	// 查询生成的数据
	orderGenData, err := orderQ.WithContext(l.ctx).Where(orderQ.UUID.Eq(id)).First()
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).Error(fmt.Sprintf("查询生成新订单数据失败，订单uuid：%v", id))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	return &pb.AddOrderResp{
		Uuid:            orderGenData.UUID,
		Id:              orderGenData.ID,
		Status:          orderGenData.Status,
		PayPath:         orderGenData.PayPath,
		PayPathOrderNum: orderGenData.PayPathOrderNum,
	}, nil
}
