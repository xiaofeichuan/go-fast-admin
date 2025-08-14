package global

import (
	"go-fast-admin/server/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB          //数据库
	CONFIG *config.AppConfig //配置
)
