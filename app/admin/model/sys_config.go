package model

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common"
)

type SysConfig struct {
	Id          int64               `gorm:"primaryKey;comment:Id"`
	ConfigName  string              `gorm:"comment:配置名称"`
	ConfigKey   string              `gorm:"comment:配置键名"`
	ConfigValue string              `gorm:"comment:配置键值"`
	Status      consts.ConfigStatus `gorm:"comment:配置状态（0正常 1停用）"`
	Remark      string              `json:"remark" gorm:"comment:备注"`
	common.BaseModel
}
