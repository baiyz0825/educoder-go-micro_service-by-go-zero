package fileRes

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFileResLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchFileResLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFileResLogic {
	return &SearchFileResLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SearchFileRes
//
//	@Description: 查询文件资源详情
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *SearchFileResLogic) SearchFileRes(req *types.SearchFileConditionReq) (resp *types.SearchFileConditionResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	reqPb := &pb.SearchFileConditionReq{}
	err = copier.Copy(reqPb, req)
	// 查询非用户隐藏文件资源
	reqPb.Status = 3
	// 禁止后缀搜索
	reqPb.Suffix = ""
	reqPb.Type = req.FileType
	if err != nil {
		return nil, err
	}
	// rpc
	pages, err := l.svcCtx.ResourcesRpc.SearchFileConditionPages(deadlineCtx, reqPb)
	if err != nil {
		return nil, err
	}
	resp = &types.SearchFileConditionResp{}
	for _, data := range pages.GetFile() {
		temp := types.File{}
		err := copier.Copy(&temp, data)
		if err != nil {
			l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
			return nil, xerr.NewErrMsg("数据处理异常")
		}
		temp.FileType = data.Type
		resp.Files = append(resp.Files, temp)
	}
	return resp, nil
}
