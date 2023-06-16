package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddThirdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddThirdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddThirdLogic {
	return &AddThirdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户三方关联数据-----------------------m

func (l *AddThirdLogic) AddThird(in *pb.AddThirdReq) (*pb.AddThirdResp, error) {
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	data := &model.UserAcc{}
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// check empty
	if len(in.AccessToken) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	err := copier.Copy(data, in)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("复制db -> pb错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// insert db
	if err := l.svcCtx.Query.UserAcc.WithContext(ctx).Create(data); err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	return &pb.AddThirdResp{}, nil
}
