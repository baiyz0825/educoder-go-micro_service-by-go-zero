package main

import (
	"context"
	"flag"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/scheduler/internal/config"
	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/scheduler/internal/logic"
	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/scheduler/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/scheduler.yaml", "指定配置文件")

func main() {
	var c config.Config

	conf.MustLoad(*configFile, &c)
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	mqueueScheduler := logic.NewCronScheduler(ctx, svcContext)
	mqueueScheduler.Register()
	if err := svcContext.Scheduler.Run(); err != nil {
		logx.Errorf("定时任务添加失败:%+v", err)
		panic(err)
	}
}
