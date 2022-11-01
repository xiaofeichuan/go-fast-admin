package router

import (
	"gin-fast-admin/server/app/example/api"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

// 初始化无权限路由
func (s *TestRouter) InitTestPublicRouter(routerGroup *gin.RouterGroup) {
	var testApi = api.TestApi{}
	userRouter := routerGroup.Group("test")
	{
		userRouter.POST("demo", testApi.Demo) // 用户登录
	}
}
