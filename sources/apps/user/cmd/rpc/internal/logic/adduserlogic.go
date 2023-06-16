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

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户数据基本信息表-----------------------

func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.AddUserResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check data username password phone number
	if len(in.GetName()) == 0 || len(in.GetPassword()) == 0 || len(in.GetPhone()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// insert db
	user := &model.User{}
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// copier  data from in to user
	err := copier.Copy(user, in)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("复制db -> pb错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// set uuid
	id, err := utils.GenSnowFlakeId()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("生成系统id错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	user.UniqueID = id
	if err := l.svcCtx.Query.User.WithContext(ctx).Create(user); err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	return &pb.AddUserResp{}, nil
}
