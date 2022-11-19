package global

import (
	"go-fast-admin/server/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG *config.Config
)
