package logic

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/bo"
	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/internal/svc"
	"github.com/hibiken/asynq"
)

type Job struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewJobs
//
//	@Description: 初始化job上下文
//	@param ctx
//	@param serviceContext
//	@return *Job
func NewJobs(ctx context.Context, serviceContext *svc.ServiceContext) *Job {
	return &Job{
		ctx:    ctx,
		svcCtx: serviceContext,
	}
}

// RegisterJobs
//
//	@Description: 注册任务
//	@receiver j
//	@return *asynq.ServeMux
func (j *Job) RegisterJobs() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	// 注册订单超时删除任务
	mux.Handle(bo.DELETE_EXPIRE_JOBS, NewDoOrderLogic(j.svcCtx))
	// 注册订单支付状态检查函数
	mux.Handle(bo.CHECK_ORDER_STATUS_JOBS, NewCheckOrderStatusJobLogic(j.svcCtx))
	// 更新用户所得定时任务
	mux.Handle(bo.UPDATE_USER_EARN_JOBS, NewUpdateUserEarnJobLogic(j.svcCtx))
	return mux
}
