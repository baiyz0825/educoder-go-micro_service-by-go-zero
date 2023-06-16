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

type SearchOnlineConditionTextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnlineConditionTextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnlineConditionTextLogic {
	return &SearchOnlineConditionTextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchOnlineConditionText
//
//	@Description: 多条件查询（大类，用户id / 全部文件中筛选）
//	@receiver l
//	@param in
//	@return *pb.SearchOnlineTextConditionResp
//	@return error
func (l *SearchOnlineConditionTextLogic) SearchOnlineConditionText(in *pb.SearchOnlineConditionTextReq) (*pb.SearchOnlineTextConditionResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	q := l.svcCtx.Query.OnlineText
	query := q.WithContext(l.ctx)
	var condition []gen.Condition
	// default condition
	condition = append(condition, q.Permission.Eq(in.GetPermission()))
	// conditions
	if in.GetClassID() != 0 {
		condition = append(condition, q.ClassID.Eq(in.GetClassID()))
	}

	// user id exist
	if in.GetOwner() != 0 {
		// user condition
		condition = append(condition, q.Owner.Eq(in.GetOwner()))
	}
	// pages
	datas, _, err := query.Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
	if err == gorm.ErrRecordNotFound {
		return &pb.SearchOnlineTextConditionResp{}, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	// copy
	var respData []*pb.OnlineText
	for _, data := range datas {
		temp := &pb.OnlineText{}
		err := copier.Copy(temp, data)
		if err != nil {
			return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
		}
		temp.UpdateTime = data.UpdateTime.UnixMilli()
		temp.CreateTime = data.CreateTime.UnixMilli()
		respData = append(respData, temp)
	}

	return &pb.SearchOnlineTextConditionResp{
		OnlineText: respData,
	}, nil
}
