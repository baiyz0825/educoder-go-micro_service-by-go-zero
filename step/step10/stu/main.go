package main

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"stu/query"
)

func main() {
	// gen.GenModel()

	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open("root:123123@(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"))
	// 注入db到查询器中
	query.SetDefault(db)
	user, err := query.User.WithContext(context.Background()).Where(query.User.UserName.Eq("bigWhite")).First()
	if err != nil {
		fmt.Printf("error is :%v", err)
		return
	}
	fmt.Printf("user name:%v,passwd:%v", *user.UserName, *user.Passwd)
	return
}

// 输出样例:user name:bigwhite,passwd:abcd1234
