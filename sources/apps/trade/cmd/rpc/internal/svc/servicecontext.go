package svc

import (
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/internal/dao"
	gormlogrus "github.com/onrik/gorm-logrus"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	Query       *dao.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	// init
	q := initDaoQuery(c)
	rd := initRedisClient(c)
	return &ServiceContext{
		Config:      c,
		Query:       q,
		RedisClient: rd,
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
