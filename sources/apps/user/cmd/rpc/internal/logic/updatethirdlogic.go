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

type UpdateThirdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateThirdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateThirdLogic {
	return &UpdateThirdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateThirdLogic) UpdateThird(in *pb.UpdateThirdReq) (*pb.UpdateThirdResp, error) {
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
	acc := l.svcCtx.Query.UserAcc
	accData := &model.UserAcc{}
	if err := copier.Copy(accData, in); err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	_, err := acc.WithContext(deadlineCtx).Where(acc.ID.Eq(in.GetID())).Updates(accData)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}
	return &pb.UpdateThirdResp{}, nil
}
