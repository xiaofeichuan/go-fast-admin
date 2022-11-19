package router

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/api"
)

type SysUserRouter struct{}

// InitSysUserPrivateRouter 初始化权限路由
func (s *SysUserRouter) InitSysUserPrivateRouter(routerGroup *gin.RouterGroup) {
	var sysUserApi = api.SysUserApi{}
	// 路由规范：业务/模板/操作，比如：taobao/user/add
	userRouter := routerGroup.Group("/system/user")
	{
		userRouter.GET("page", sysUserApi.Page)      // 用户分页查询
		userRouter.GET("detail", sysUserApi.Detail)  // 用户详情
		userRouter.GET("list", sysUserApi.List)      // 用户列表
		userRouter.POST("add", sysUserApi.Add)       // 添加用户
		userRouter.POST("update", sysUserApi.Update) // 更新用户
		userRouter.POST("delete", sysUserApi.Delete) // 删除用户
	}
}
