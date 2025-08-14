package router

import (
	"github.com/gin-gonic/gin"
	adminRouter "go-fast-admin/server/app/admin/router"
	exampleRouter "go-fast-admin/server/app/example/router"
	"go-fast-admin/server/middleware"
)

// InitRouters 初始化路由
func InitRouters() *gin.Engine {
	routers := gin.Default()
	//允许修复当前请求路径，如/FOO和/..//Foo会被修复为/foo，并进行重定向，默认为 false。
	routers.RedirectFixedPath = true

	//异常处理
	routers.Use(middleware.Recover)

	//跨域
	routers.Use(middleware.Cors())

	//实例化路由
	systemRouter := adminRouter.SystemRouter{}
	exampleRouter := exampleRouter.TestRouter{}

	//授权路由
	privateGroup := routers.Group("/")
	privateGroup.Use(middleware.JwtAuth()).Use(middleware.Casbin())
	{
		systemRouter.InitPrivateRouter(privateGroup) //系统路由
	}

	//公开路由
	publicGroup := routers.Group("")
	{
		systemRouter.InitPublicRouter(publicGroup) //系统路由
		exampleRouter.InitPublicRouter(publicGroup)
	}

	return routers
}
