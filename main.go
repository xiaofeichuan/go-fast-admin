package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/common/tools"
	"go-fast-admin/server/config"
	docs "go-fast-admin/server/docs" //注意：需要引用docs的路径，否则打开swagger时会出现错误
	"go-fast-admin/server/global"
	"go-fast-admin/server/router"
	"net/http"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 系统模块
// @version 1.0.0
// @description go-fast-admin 接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	//初始化配置
	global.CONFIG = config.InitConfig()

	//初始化数据库连接
	global.DB = tools.Gorm.Database().InitDB()

	//初始redis连接
	tools.InitRedis()

	//初始化路由
	var routers = router.InitRouters()

	// 初始化swagger文档
	InitSwagger(routers)

	//自定义 HTTP 服务器
	server := &http.Server{
		Addr:           ":9999",
		Handler:        routers,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Swagger文档地址:http://127.0.0.1%s/swagger/index.html", server.Addr)
	server.ListenAndServe()

}

// InitSwagger 初始化swagger文档
func InitSwagger(routers *gin.Engine) {
	//swagger文档
	docs.SwaggerInfo.Title = "go-fast-admin"
	docs.SwaggerInfo.Description = "Swagger Admin API"
	docs.SwaggerInfo.Version = "1.0"
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
