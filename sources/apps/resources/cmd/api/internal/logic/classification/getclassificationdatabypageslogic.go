package classification

import (
	"context"
	"encoding/json"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassificationDataByPagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassificationDataByPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassificationDataByPagesLogic {
	return &GetClassificationDataByPagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetClassificationDataByPages
//
//	@Description: 获取分类下的资源数据
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GetClassificationDataByPagesLogic) GetClassificationDataByPages(req *types.SearchClassificationSubDataReq) (resp *types.SearchClassificationSubDataResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	pagesReq := &pb.ClassificationDataByPagesReq{
		Page:             req.Page,
		Limit:            req.Limit,
		ClassificationID: req.ClassificationID,
	}
	uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 查询用户自己的
	if req.IsUser {
		pagesReq.UserId = uid
	}
	// 需要查询指定分类
	if req.ResType != 0 {
		pagesReq.ResType = req.ResType
	}
	// 关键词名称查询
	if len(strings.Trim(req.KeyWord, " ")) >= 0 {
		pagesReq.KeyWord = strings.Trim(req.KeyWord, " ")
	}
	// 创建context
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	pages, err := l.svcCtx.ResourcesRpc.GetClassificationDataByPages(deadlineCtx, pagesReq)
	if err != nil {
		return nil, err
	}
	files := pages.GetFiles()
	onlineText := pages.GetOnlineText()
	wait := errgroup.Group{}
	// 处理文件
	resp = &types.SearchClassificationSubDataResp{}
	wait.Go(func() error {
		for _, file := range files {
			temp := types.File{}
			err := copier.Copy(&temp, file)
			// 补齐不一致类型值
			temp.FileType = file.Type
			if err != nil {
				return err
			}
			resp.Files = append(resp.Files, temp)
		}
		resp.FilesTotal = pages.FilesTotal
		return nil
	})

	// 处理在线文档
	wait.Go(func() error {
		for _, text := range onlineText {
			temp := types.OnlineText{}
			err := copier.Copy(&temp, text)
			if err != nil {
				return err
			}
			resp.OnlineText = append(resp.OnlineText, temp)
		}
		resp.TextsTotal = pages.TextsTotal
		return nil
	})
	err = wait.Wait()
	if err != nil {
		l.Logger.WithFields(logx.Field("err: ", err)).Error("分页获取分类资源，发生copier错误")
		return nil, xerr.NewErrMsg("数据处理异常")
	}
	return resp, nil
}
