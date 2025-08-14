package tools

import (
	"go-fast-admin/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type GormMySQLTool struct{}

var GormMySQL = new(GormMySQLTool)

// InitDB 初始化数据库
func (g *GormMySQLTool) InitDB() *gorm.DB {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               global.CONFIG.Database.DSN, // DSN data source name
		DefaultStringSize: 256,                        // string 类型字段的默认长度
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		//MySQL connection error
		panic(err)
	}
	return db
}

// GetTableInfoList 获取表信息
func (g *GormMySQLTool) GetTableInfoList() (list []DBTableInfo, err error) {
	tableSql := `
		SELECT 
			table_name AS TableName,
			table_comment AS TableDescription
		FROM 
			information_schema.TABLES 
		WHERE 
			table_schema = (SELECT DATABASE())`
	err = global.DB.Raw(tableSql).Find(&list).Error
	return list, err
}

// GetColumnInfoList 获取表信息
func (g *GormMySQLTool) GetColumnInfoList(tableName string) (list []DBColumnInfo, err error) {
	columnSql := `		
	SELECT
		table_name AS TableName,
		column_name AS ColumnName,
		column_comment AS ColumnDescription,
		data_type AS ColumnType,
		( CASE WHEN column_key = 'PRI' THEN true ELSE false END ) AS IsPk 
	FROM
		information_schema.COLUMNS 
	WHERE
		table_schema = (SELECT DATABASE()) AND table_name = ? ORDER BY ORDINAL_POSITION`
	err = global.DB.Raw(columnSql, tableName).Find(&list).Error
	return list, err
}
