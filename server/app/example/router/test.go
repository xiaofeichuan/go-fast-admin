package router

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/example/api"
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
