package router

import (
	"gin-fast-admin/server/app/admin/api"
	"github.com/gin-gonic/gin"
)

type SysUserRouter struct{}

// InitSysUserPrivateRouter 初始化权限路由
func (s *SysUserRouter) InitSysUserPrivateRouter(routerGroup *gin.RouterGroup) {
	var sysUserApi = api.SysUserApi{}
	userRouter := routerGroup.Group("sysUser")
	{
		userRouter.GET("page", sysUserApi.Page)      // 用户分页查询
		userRouter.GET("detail", sysUserApi.Detail)  // 用户详情
		userRouter.GET("list", sysUserApi.List)      // 用户列表
		userRouter.POST("add", sysUserApi.Add)       // 添加用户
		userRouter.POST("update", sysUserApi.Update) // 更新用户
		userRouter.POST("delete", sysUserApi.Delete) // 删除用户
	}
}

// InitSysUserPublicRouter 初始化无权限路由
func (s *SysUserRouter) InitSysUserPublicRouter(routerGroup *gin.RouterGroup) {
	//var sysUserApi = new(api.SysUserApi)
	//userRouter := routerGroup.Group("sysUser")
	//{
	//	userRouter.GET("login", sysUserApi.UserLogin) // 用户登录
	//}
}
