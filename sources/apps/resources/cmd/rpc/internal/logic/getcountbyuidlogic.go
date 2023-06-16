package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCountByUIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCountByUIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountByUIdLogic {
	return &GetCountByUIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCountByUId
// @Description: 使用uid 查询用户数据统计
// @receiver l
// @param in
// @return *pb.GetCountByUIdResp
// @return error
func (l *GetCountByUIdLogic) GetCountByUId(in *pb.GetCountByUIdReq) (*pb.GetCountByUIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetUid() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// search
	count := l.svcCtx.Query.Count
	data, err := count.WithContext(l.ctx).Where(count.UID.Eq(in.GetUid())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetCountByUIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return &pb.GetCountByUIdResp{}, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 复制数据
	countData := &pb.Count{}
	err = copier.Copy(countData, data)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return &pb.GetCountByUIdResp{}, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	countData.UpdateTime = data.UpdateTime.UnixMilli()
	countData.CreateTime = data.CreateTime.UnixMilli()

	return &pb.GetCountByUIdResp{
		Count: countData,
	}, nil
}
