package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetThirdDataByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetThirdDataByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetThirdDataByIdLogic {
	return &GetThirdDataByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetThirdDataByIdLogic) GetThirdDataById(in *pb.GetThirdDataByIdReq) (*pb.GetThirdDataByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// search db
	third := l.svcCtx.Query.ThirdData
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	first, err := third.WithContext(deadlineCtx).Where(third.ID.Eq(in.GetID())).First()
	if first == nil {
		// return
		return &pb.GetThirdDataByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// copy
	p := &pb.ThirdData{}
	err = copier.Copy(p, first)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("复制db -> pb错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// return
	return &pb.GetThirdDataByIdResp{
		ThirdData: p,
	}, nil
}
