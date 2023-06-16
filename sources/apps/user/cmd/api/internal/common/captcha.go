package common

import (
	"log"
	"strings"

	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var store *rdsStore

// InitCaptcha 初始化验证吗生成
func InitCaptcha(rd *redis.Redis) *base64Captcha.Captcha {
	// 默认生成器
	driver := base64Captcha.DefaultDriverDigit
	// 初始化redis store
	store = &rdsStore{redisClient: rd}
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha
}

// rdsStore An object implementing Store interface
type rdsStore struct {
	redisClient *redis.Redis
}

func (s *rdsStore) Set(id string, value string) error {
	// 5min有效
	ok, err := s.redisClient.SetnxEx(id, value, xconst.REDIS_CAPTCHA_EXPIRE_TIME)
	if !ok || err != nil {
		return err
	}
	return nil
}

func (s *rdsStore) Get(id string, clear bool) (value string) {
	// 拿
	val, err := s.redisClient.Get(id)
	// 出现错误
	if err != nil {
		log.Println(err)
		return ""
	}
	// 删除
	if clear {
		_, err := s.redisClient.Del(id)
		// 删除出错
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	// 没问题返回val
	return val
}

func (s *rdsStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	v = strings.TrimSpace(v)
	return v == strings.TrimSpace(answer)
}
