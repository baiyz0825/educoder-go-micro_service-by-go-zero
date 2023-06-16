package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // 传入Gorm相关属性
	ID         uint   `gorm:"primary_key"`
	UserName   string `gorm:"column:name"`
	Passwd     string `gorm:"column:password"`
}
