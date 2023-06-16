package logic

import (
	"context"

	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCountByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCountByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountByIdLogic {
	return &GetCountByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCountByIdLogic) GetCountById(in *pb.GetCountByIdReq) (*pb.GetCountByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetId() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// search
	count := l.svcCtx.Query.Count
	data, err := count.WithContext(l.ctx).Where(count.ID.Eq(in.GetId())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetCountByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return &pb.GetCountByIdResp{}, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 复制数据
	countData := &pb.Count{}
	err = copier.Copy(countData, data)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return &pb.GetCountByIdResp{}, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	countData.UpdateTime = data.UpdateTime.UnixMilli()
	countData.CreateTime = data.CreateTime.UnixMilli()

	return &pb.GetCountByIdResp{
		Count: countData,
	}, nil
}
