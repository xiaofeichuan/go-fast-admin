package initialize

import (
	adminRouter "gin-fast-admin/server/app/admin/router"
	exampleRouter "gin-fast-admin/server/app/example/router"
	"github.com/gin-gonic/gin"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	Router := gin.Default()

	adminRouter := adminRouter.AdminRouterGroup
	exampleRouter := exampleRouter.ExampleRouterGroup

	// 授权
	privateGroup := Router.Group("")
	//middleware.JwtAuth()
	privateGroup.Use()
	{
		adminRouter.InitSysUserPrivateRouter(privateGroup)     //用户
		adminRouter.InitSysGenTablePrivateRouter(privateGroup) //代码生成器
	}

	// 无授权
	publicGroup := Router.Group("")
	{
		adminRouter.InitSysUserPublicRouter(publicGroup) //用户
		exampleRouter.InitTestPublicRouter(publicGroup)
	}

	return Router
}
