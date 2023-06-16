package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilResourcesUSerIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFilResourcesUSerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilResourcesUSerIdLogic {
	return &GetFilResourcesUSerIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFilResourcesUSerId
//
//	@Description: 通过资源id查询资源owner
//	@receiver l
//	@param in
//	@return *pb.GetFilResourcesUSerIdResp
//	@return error
func (l *GetFilResourcesUSerIdLogic) GetFilResourcesUSerId(in *pb.GetFilResourcesUSerIdReq) (*pb.GetFilResourcesUSerIdResp, error) {
	file := l.svcCtx.Query.File
	var uid int64
	err := file.WithContext(context.Background()).Select(file.Owner).Where(file.ID.Eq(in.GetResourcesId())).Scan(&uid)
	if err != nil {
		return nil, err
	}
	return &pb.GetFilResourcesUSerIdResp{
		UserId: uid,
	}, nil
}
