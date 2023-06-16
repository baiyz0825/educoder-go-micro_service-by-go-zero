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

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 必要检查
	if in.GetUID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// update to db
	u := l.svcCtx.Query.User
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	user := &model.User{}
	if len(in.GetAvatar()) > 0 {
		_, err := u.WithContext(deadlineCtx).Where(u.UID.Eq(in.GetUID())).Update(u.Avatar, in.GetAvatar())
		if err != nil {
			l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_UPDATE_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
		}
	} else {
		if err := copier.Copy(user, in); err != nil {
			return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
		}
		// 补齐uid
		_, err := u.WithContext(deadlineCtx).Where(u.UID.Eq(in.GetUID())).Omit(u.Avatar).Updates(user)
		if err != nil {
			l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_UPDATE_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
		}
	}
	return &pb.UpdateUserResp{}, nil
}
