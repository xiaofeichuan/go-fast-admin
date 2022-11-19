package dto

import (
	"go-fast-admin/server/app/admin/model"
	"time"
)

type TableInfoVo struct {
	TableName    string                     `json:"tableName"`           //表名称
	TableComment string                     `json:"tableComment"`        //表注释
	ParamName    string                     `json:"paramName"`           //参数名称
	ModelName    string                     `json:"modelName"`           //实体名称
	BusinessName string                     `json:"businessName"`        //业务名称
	FunctionName string                     `json:"functionName"`        //功能名称
	CreateTime   time.Time                  `json:"createTime"`          //创建时间
	UpdateTime   *time.Time                 `json:"updateTime"`          //更新时间
	ColumnList   *[]model.SysGenTableColumn `json:"columnList" gorm:"-"` //字段列表
}

type TableColumnInfoVo struct {
	TableName     string `json:"tableName"`     //表名称
	ColumnName    string `json:"columnName"`    //字段名称
	ColumnComment string `json:"columnComment"` //字段注释
	ColumnType    string `json:"columnType" `   //字段类型
	IsPk          bool   `json:"isPk"`          //是否主键
}

type PreviewCodeVo struct {
	FileName    string `json:"fileName"`    //文件名称
	FileContent string `json:"fileContent"` //文件内容
}
