package global

import (
	"gin-fast-admin/server/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG *config.Config
)
