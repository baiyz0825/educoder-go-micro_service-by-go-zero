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

type GetMajorPagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMajorPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMajorPagesLogic {
	return &GetMajorPagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMajorPages
//
//	@Description: 分页查询所有的主修情况
//	@receiver l
//	@param in
//	@return *pb.SearchMajorResp
//	@return error
func (l *GetMajorPagesLogic) GetMajorPages(in *pb.GetMajorPagesReq) (*pb.GetMajorPagesResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 必要参数检测
	// 负分页
	if in.GetPage() < 0 || in.GetLimit() < 0 {
		in.Page = 0
		in.Limit = 0
	}
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// 使用检索 分页 分页于limit 为默认相当于查询全部
	if in.GetPage() != 0 && in.GetLimit() != 0 {
		// pages
		pageData, _, err := l.svcCtx.Query.Major.WithContext(ctx).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
		if err != nil {
			l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}
		var respData []*pb.Major
		for _, value := range pageData {
			temp := &pb.Major{}
			if err := copier.Copy(temp, value); err != nil {
				return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
			}
			respData = append(respData, temp)
		}
		return &pb.GetMajorPagesResp{Major: respData}, nil
	}

	// all
	find, err := l.svcCtx.Query.Major.WithContext(ctx).Find()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	var respData []*pb.Major
	for _, value := range find {
		temp := &pb.Major{}
		if err := copier.Copy(temp, value); err != nil {
			return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
		}
		respData = append(respData, temp)
	}
	return &pb.GetMajorPagesResp{Major: respData}, nil
}
