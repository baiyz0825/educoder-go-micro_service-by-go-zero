package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/internal/logic"
	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/mq.yaml", "指定配置文件")

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	// 创建任务统一配置
	jobs := logic.NewJobs(ctx, svcContext)
	// 注册任务
	mux := jobs.RegisterJobs()

	// 处理任务
	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("配置初始化定时任务失败！:%+v", err)
		panic(fmt.Sprintf("配置初始化定时任务失败！"))
	}
}
