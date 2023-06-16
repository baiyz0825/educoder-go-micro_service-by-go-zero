package gen

import (
	"gorm.io/gen"
	"model/model"
)

func GenCode() {

	// 声明一个生成器
	g := gen.NewGenerator(gen.Config{
		// 生成器输出目录
		OutPath: "./query",
		Mode:    gen.WithDefaultQuery,
	})

	// 生成Gorm自带的基本查询结构
	g.ApplyBasic(model.User{})

	// 提交需要生成的基本自定义接口
	g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	// 执行生成器
	g.Execute()
}
