package dto

import (
	"time"
)

// SysGenTableColumnQuery 查询
type SysGenTableColumnQuery struct {
	TableId int64 `json:"tableId" form:"tableId"` //表id
}

type SysGenTableColumnVo struct {
	Id                int64     `json:"id"`                //Id
	TableId           int64     `json:"tableId"`           //表编号
	ColumnName        string    `json:"columnName"`        //字段名
	ColumnDescription string    `json:"columnDescription"` //字段描述
	ColumnType        string    `json:"columnType"`        //字段类型
	ParamName         string    `json:"paramName"`         //参数名称
	CodeField         string    `json:"codeField"`         //代码字段
	CodeType          string    `json:"codeType"`          //代码类型
	IsPk              bool      `json:"isPk"`              //是否主键（0否 1是）
	IsEdit            bool      `json:"isEdit"`            //是否编辑字段（0否 1是）
	IsList            bool      `json:"isList"`            //是否列表字段（0否 1是）
	IsQuery           bool      `json:"isQuery"`           //是否查询字段（0否 1是）
	QueryMethod       string    `json:"queryMethod"`       //查询方式
	ComponentType     string    `json:"componentType"`     //组件类型
	DictCode          string    `json:"dictCode"`          //字典代码
	CreatedAt         time.Time `json:"createdAt"`         //创建时间
	UpdatedAt         time.Time `json:"updatedAt"`         //更新时间
}

type SysGenTableColumnUpdateDto struct {
	Id                int64  `json:"id"`                //Id
	ColumnName        string `json:"columnName"`        //字段名
	ColumnDescription string `json:"columnDescription"` //字段描述
	ColumnType        string `json:"columnType"`        //字段类型
	ParamName         string `json:"paramName"`         //参数名称
	CodeField         string `json:"codeField"`         //代码字段
	CodeType          string `json:"codeType"`          //代码类型
	IsPk              bool   `json:"isPk"`              //是否主键（0否 1是）
	IsEdit            bool   `json:"isEdit"`            //是否编辑字段（0否 1是）
	IsList            bool   `json:"isList"`            //是否列表字段（0否 1是）
	IsQuery           bool   `json:"isQuery"`           //是否查询字段（0否 1是）
	QueryMethod       string `json:"queryMethod"`       //查询方式
	ComponentType     string `json:"componentType"`     //组件类型
	DictCode          string `json:"dictCode"`          //字典代码
}
