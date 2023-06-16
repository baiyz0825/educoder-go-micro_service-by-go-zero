package logic

import (
	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/bo"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func (m *MqScheduler) updateUserEarnLogic() {
	task := asynq.NewTask(bo.UPDATE_USER_EARN_JOBS, nil)
	// every 20min minute exec
	entryID, err := m.svcCtx.Scheduler.Register("*/20 * * * *", task)
	if err != nil {
		logx.WithContext(m.ctx).Errorf("定时任务:更新用户所得任务，创建失败：err:%+v , task:%+v", err, task)
	}
	logx.WithContext(m.ctx).Infof("定时任务:更新用户所得任务，创建成功，任务id:%v", entryID)
}
