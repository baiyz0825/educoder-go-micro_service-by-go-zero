package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetThirdBindDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetThirdBindDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetThirdBindDataLogic {
	return &GetThirdBindDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type ThirdBindData struct {
	Name    string
	Type    int64
	ThirdID int64
}

func (l *GetThirdBindDataLogic) GetThirdBindData(in *pb.GetThirdBindDataReq) (*pb.GetThirdBindDataResp, error) {
	// check pb
	if in == nil || in.GetUserID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// 查询third表获取third id
	acc := l.svcCtx.Query.UserAcc
	thirds, err := acc.WithContext(ctx).Select(acc.ID).Where(acc.UserID.Eq(in.UserID)).Find()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	var ids []int64
	for _, data := range thirds {
		ids = append(ids, data.ID)
	}
	var resultData []ThirdBindData
	// 查询third_data表获取用户名
	thirdDataQ := l.svcCtx.Query.ThirdData
	err = acc.WithContext(ctx).
		Select(thirdDataQ.ThirdID, thirdDataQ.Name, acc.Type).
		LeftJoin(thirdDataQ, thirdDataQ.ThirdID.EqCol(acc.ID)).Where(acc.UserID.In(ids...)).Scan(resultData)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.DB_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	var respThirdBind []*pb.ThirdBind
	for _, res := range resultData {
		temp := &pb.ThirdBind{
			ThirdId:   res.ThirdID,
			ThirdName: res.Name,
			ThirdType: res.Type,
		}
		respThirdBind = append(respThirdBind, temp)
	}
	return &pb.GetThirdBindDataResp{
		Thirds: respThirdBind,
	}, nil
}
