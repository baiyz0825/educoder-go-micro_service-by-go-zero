package logic

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFileConditionPagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFileConditionPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFileConditionPagesLogic {
	return &SearchFileConditionPagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchFileConditionPages
//
//	@Description: 多条件查询（用户id / 全部文件中筛选）
//	@receiver l
//	@param in
//	@return *pb.SearchFileConditionResp
//	@return error
func (l *SearchFileConditionPagesLogic) SearchFileConditionPages(in *pb.SearchFileConditionReq) (*pb.SearchFileConditionResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	q := l.svcCtx.Query.File
	query := q.WithContext(l.ctx)
	var condition []gen.Condition
	// default condition
	condition = append(condition, q.Status.Neq(in.GetStatus()), q.Type.Eq(in.GetType()))
	// conditions
	if in.GetClass() != 0 {
		condition = append(condition, q.Class.Eq(in.GetClass()))
	}

	// user id exist
	if in.GetOwner() != 0 {
		// user condition
		condition = append(condition, q.Owner.Eq(in.GetOwner()))
	}
	// file name
	if len(in.GetName()) != 0 {
		condition = append(condition, q.Name.Eq(in.GetName()))
	}
	// file suffix
	if len(in.GetSuffix()) != 0 {
		condition = append(condition, q.Suffix.Eq(in.GetSuffix()))
	}
	// pages
	datas, _, err := query.Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
	if err == gorm.ErrRecordNotFound {
		return &pb.SearchFileConditionResp{}, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	// copy
	var respData []*pb.File
	for _, data := range datas {
		temp := &pb.File{}
		err := copier.Copy(temp, data)
		if err != nil {
			return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
		}
		temp.UpdateTime = data.UpdateTime.UnixMilli()
		temp.CreateTime = data.CreateTime.UnixMilli()
		respData = append(respData, temp)
	}

	return &pb.SearchFileConditionResp{
		File: respData,
	}, nil
}
