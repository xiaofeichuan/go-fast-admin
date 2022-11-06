package initialize

import (
	"fmt"
	"gin-fast-admin/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {

	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	//	logger.Config{
	//		SlowThreshold:             time.Second,   // 慢 SQL 阈值
	//		LogLevel:                  logger.Silent, // 日志级别
	//		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  false,         // 禁用彩色打印
	//	},
	//)

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
