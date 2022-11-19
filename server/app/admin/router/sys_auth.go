package router

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/api"
)

// InitSysAuthPrivateRouter 初始化权限路由
func (s *SysUserRouter) InitSysAuthPrivateRouter(routerGroup *gin.RouterGroup) {
	var sysAuthApi = api.SysAuthApi{}
	authRouter := routerGroup.Group("/")
	{
		authRouter.GET("getUserInfo", sysAuthApi.GetUserInfo) // 获取用户信息
		//authRouter.GET("router", sysAuthApi.GetMenuRouter) // 获取菜单路由（树形）
	}
}

// InitSysAuthPublicRouter 初始化无权限路由
func (s *SysUserRouter) InitSysAuthPublicRouter(routerGroup *gin.RouterGroup) {
	var sysAuthApi = api.SysAuthApi{}
	authRouter := routerGroup.Group("/")
	{
		authRouter.GET("captcha", sysAuthApi.GenerateCaptcha) // 生成验证码
		authRouter.POST("login", sysAuthApi.Login)            // 用户登录
	}
}
