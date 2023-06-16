package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnlineTextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnlineTextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnlineTextLogic {
	return &UpdateOnlineTextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnlineTextLogic) UpdateOnlineText(in *pb.UpdateOnlineTextReq) (*pb.UpdateOnlineTextResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check param
	if in.GetID() == 0 || in.GetClassID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// dirty work check
	// create db
	content := []byte(in.Content)
	text := &model.OnlineText{
		Content:    &content,
		ClassID:    in.GetClassID(),
		Permission: &in.Permission,
		TextPoster: &in.TextPoster,
		TextName:   &in.TextName,
	}
	onlineText := l.svcCtx.Query.OnlineText
	_, err := onlineText.WithContext(l.ctx).Where(onlineText.ID.Eq(in.GetID())).Updates(text)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_UPDATE_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_UPDATE_ERR)
	}

	return &pb.UpdateOnlineTextResp{}, nil
}
