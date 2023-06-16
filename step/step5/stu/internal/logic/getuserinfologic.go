package logic

import (
	"context"
	"errors"

	"stu/internal/svc"
	"stu/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.UserReq) (*pb.UserResp, error) {
	if in == nil {
		return nil, errors.New("不允许传递空数据")
	}
	var userMoneyCardNum []int64
	for i := 0; i < 3; i++ {
		userMoneyCardNum = append(userMoneyCardNum, in.GetPasswd())
	}
	return &pb.UserResp{
		UserHome:         in.UserName,
		UserMoneyCardNum: userMoneyCardNum,
	}, nil
}
