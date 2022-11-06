package initialize

import (
	"fmt"
	"gin-fast-admin/server/config"
	"github.com/spf13/viper"
)

func InitConfig() *config.Config {
	//导入配置文件
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置出现错误：" + err.Error())
	}
	var _config *config.Config
	//将配置文件读到结构体中
	err = viper.Unmarshal(&_config)
	if err != nil {
		fmt.Println("配置文件转实体出现错误" + err.Error())
	}
	return _config
}
