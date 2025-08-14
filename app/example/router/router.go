package router

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/example/api"
)

type TestRouter struct{}

var (
	testApi = &api.TestApi{}
)

// InitPublicRouter 初始化公开路由
func (t *TestRouter) InitPublicRouter(routerGroup *gin.RouterGroup) {

	//userRouter := routerGroup.Group("test")
	//{
	//	userRouter.POST("demo", testApi.Demo) // 用户登录
	//}
}

// InitPrivateRouter 初始化私有路由
func (t *TestRouter) InitPrivateRouter(routerGroup *gin.RouterGroup) {

}
