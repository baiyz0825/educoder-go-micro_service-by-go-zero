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

type GetUserEarnByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserEarnByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserEarnByIdLogic {
	return &GetUserEarnByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserEarnByIdLogic) GetUserEarnById(in *pb.GetUserEarnByIdReq) (*pb.GetUserEarnByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check id
	if in.GetId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	order := l.svcCtx.Query.UserEarn
	data, err := order.WithContext(l.ctx).Where(order.ID.Eq(in.GetId())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetUserEarnByIdResp{}, nil
	}
	if err != nil {
		return nil, err
	}
	respData := &pb.UserEarn{}
	err = copier.Copy(respData, data)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}

	return &pb.GetUserEarnByIdResp{
		UserEarn: respData,
	}, nil
}
