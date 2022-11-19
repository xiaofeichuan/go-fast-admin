package model

import "go-fast-admin/server/common"

type SysGenTableColumn struct {
	Id            uint64 `json:"id" gorm:"primaryKey;comment:Id"`
	TableId       uint64 `json:"tableId" gorm:"comment:表编号"`
	ColumnName    string `json:"columnName" gorm:"comment:字段名"`
	ColumnComment string `json:"columnComment" gorm:"comment:字段注释"`
	ColumnType    string `json:"columnType" gorm:"comment:字段类型"`
	ParamName     string `json:"jsonField" gorm:"comment:参数名称"`
	GoField       string `json:"goField" gorm:"comment:go字段名"`
	GoType        string `json:"goType" gorm:"comment:go类型"`
	IsPk          bool   `json:"isPk" gorm:"comment:是否主键（0否 1是）"`
	IsEdit        bool   `json:"isEdit" gorm:"comment:是否编辑字段（0否 1是）"`
	IsList        bool   `json:"isList" gorm:"comment:是否列表字段（0否 1是）"`
	IsQuery       bool   `json:"isQuery" gorm:"comment:是否查询字段（0否 1是）"`
	QueryMethod   string `json:"queryMethod" gorm:"comment:是否查询字段（0否 1是）"`
	common.BaseModel
}

func (SysGenTableColumn) TableName() string {
	return "sys_gen_table_column"
}
