package utils

import (
	"time"
)

// GetContextDefaultTime
//
//	@Description: 获取默认上下文时间
//	@return time.Time
func GetContextDefaultTime() time.Time {
	return time.Now().Add(time.Second * 40)
}

// GetContextDuration
//
//	@Description: 默认context时长
//	@return time.Duration
func GetContextDuration() time.Duration {
	return 40 * time.Second
}

// GetJwtIatTime
//
//	@Description: 获取jwt产生时间（当前时间）
//	@return time.Time
func GetJwtIatTime() int64 {
	return time.Now().Unix()
}

// GetJwtExpireDefaultTime
//
//	@Description: 获取默认jwt过期时间(12h)
//	@return time.Time
func GetJwtExpireDefaultTime() int64 {
	return 3600 * 24
}
