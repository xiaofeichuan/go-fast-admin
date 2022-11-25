package initialize

import (
	"github.com/gin-gonic/gin"
	adminRouter "go-fast-admin/server/app/admin/router"
	exampleRouter "go-fast-admin/server/app/example/router"
	"go-fast-admin/server/middleware"
)

// InitRouters 初始化路由
func InitRouters() *gin.Engine {
	Router := gin.Default()
	Router.RedirectFixedPath = true

	//跨域
	Router.Use(middleware.Cors())

	adminRouter := adminRouter.AdminRouterGroup
	exampleRouter := exampleRouter.ExampleRouterGroup

	//授权
	privateGroup := Router.Group("/")
	privateGroup.Use(middleware.JwtAuth())
	{
		adminRouter.InitSysAuthPrivateRouter(privateGroup)     //授权
		adminRouter.InitSysUserPrivateRouter(privateGroup)     //用户
		adminRouter.InitSysGenTablePrivateRouter(privateGroup) //代码生成器
	}

	//无授权
	publicGroup := Router.Group("")
	{
		adminRouter.InitSysAuthPublicRouter(publicGroup) //用户

		exampleRouter.InitTestPublicRouter(publicGroup)
	}

	return Router
}
