package initialize

import (
	"fmt"
	"go-fast-admin/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitDB 初始化数据库
func InitDB() *gorm.DB {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               global.CONFIG.Database.Default, // DSN data source name
		DefaultStringSize: 256,                            // string 类型字段的默认长度
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("连接数据库出现错误：" + err.Error())
	}
	return db
}
