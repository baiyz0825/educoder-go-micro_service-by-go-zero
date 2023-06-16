// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameResComment = "res_comment"

// ResComment mapped from table <res_comment>
type ResComment struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 评论自增id
	Owner      int64          `gorm:"column:owner;not null" json:"owner"`                                     // 评论所属人信息
	ResourceID int64          `gorm:"column:resource_id;not null" json:"resource_id"`                         // 资源id
	Content    *string        `gorm:"column:content;not null;default:none set" json:"content"`                // 评论内容
	CreateTime *time.Time     `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime *time.Time     `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"` // 更新时间
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"delete_time"`                                  // 删除时间
}

// TableName ResComment's table name
func (*ResComment) TableName() string {
	return TableNameResComment
}
