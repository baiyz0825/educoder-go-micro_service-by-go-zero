package textRes

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTextResLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTextResLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTextResLogic {
	return &SearchTextResLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SearchTextRes
//
//	@Description: 条件查询所有在线文本
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *SearchTextResLogic) SearchTextRes(req *types.SearchOnlineConditionTextReq) (resp *types.SearchOnlineTextConditionResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	textReq := &pb.SearchOnlineConditionTextReq{
		Page:       req.Page,
		Limit:      req.Limit,
		Owner:      req.Owner,
		ClassID:    req.ClassID,
		Permission: req.Permission,
	}

	searchData, err := l.svcCtx.ResourcesRpc.SearchOnlineConditionText(l.ctx, textReq)
	if err != nil {
		return nil, err
	}
	resp = &types.SearchOnlineTextConditionResp{}
	for _, text := range searchData.OnlineText {
		temp := types.OnlineText{}
		err := copier.Copy(&temp, text)
		if err != nil {
			l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
			return nil, xerr.NewErrMsg("数据处理异常")
		}
		resp.OnlineText = append(resp.OnlineText, temp)
	}
	return resp, nil
}
