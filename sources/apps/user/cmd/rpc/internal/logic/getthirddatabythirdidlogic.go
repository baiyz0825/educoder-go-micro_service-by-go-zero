package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetThirdDataByThirdIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetThirdDataByThirdIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetThirdDataByThirdIDLogic {
	return &GetThirdDataByThirdIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetThirdDataByThirdId
//
//	@Description: 通过用户ids获取三方数据
//	@receiver l
//	@param in
//	@return *pb.GetThirdDataByIdsResp
//	@return error
func (l *GetThirdDataByThirdIDLogic) GetThirdDataByThirdId(in *pb.GetThirdDataByThirdIdReq) (*pb.GetThirdDataByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 必要参数检测
	if in.GetThirdID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// 用户id 搜索三方数据信息
	q := l.svcCtx.Query.ThirdData
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	data, err := q.WithContext(deadlineCtx).Where(q.ThirdID.Eq(in.GetThirdID())).First()
	if data == nil {
		// return
		return &pb.GetThirdDataByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	thirdData := &pb.ThirdData{}
	if err := copier.Copy(thirdData, data); err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	return &pb.GetThirdDataByIdResp{
		ThirdData: thirdData,
	}, nil
}
