package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchClassificationAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchClassificationAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchClassificationAllLogic {
	return &SearchClassificationAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchClassificationAll
//
//	@Description: 查询全部分类
//	@receiver l
//	@param in
//	@return *pb.SearchClassificationAllResp
//	@return error
func (l *SearchClassificationAllLogic) SearchClassificationAll(in *pb.SearchClassificationAllReq) (*pb.SearchClassificationAllResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}

	// search
	finds, err := l.svcCtx.Query.Classification.WithContext(l.ctx).Find()
	if err == gorm.ErrRecordNotFound {
		l.Logger.WithFields(logx.Field("error:", err)).Error("记录不存在")
		return &pb.SearchClassificationAllResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// tree
	tree := getTreeClassification(finds, 0)
	// return
	return &pb.SearchClassificationAllResp{
		Classifications: tree,
	}, nil
}

// getTreeClassification
//
//	@Description: 获取子分类所有树形菜单
//	@param sources
//	@param parentId
//	@return []*pb.ClassificationTreeMenu
func getTreeClassification(sources []*model.Classification, parentId int64) []*pb.ClassificationTreeMenu {
	var result []*pb.ClassificationTreeMenu
	for _, menu := range sources {
		if menu.ClassParentID == parentId {
			children := getTreeClassification(sources, menu.ClassID)
			tempRootNode := &pb.ClassificationTreeMenu{
				ClassID:          menu.ClassID,
				ClassParentID:    menu.ClassParentID,
				ClassName:        *menu.ClassName,
				ClassResourceNum: *menu.ClassResourceNum,
				Children:         children,
			}
			result = append(result, tempRootNode)
		}
	}
	return result
}
