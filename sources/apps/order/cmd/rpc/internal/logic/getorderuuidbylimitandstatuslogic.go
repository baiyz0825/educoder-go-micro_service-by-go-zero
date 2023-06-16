package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderUUIdByLimitAndStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderUUIdByLimitAndStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderUUIdByLimitAndStatusLogic {
	return &GetOrderUUIdByLimitAndStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderUUIdByLimitAndStatus
//
//	@Description: 按照订单状态以及数量限制批量获取订单UUID
//	@receiver l
//	@param in
//	@return *pb.GetOrderUUIdByLimitAndStatusResp
//	@return error
func (l *GetOrderUUIdByLimitAndStatusLogic) GetOrderUUIdByLimitAndStatus(in *pb.GetOrderUUIdByLimitAndStatusReq) (*pb.GetOrderUUIdByLimitAndStatusResp, error) {
	if in.Limit == 0 {
		return &pb.GetOrderUUIdByLimitAndStatusResp{}, nil
	}
	// 查询数据库
	Q := l.svcCtx.Query.Order
	deadline, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	find, err := Q.WithContext(deadline).Select(Q.UUID, Q.ID).Where(Q.Status.Eq(in.GetStatus())).Order(Q.CreateTime).Limit(int(in.GetLimit())).Find()
	if err != nil {
		return nil, err
	}
	var respUUIDList []int64
	for _, data := range find {
		respUUIDList = append(respUUIDList, data.UUID)
	}
	return &pb.GetOrderUUIdByLimitAndStatusResp{
		OrderUUid: respUUIDList,
	}, nil
}
