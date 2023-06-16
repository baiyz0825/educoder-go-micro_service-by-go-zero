package deplete

import (
	"context"
	"encoding/json"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileAndSpaceInsightLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileAndSpaceInsightLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileAndSpaceInsightLogic {
	return &GetFileAndSpaceInsightLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetFileAndSpaceInsight
//
//	@Description: 查询用户空间以及文件使用情况
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GetFileAndSpaceInsightLogic) GetFileAndSpaceInsight() (resp *types.GetCountUiDResp, err error) {
	uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户信息失败")
	}
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// rpc
	insightData, err := l.svcCtx.ResourcesRpc.GetCountByUId(deadlineCtx, &pb.GetCountByUIdReq{Uid: uid})
	if err != nil {
		return nil, err
	}
	resp = &types.GetCountUiDResp{}
	err = copier.Copy(&resp.UserFileCount, insightData.GetCount())
	if err != nil {
		l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
		return nil, xerr.NewErrMsg("数据处理异常")
	}

	return resp, nil
}
