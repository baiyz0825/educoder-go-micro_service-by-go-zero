package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnlineTextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnlineTextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnlineTextLogic {
	return &AddOnlineTextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------在线文本资源信息-----------------------

func (l *AddOnlineTextLogic) AddOnlineText(in *pb.AddOnlineTextReq) (*pb.AddOnlineTextResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// check param
	if in.GetOwner() == 0 || in.GetClassID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// dirty work check
	// gen uuid && db data
	id, err := utils.GenSnowFlakeId()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("生成系统id错误")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// create db
	content := []byte(in.Content)
	text := &model.OnlineText{
		UUID:       id,
		TypeSuffix: in.GetTypeSuffix(),
		Owner:      in.GetOwner(),
		Content:    &content,
		ClassID:    in.GetClassID(),
		Permission: &in.Permission,
		TextPoster: &in.TextPoster,
		TextName:   &in.TextName,
	}
	err = l.svcCtx.Query.OnlineText.WithContext(l.ctx).Create(text)
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error(xerr.NewErrCode(xerr.RPC_INSERT_ERR))
		return nil, xerr.NewErrCode(xerr.RPC_INSERT_ERR)
	}
	return &pb.AddOnlineTextResp{}, nil
}
