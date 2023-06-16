package order

import (
	"context"
	"encoding/json"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type DoOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoOrderLogic {
	return &DoOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DoOrder
//
//	@Description: 用户下单
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *DoOrderLogic) DoOrder(req *types.DoOrderReq) (resp *types.DoOrderResp, err error) {
	// 参数检查
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	deadline, cancelFunc := context.WithTimeout(context.Background(), utils.GetContextDuration())
	defer cancelFunc()
	uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.SERVER_ERROR)
	}
	createOrderData, err := l.svcCtx.OrderRpc.DoOrder(deadline, &pb.DoOrderReq{
		ProductId: req.ProductId,
		UserId:    uid,
		PayPath:   req.PayPath,
	})
	if err != nil {
		return nil, err
	}
	return &types.DoOrderResp{
		PayPathOrderNum: createOrderData.PayPathOrderNum,
		Status:          createOrderData.Status,
		PayUrl:          createOrderData.PayUrl,
	}, nil
}
