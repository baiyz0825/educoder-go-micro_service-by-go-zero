package major

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllMajorsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllMajorsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllMajorsLogic {
	return &GetAllMajorsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAllMajors
//
//	@Description: 获取全部主修课程
//	@receiver l
//	@return resp
//	@return err
func (l *GetAllMajorsLogic) GetAllMajors() (resp *types.GetAllMajorsResp, err error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), utils.GetContextDuration())
	defer cancelFunc()
	pages, err := l.svcCtx.UserRpc.GetMajorPages(timeout, &pb.GetMajorPagesReq{
		Page:  0,
		Limit: 0,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.SERVER_ERROR)
	}
	resp = &types.GetAllMajorsResp{}
	for _, major := range pages.Major {
		temp := types.Major{}
		err := copier.Copy(&temp, major)
		if err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).Error("拷贝错误")
			return nil, xerr.NewErrCode(xerr.SERVER_ERROR)
		}
		resp.Major = append(resp.Major, temp)
	}
	return resp, nil
}
