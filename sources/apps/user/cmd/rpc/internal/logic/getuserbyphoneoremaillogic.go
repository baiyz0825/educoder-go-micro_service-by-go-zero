package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByPhoneOrEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByPhoneOrEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByPhoneOrEmailLogic {
	return &GetUserByPhoneOrEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserByPhoneOrEmail
//
//	@Description: 使用手机号或者邮箱查询用户 传哪个用哪个
//	@receiver l
//	@param in
//	@return *pb.GetUserByIdResp
//	@return error
func (l *GetUserByPhoneOrEmailLogic) GetUserByPhoneOrEmail(in *pb.GetUserByPhoneOrEmailReq) (*pb.GetUserByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if len(in.GetPhone()) == 0 && len(in.GetEmail()) == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	q := l.svcCtx.Query.User
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	if len(in.GetPhone()) != 0 && len(in.GetEmail()) != 0 {
		// 精准查询
		user, err := q.WithContext(deadlineCtx).Where(q.Phone.Eq(in.GetPhone()), q.Email.Eq(in.GetEmail())).First()
		if err == gorm.ErrRecordNotFound {
			return &pb.GetUserByIdResp{}, nil
		}
		if err != nil {
			l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}
		data := &pb.User{}
		if err := copier.Copy(data, user); err != nil {
			return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
		}
		// 手动拷贝不兼容字段
		data.CreateTime = user.CreateTime.UnixMilli()
		data.UpdateTime = user.UpdateTime.UnixMilli()
		return &pb.GetUserByIdResp{
			User: data,
		}, nil
	} else if len(in.GetPhone()) != 0 {
		// 手机号检索
		user, err := q.WithContext(deadlineCtx).Where(q.Phone.Eq(in.GetPhone())).First()
		if err == gorm.ErrRecordNotFound {
			return &pb.GetUserByIdResp{}, nil
		}
		if err != nil {
			l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}
		data := &pb.User{}
		if err := copier.Copy(data, user); err != nil {
			return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
		}
		// 手动拷贝不兼容字段
		data.CreateTime = user.CreateTime.UnixMilli()
		data.UpdateTime = user.UpdateTime.UnixMilli()
		return &pb.GetUserByIdResp{
			User: data,
		}, nil
	} else {
		// 邮箱检索
		user, err := q.WithContext(deadlineCtx).Where(q.Email.Eq(in.GetEmail())).First()
		if err == gorm.ErrRecordNotFound {
			return &pb.GetUserByIdResp{}, nil
		}
		if err != nil {
			l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}
		data := &pb.User{}
		if err := copier.Copy(data, user); err != nil {
			return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
		}
		// 手动拷贝不兼容字段
		data.CreateTime = user.CreateTime.UnixMilli()
		data.UpdateTime = user.UpdateTime.UnixMilli()
		return &pb.GetUserByIdResp{
			User: data,
		}, nil
	}
}
