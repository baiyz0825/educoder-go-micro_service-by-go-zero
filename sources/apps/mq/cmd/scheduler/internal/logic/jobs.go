package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/scheduler/internal/svc"
)

type MqScheduler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewCronScheduler
//
//	@Description: 创建定时任务定时器
//	@param ctx
//	@param svcCtx
//	@return *MqScheduler
func NewCronScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *MqScheduler {
	return &MqScheduler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register
//
//	@Description: 注册任务
//	@receiver m
func (m *MqScheduler) Register() {
	m.updateUserEarnLogic()
}
