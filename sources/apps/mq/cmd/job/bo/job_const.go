package bo

const (
	// DELETE_EXPIRE_JOBS 删除过期订单任务
	DELETE_EXPIRE_JOBS = "jobs:order:deleteExpire"
	// CHECK_ORDER_STATUS_JOBS 检查订单支付状态
	CHECK_ORDER_STATUS_JOBS = "jobs:order:checkStatus"
	// UPDATE_USER_EARN
	UPDATE_USER_EARN_JOBS = "jobs:order:updateUSerEarn"
)
