package model

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common"
)

type SysGenTableColumn struct {
	Id                int64                `json:"id" gorm:"primaryKey;comment:Id"`
	TableId           int64                `json:"tableId" gorm:"comment:表编号"`
	ColumnName        string               `json:"columnName" gorm:"comment:字段名"`
	ColumnDescription string               `json:"columnDescription" gorm:"comment:字段描述"`
	ColumnType        string               `json:"columnType" gorm:"comment:字段类型"`
	ParamName         string               `json:"jsonField" gorm:"comment:参数名称"`
	CodeField         string               `json:"codeField" gorm:"comment:代码字段"`
	CodeType          string               `json:"codeType" gorm:"comment:代码类型"`
	IsPk              bool                 `json:"isPk" gorm:"comment:是否主键（0否 1是）"`
	IsEdit            bool                 `json:"isEdit" gorm:"comment:是否编辑字段（0否 1是）"`
	IsList            bool                 `json:"isList" gorm:"comment:是否列表字段（0否 1是）"`
	IsQuery           bool                 `json:"isQuery" gorm:"comment:是否查询字段（0否 1是）"`
	QueryMethod       string               `json:"queryMethod" gorm:"comment:查询方式"`
	ComponentType     consts.ComponentType `json:"componentType" gorm:"comment:组件类型"`
	DictCode          string               `json:"dictCode" gorm:"comment:字典代码"`
	common.BaseModel
}

func (SysGenTableColumn) TableName() string {
	return "sys_gen_table_column"
}
