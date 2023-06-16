package logic

import (
	"context"
	"errors"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	if in == nil {
		return nil, errors.New("pb输入为空")
	}
	// 查询映射数据表
	u := l.svcCtx.Query.User
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	user, err := u.WithContext(deadlineCtx).Where(u.UID.Eq(in.GetID())).First()
	if user == nil {
		return &pb.GetUserByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
		return nil, errors.New("数据不存在")
	}
	data := &pb.User{}
	err = copier.Copy(data, user)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("复制db -> pb错误")
		return nil, errors.New("数据转化错误")
	}
	// 手动拷贝不兼容字段
	data.CreateTime = user.CreateTime.UnixMilli()
	data.UpdateTime = user.UpdateTime.UnixMilli()
	// return
	return &pb.GetUserByIdResp{
		User: data,
	}, nil
}
