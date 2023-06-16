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

type GetOnlineTextByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnlineTextByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnlineTextByIdLogic {
	return &GetOnlineTextByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnlineTextByIdLogic) GetOnlineTextById(in *pb.GetOnlineTextByIdReq) (*pb.GetOnlineTextByIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// delete
	onlineText := l.svcCtx.Query.OnlineText
	data, err := onlineText.WithContext(l.ctx).Where(onlineText.ID.Eq(in.GetID())).First()
	if err == gorm.ErrRecordNotFound {
		return &pb.GetOnlineTextByIdResp{}, nil
	}
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	// 复制数据
	onlineTextData := &pb.OnlineText{}
	err = copier.Copy(onlineTextData, data)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
	}
	onlineTextData.UpdateTime = data.UpdateTime.UnixMilli()
	onlineTextData.CreateTime = data.CreateTime.UnixMilli()
	return &pb.GetOnlineTextByIdResp{
		OnlineText: onlineTextData,
	}, nil
}
