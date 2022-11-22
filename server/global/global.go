package global

import (
	"github.com/go-redis/redis/v8"
	"go-fast-admin/server/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB       //数据库
	CONFIG *config.Config //配置
	REDIS  *redis.Client  //redis
)
