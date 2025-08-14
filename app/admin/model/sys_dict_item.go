package model

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common"
)

type SysDictItem struct {
	Id            int64                 `gorm:"primaryKey;comment:选项id"`
	DictId        int64                 `gorm:"comment:字典id"`
	DictItemName  string                `gorm:"comment:选项名称"`
	DictItemValue string                `gorm:"comment:选项值"`
	Status        consts.DictItemStatus `gorm:"comment:选项状态（0正常 1停用）"`
	Remark        string                `json:"remark" gorm:"comment:备注"`
	common.BaseModel
}
