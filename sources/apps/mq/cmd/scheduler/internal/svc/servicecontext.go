package svc

import (
	"time"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/scheduler/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config    config.Config
	Scheduler *asynq.Scheduler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Scheduler: initScheduler(c),
	}
}

// 初始化定时任务
func initScheduler(c config.Config) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}, &asynq.SchedulerOpts{
			Location: location,
			// 任务enqueue回调
			PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
				logx.Infof("开始处理任务：%v", info)
			},
		})
}
