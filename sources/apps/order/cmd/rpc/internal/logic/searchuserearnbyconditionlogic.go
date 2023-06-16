package logic

import (
	"context"
	"time"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserEarnByConditionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserEarnByConditionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserEarnByConditionLogic {
	return &SearchUserEarnByConditionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserEarnByConditionLogic) SearchUserEarnByCondition(in *pb.SearchUserEarnByConditionReq) (*pb.SearchUserEarnByConditionResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// 分页加时间范围查询
	q := l.svcCtx.Query.UserEarn
	query := q.WithContext(l.ctx)
	var condition []gen.Condition
	condition = append(condition, q.UpdateTime.Between(time.UnixMilli(in.GetFromTime()), time.UnixMilli(in.ToTime)))
	if in.GetUserId() != 0 {
		condition = append(condition, q.UserID.Eq(in.GetUserId()))
	}
	datas, _, err := query.Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
	if err == gorm.ErrRecordNotFound {
		return &pb.SearchUserEarnByConditionResp{}, nil
	}
	if err != nil {
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	respData := make([]*pb.UserEarn, len(datas))
	for _, data := range datas {
		temp := &pb.UserEarn{}
		err := copier.Copy(temp, data)
		if err != nil {
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}
		respData = append(respData, temp)
	}
	return &pb.SearchUserEarnByConditionResp{
		UserEarn: respData,
	}, nil
}
