// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFile = "file"

// File mapped from table <file>
type File struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 文件自增id
	UUID          int64          `gorm:"column:uuid;not null" json:"uuid"`                                       // 文件uuid唯一标识
	Name          string         `gorm:"column:name;not null" json:"name"`                                       // 文件名称
	ObfuscateName string         `gorm:"column:obfuscate_name;not null" json:"obfuscate_name"`                   // 文件混淆名称
	Size          int64          `gorm:"column:size;not null" json:"size"`                                       // 文件占用空间大小（kb）
	Owner         int64          `gorm:"column:owner;not null" json:"owner"`                                     // 对应用户id
	Status        *int64         `gorm:"column:status;not null;default:1" json:"status"`                         // 0:已删除（云端） 1:（本地存储状态） 2:（云端存储状态，末态） 3:(用户隐藏状态）
	Type          int64          `gorm:"column:type;not null" json:"type"`                                       // 文件所属类型 文本0、文件1、视频2、图片3
	Class         int64          `gorm:"column:class;not null" json:"class"`                                     // 文件所属分类
	Suffix        string         `gorm:"column:suffix;not null" json:"suffix"`                                   // 文件后缀信息
	DownloadAllow *int64         `gorm:"column:download_allow;not null;default:1" json:"download_allow"`         // 是否允许查看 0 no 1 yes
	Link          *string        `gorm:"column:link" json:"link"`                                                // 文件云端存储目录
	CreateTime    *time.Time     `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime    *time.Time     `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"` // 更新时间
	DeleteTime    gorm.DeletedAt `gorm:"column:delete_time" json:"delete_time"`                                  // 删除时间
	FilePoster    *string        `gorm:"column:file_poster" json:"file_poster"`                                  // 文件头图
}

// TableName File's table name
func (*File) TableName() string {
	return TableNameFile
}