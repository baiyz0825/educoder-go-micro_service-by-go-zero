package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateThirdDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateThirdDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateThirdDataLogic {
	return &UpdateThirdDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateThirdDataLogic) UpdateThirdData(in *pb.UpdateThirdDataReq) (*pb.UpdateThirdDataResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 必要检查
	if in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// update to db
	// update db
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	t := l.svcCtx.Query.ThirdData
	third := &model.ThirdData{}
	if err := copier.Copy(third, in); err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	_, err := t.WithContext(deadlineCtx).Where(t.ID.Eq(in.GetID())).Updates(third)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}
	return &pb.UpdateThirdDataResp{}, nil
}
