package model

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common"
)

type SysDict struct {
	Id       int64             `gorm:"primaryKey;comment:Id"`
	DictName string            `gorm:"comment:字典名称"`
	DictCode string            `gorm:"comment:字典代码"`
	DictType consts.DictType   `gorm:"comment:字典类型（0int，1string）"`
	Status   consts.DictStatus `gorm:"comment:字典状态（0正常 1停用）"`
	Remark   string            `json:"remark" gorm:"comment:备注"`
	common.BaseModel
}
