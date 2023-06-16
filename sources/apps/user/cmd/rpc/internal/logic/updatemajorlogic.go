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

type UpdateMajorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMajorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMajorLogic {
	return &UpdateMajorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMajor
//
//	@Description: 更新分类信息
//	@receiver l
//	@param in
//	@return *pb.UpdateMajorResp
//	@return error
func (l *UpdateMajorLogic) UpdateMajor(in *pb.UpdateMajorReq) (*pb.UpdateMajorResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 必要检查
	if in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// update db
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	m := l.svcCtx.Query.Major
	major := &model.Major{}
	if err := copier.Copy(major, in); err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	_, err := m.WithContext(deadlineCtx).Where(m.ID.Eq(in.GetID())).Omit(m.ID).Updates(major)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}
	return &pb.UpdateMajorResp{}, nil
}
