package xconst

var thirdMapping map[int64]string

func init() {
	// 微信0、QQ1、Github2、Gitee3
	thirdMapping = map[int64]string{
		0: "微信",
		1: "QQ",
		2: "Github",
		3: "Gitee",
	}
}

// GetThirdMapping
//
//	@Description: 获取账户绑定分类
//	@param typeCode
//	@return string
func GetThirdMapping(typeCode int64) string {
	if s, ok := thirdMapping[typeCode]; ok {
		return s
	}
	return ""
}

// 业务常量
const (
	// REDIS_CAPTCHA_EXPIRE_TIME redis验证码过期时间
	REDIS_CAPTCHA_EXPIRE_TIME = 5 * 60

	// JWT 生成KV对
	JWT_USER_USERUNIQUEID = "userUniqueId"
	JWT_USER_ID           = "userId"
)

// 文件资源类型
const (
	TEXT    int64 = 0
	FILE    int64 = 1
	VIDEO   int64 = 2
	PICTURE int64 = 3
	UNKONWN int64 = 9999
)

// order
const (
	ORDER_STATUS_CREATE = 0
	ORDER_STATUS_PAYING = 1
	ORDER_STATUS_PAYED  = 2
	// ORDER_SYSTEM_MODE 商品模块
	ORDER_SYSTEM_MODE_TRADE = 0
	PAY_PATH_ALIPAY         = 1
	PAY_PATH_WECHAT         = 0
)

const (
	// TRADE_SUCCESS 阿里支付成功
	TRADE_SUCCESS = "TRADE_SUCCESS"
	// TRADE_FINISHED 阿里支付结束
	TRADE_FINISHED = "TRADE_FINISHED"
)

// 支付宝侧错误码
const (
	// ALIPAY_TRADE_NOT_EXIST 交易不存在（订单扫码之后才会创建）
	ALIPAY_TRADE_NOT_EXIST = "ACQ.TRADE_NOT_EXIST"
	// ALIPAY_WAIT_BUYER_PAY 等待付款
	ALIPAY_WAIT_BUYER_PAY = "WAIT_BUYER_PAY"
	// AlIPAY_TRADE_CLOSED 交易关闭
	AlIPAY_TRADE_CLOSED = "TRADE_CLOSED"
	// ALIPAY_TRADE_SUCCESS 交易成功
	ALIPAY_TRADE_SUCCESS = "TRADE_SUCCESS"
	// ALIPAY_TRADE_FINISHED 结束不可退款
	ALIPAY_TRADE_FINISHED = "TRADE_FINISHED"
)

const (
	NOT_SET_TYPE = 0
	FILE_TYEP    = 1
	TEXT_TYPE    = 2
)

const (
	PERMISSION_TRUE  = 1
	PERMISSION_FALSE = 0
)
