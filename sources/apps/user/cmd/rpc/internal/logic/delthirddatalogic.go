package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelThirdDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelThirdDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelThirdDataLogic {
	return &DelThirdDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelThirdDataLogic) DelThirdData(in *pb.DelThirdDataReq) (*pb.DelThirdDataResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// delete db
	thirdData := l.svcCtx.Query.ThirdData
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	if _, err := thirdData.WithContext(ctx).Where(thirdData.ID.Eq(in.ID)).Delete(); err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_DELETE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_DELETE_ERR)
	}
	return &pb.DelThirdDataResp{}, nil
}
