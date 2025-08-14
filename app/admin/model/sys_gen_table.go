package model

import (
	"go-fast-admin/server/common"
)

type SysGenTable struct {
	Id               int64  `json:"id" gorm:"primaryKey;comment:Id"`
	TableName        string `json:"tableName" gorm:"column:table_name;comment:表名称"`
	TableDescription string `json:"tableDescription" gorm:"comment:表描述"`
	ModelName        string `json:"modelName" gorm:"comment:实体名称"`
	BusinessName     string `json:"businessName" gorm:"comment:业务名称"`
	ModuleName       string `json:"moduleName" gorm:"comment:模块名称"`
	MenuParentId     int64  `json:"menuParentId" gorm:"comment:上级菜单"`
	common.BaseModel
}

//func (SysTable) TableName() string {
//	return "sys_table"
//}
