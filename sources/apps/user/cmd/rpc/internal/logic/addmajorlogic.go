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

type AddMajorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMajorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMajorLogic {
	return &AddMajorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户专业统计表-----------------------

func (l *AddMajorLogic) AddMajor(in *pb.AddMajorReq) (*pb.AddMajorResp, error) {
	// 参数校验
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if len(in.GetName()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)

	}
	// 插入数据库
	insertData := &model.Major{}
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	err := copier.Copy(insertData, in)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("复制错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	err = l.svcCtx.Query.Major.WithContext(ctx).Create(insertData)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	// 返回数据
	return &pb.AddMajorResp{}, nil
}
