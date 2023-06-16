package logic

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOrderByConditionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOrderByConditionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOrderByConditionLogic {
	return &SearchOrderByConditionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchOrderByCondition
//
//	@Description: 条件检索用户订单信息
//	@receiver l
//	@param in
//	@return *pb.SearchOrderByConditionResp
//	@return error
func (l *SearchOrderByConditionLogic) SearchOrderByCondition(in *pb.SearchOrderByConditionReq) (*pb.SearchOrderByConditionResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check param
	if in.GetProductId() == 0 && in.GetUserId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// search condition
	q := l.svcCtx.Query.Order
	query := q.WithContext(l.ctx)
	var condition []gen.Condition
	condition = append(condition, q.SysModel.Eq(in.GetSysModel()), q.PayPath.Eq(in.GetPayPath()))
	// 订单状态非为创建支付订单
	if in.GetStatus() != 0 {
		condition = append(condition, q.Status.Eq(in.GetStatus()))
	} else {
		// 默认状态
		condition = append(condition, q.Status.In(xconst.ORDER_STATUS_PAYING, xconst.ORDER_STATUS_PAYED))
	}
	if in.GetUserId() != 0 {
		condition = append(condition, q.UserID.Eq(in.GetUserId()))
	}
	if in.GetProductId() != 0 {
		condition = append(condition, q.ProductID.Eq(in.GetProductId()))
	}
	datas, total, err := query.Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
	if err == gorm.ErrRecordNotFound {
		return &pb.SearchOrderByConditionResp{}, nil
	}
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// copy data
	var respData []*pb.Order
	for _, data := range datas {
		temp := &pb.Order{
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
			temp.PayCodeURL = *data.PayCodeURL
		}
		respData = append(respData, temp)
	}
	return &pb.SearchOrderByConditionResp{
		Order: respData,
		Total: total,
	}, nil
}
