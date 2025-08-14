package tools

import (
	"go-fast-admin/server/global"
	"gorm.io/gorm"
)

type GormTool struct{}

var Gorm = new(GormTool)

type Database interface {
	// InitDB 初始化数据库
	InitDB() *gorm.DB

	// GetTableInfoList 获取表信息
	GetTableInfoList() (list []DBTableInfo, err error)

	// GetColumnInfoList 获取表信息
	GetColumnInfoList(tableName string) (list []DBColumnInfo, err error)
}

func (gt *GormTool) Database() Database {
	switch global.CONFIG.Database.Type {
	case "mysql":
		return GormMySQL
	case "pgsql":
		return GormPgSQL
	default:
		return GormMySQL
	}
}

type DBTableInfo struct {
	TableName        string `json:"tableName"`        //表名称
	TableDescription string `json:"tableDescription"` //表描述
	SchemaName       string `json:"schemaName"`       //pgsql架构名称
}

type DBColumnInfo struct {
	TableName         string `json:"tableName"`         //表名称
	ColumnName        string `json:"columnName"`        //字段名称
	ColumnDescription string `json:"columnDescription"` //字段描述
	ColumnType        string `json:"columnType"`        //字段类型
	IsPk              bool   `json:"isPk"`              //是否主键
}
