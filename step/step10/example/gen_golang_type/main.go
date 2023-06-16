package main

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"model/query"
)

func main() {
	// 调用生成器
	// gen.GenCode()
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/user_test?charset=utf8mb4&parseTime=True&loc=Local"))
	// 注入db到查询器中
	query.SetDefault(db)

	user, err := query.User.WithContext(context.Background()).FindByUserName("bigWhite")
	if err != nil {
		fmt.Printf("error is :%v", err)
		return
	}
	fmt.Printf("user is :%v", user)
	return
}
