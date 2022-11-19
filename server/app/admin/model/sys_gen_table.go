package model

import "go-fast-admin/server/common"

type SysGenTable struct {
	Id           uint64 `json:"id" gorm:"primaryKey;comment:Id"`
	TableName    string `json:"tableName" gorm:"column:table_name;comment:表名称"`
	TableComment string `json:"tableComment" gorm:"comment:表注释"`
	ModelName    string `json:"modelName" gorm:"comment:实体名称"`
	BusinessName string `json:"businessName" gorm:"comment:业务名称"`
	FunctionName string `json:"functionName" gorm:"comment:功能名称"`
	ParamName    string `json:"paramName" gorm:"comment:参数名称"`
	common.BaseModel
}

//func (SysTable) TableName() string {
//	return "sys_table"
//}
