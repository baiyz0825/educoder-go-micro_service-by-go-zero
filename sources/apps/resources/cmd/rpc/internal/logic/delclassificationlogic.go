package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelClassificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelClassificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelClassificationLogic {
	return &DelClassificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DelClassification
//
//	@Description: 递归删除所有父分类下子分类信息
//	@receiver l
//	@param in
//	@return *pb.DelClassificationResp
//	@return error
func (l *DelClassificationLogic) DelClassification(in *pb.DelClassificationReq) (*pb.DelClassificationResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// delete
	parentId := in.GetID()
	// 查询全部菜单数据
	class := l.svcCtx.Query.Classification
	find, err := class.WithContext(l.ctx).Select(class.ClassID, class.ClassParentID).Order(class.ClassID).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		logx.WithContext(l.ctx).WithFields(logx.Field("error:", err)).Error("递归查询分类数据失败")
		return nil, err
	}
	var deletedIds []int64
	// 进行递归获取id
	recursiveIds(find, parentId, &deletedIds)
	// 删除数据
	_, err = class.WithContext(l.ctx).Where(class.ClassID.In(deletedIds...)).Delete()
	if err != nil {
		logx.WithContext(l.ctx).WithFields(logx.Field("error:", err)).Error("批量删除错误")
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_DELETE_ERR))
	}
	return &pb.DelClassificationResp{}, nil
}

// recursiveIds
//
//	@Description: 递归获取所有id
//	@param sources
//	@param parentId
//	@param resId
func recursiveIds(sources []*model.Classification, parentId int64, resId *[]int64) {
	// 加入当前子节点id
	*resId = append(*resId, parentId)
	// 获取子节点
	node := getChildrenNode(sources, parentId)
	// 出口
	if len(node) == 0 {
		return
	}
	for _, value := range node {
		// 递归获取子节点
		recursiveIds(sources, value, resId)
	}
}

// getChildrenNode
//
//	@Description: 获取子节点id
//	@param sources
//	@param parentId
//	@return []int64
func getChildrenNode(sources []*model.Classification, parentId int64) []int64 {
	var childrenNodesID []int64
	for _, value := range sources {
		if value.ClassParentID == parentId {
			childrenNodesID = append(childrenNodesID, value.ClassID)
		}
	}
	return childrenNodesID
}
