package svc

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/dao"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/traderpc"
	"github.com/hibiken/asynq"
	gormlogrus "github.com/onrik/gorm-logrus"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Query        *dao.Query
	RedisClient  *redis.Redis
	AliPayClient *alipay.Client
	TradeRpc     traderpc.Traderpc
	AsynqClient  *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// init
	q := initDaoQuery(c)
	rd := initRedisClient(c)
	return &ServiceContext{
		Config:       c,
		Query:        q,
		RedisClient:  rd,
		AliPayClient: initAlipayClient(c),
		TradeRpc:     traderpc.NewTraderpc(zrpc.MustNewClient(c.TradeRpcConfig)),
		AsynqClient:  initAsyncClient(c),
	}
}

// 初始化数据库设置
func initDaoQuery(c config.Config) *dao.Query {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		// 跳过默认事务
		SkipDefaultTransaction: false,
		// 命名策略（表名、列名生成规则）
		NamingStrategy: nil,
		// 创建更新时候，是否更新关联数据
		FullSaveAssociations: false,
		// 自定义日志 "github.com/onrik/gorm-logrus" 提供logrus包装实现
		Logger: gormlogrus.New(),
		// 创建时间函数
		NowFunc: nil,
		// 生成SQL不执行
		DryRun: false,
		// 是否禁止创建prepare stm
		PrepareStmt: false,
		// 禁用数据库健康检查
		DisableAutomaticPing: true,
		// 是否禁止自动创建外间约束
		DisableForeignKeyConstraintWhenMigrating: false,
		// 是否自动禁止外键约束
		IgnoreRelationshipsWhenMigrating: false,
		// 是否禁止嵌套事务
		DisableNestedTransaction: false,
		// 是否允许全局更新
		AllowGlobalUpdate: false,
		// 查询是否带上全部字段
		QueryFields: false,
		// 默认批量插入大小
		CreateBatchSize: 400,
		ClauseBuilders:  nil,
		ConnPool:        nil,
		Dialector:       nil,
		Plugins:         nil,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * 20)
	// 注入dao
	q := dao.Use(db)
	return q
}

// 初始化redis
func initRedisClient(c config.Config) *redis.Redis {
	// 创建 Redis 配置
	r := redis.MustNewRedis(c.Redis.RedisConf)
	if r != nil {
		return r
	}
	panic(errors.New("初始化Redis失败"))
}

// initAlipayClient
//
//	@Description: 初始化支付宝客户端
//	@param c
//	@return *alipay.Client
func initAlipayClient(c config.Config) *alipay.Client {
	privateKey, err := os.ReadFile(c.AliPay.AppPrivateKeyPath)
	if err != nil {
		panic(fmt.Sprintf("初始化支付渠道: 支付宝应用私钥失败:%v", err))
	}
	payClient, err := alipay.New(c.AliPay.AppId, string(privateKey), c.AliPay.IsProduction)
	// 加载应用公钥证书
	err = payClient.LoadAppPublicCertFromFile(c.AliPay.AppPublicCertPath)
	if err != nil {
		panic(fmt.Sprintf("初始化支付渠道: 加载应用公钥证书失败:%v", err))
	}
	// 加载支付宝根证书
	err = payClient.LoadAliPayRootCertFromFile(c.AliPay.ALiPayRootCertPath)
	if err != nil {
		panic(fmt.Sprintf("初始化支付渠道: 加载支付宝根证书失败:%v", err))
	}
	// 加载支付宝证书
	err = payClient.LoadAliPayPublicCertFromFile(c.AliPay.ALiPayPublicCertPath)
	if err != nil {
		panic(fmt.Sprintf("初始化支付渠道: 加载支付宝证书失败:%v", err))
	}
	return payClient
}

// initAsyncClient
//
//	@Description: 初始化 https://github.com/hibiken/asynq 异步任务队列客户端
//	@param c
//	@return *asynq.Client
func initAsyncClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
	})
}
