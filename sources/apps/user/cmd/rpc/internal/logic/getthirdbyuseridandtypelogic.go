package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetThirdByUserIdAndTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetThirdByUserIdAndTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetThirdByUserIdAndTypeLogic {
	return &GetThirdByUserIdAndTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetThirdByUserIdAndType
//
//	@Description: 使用userId以及相应的type精准查询
//	@receiver l
//	@param in
//	@return *pb.PagesThirdResp
//	@return error
func (l *GetThirdByUserIdAndTypeLogic) GetThirdByUserIdAndType(in *pb.GetThirdByUserIdAndTypeReq) (*pb.GetThirdByUserIdAndTypeResp, error) {
	// 使用userId 和 type查询密钥
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 必要参数检查
	if in.GetUserID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// 精准查询
	acc := l.svcCtx.Query.UserAcc
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	first, err := acc.WithContext(deadlineCtx).Where(acc.UserID.Eq(in.GetUserID()), acc.Type.Eq(in.GetType())).First()
	if first == nil {
		return &pb.GetThirdByUserIdAndTypeResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	data := &pb.Third{}
	if err := copier.Copy(data, first); err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	return &pb.GetThirdByUserIdAndTypeResp{
		Third: []*pb.Third{data},
	}, nil
}
