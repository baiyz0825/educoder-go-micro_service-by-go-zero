package classification

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllClassificationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllClassificationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllClassificationsLogic {
	return &GetAllClassificationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAllClassifications
//
//	@Description: 获取所有分类数据
//	@receiver l
//	@return resp
//	@return err
func (l *GetAllClassificationsLogic) GetAllClassifications() (resp *types.ClassificationTreeMenuResp, err error) {
	ctx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	in := &pb.SearchClassificationAllReq{}
	// rpc调用
	classificationAll, err := l.svcCtx.ResourcesRpc.SearchClassificationAll(ctx, in)
	if err != nil {
		return nil, err
	}
	resp = &types.ClassificationTreeMenuResp{}
	// 拷贝数据
	err = copier.Copy(resp, classificationAll)
	// 创建默认选项
	defaults := types.MenuItem{
		ClassID:          -1,
		ClassParentID:    -1,
		ClassName:        "请选择分类",
		ClassResourceNum: 0,
		Children:         nil,
	}
	resp.Classifications = append(resp.Classifications, defaults)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
