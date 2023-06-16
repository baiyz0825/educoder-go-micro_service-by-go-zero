// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserEarn = "user_earn"

// UserEarn mapped from table <user_earn>
type UserEarn struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 统计表id
	UserID     int64          `gorm:"column:user_id;not null" json:"user_id"`                                 // 用户id
	EarnNum    float64        `gorm:"column:earn_num;not null" json:"earn_num"`                               // 用户入账
	PayNum     float64        `gorm:"column:pay_num;not null" json:"pay_num"`                                 // 用户支出价格
	CreateTime *time.Time     `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime *time.Time     `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"` // 更新时间
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"delete_time"`                                  // 删除时间
}

// TableName UserEarn's table name
func (*UserEarn) TableName() string {
	return TableNameUserEarn
}