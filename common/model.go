package common

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel
// GORM 会直接填充更新时间、更新时间、删除时间
// 文档地址：https://gorm.io/zh_CN/docs/models.html#创建-x2F-更新时间追踪（纳秒、毫秒、秒、Time）
type BaseModel struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
