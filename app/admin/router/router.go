package router

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/api"
)

// SystemRouter 系统路由
// 建议：
// 当接口数量超出一定范围，可适当分类或将复杂模块独立文件整理
// 路由规范：业务/模块/操作，比如：mall/user/add
type SystemRouter struct{}

var (
	authApi           = &api.AuthApi{}
	monitorApi        = &api.SysMonitorApi{}
	configApi         = &api.SysConfigApi{}
	dictApi           = &api.SysDictApi{}
	dictItemApi       = &api.SysDictItemApi{}
	menuApi           = &api.SysMenuApi{}
	userApi           = &api.SysUserApi{}
	roleApi           = &api.SysRoleApi{}
	genTableApi       = &api.SysGenTableApi{}
	genTableColumnApi = &api.SysGenTableColumnApi{}
	loginLogApi       = &api.SysLoginLogApi{}
)

// InitPublicRouter 初始化公开路由
func (s *SystemRouter) InitPublicRouter(routerGroup *gin.RouterGroup) {

	//权限
	authRouter := routerGroup.Group("/system/auth")
	{
		authRouter.GET("captcha", authApi.GenerateCaptcha) // 生成验证码
		authRouter.POST("login", authApi.Login)            // 用户登录
	}
}

// InitPrivateRouter 初始化私有路由
func (s *SystemRouter) InitPrivateRouter(routerGroup *gin.RouterGroup) {

	//权限
	authRouter := routerGroup.Group("/system/auth")
	{
		authRouter.GET("userInfo", authApi.GetUserInfo) // 获取用户信息
		authRouter.GET("menu", authApi.GetAuthMenu)     // 获取用户菜单（树状）
		authRouter.POST("updatePwd", authApi.UpdatePwd) // 更新密码
	}

	//用户
	userRouter := routerGroup.Group("/system/user")
	{
		userRouter.GET("query", userApi.Query)              // 用户分页查询
		userRouter.POST("add", userApi.Add)                 // 添加用户
		userRouter.POST("update", userApi.Update)           // 更新用户
		userRouter.POST("delete", userApi.Delete)           // 删除用户
		userRouter.GET("detail", userApi.Detail)            // 用户详情
		userRouter.POST("resetPwd", userApi.ResetPwd)       // 重置密码
		userRouter.POST("setStatus", userApi.SetUserStatus) // 设置状态
	}

	//菜单
	menuRouter := routerGroup.Group("/system/menu")
	{
		menuRouter.GET("query", menuApi.Query)      // 菜单分页查询
		menuRouter.POST("add", menuApi.Add)         // 添加菜单
		menuRouter.POST("update", menuApi.Update)   // 更新菜单
		menuRouter.POST("delete", menuApi.Delete)   // 删除菜单
		menuRouter.GET("detail", menuApi.Detail)    // 菜单详情
		menuRouter.GET("tree", menuApi.GetMenuTree) // 菜单树状
	}

	//角色
	roleRouter := routerGroup.Group("/system/role")
	{
		roleRouter.GET("query", roleApi.Query)    // 角色查询
		roleRouter.POST("add", roleApi.Add)       // 添加角色
		roleRouter.POST("update", roleApi.Update) // 更新角色
		roleRouter.POST("delete", roleApi.Delete) // 删除角色
		roleRouter.GET("detail", roleApi.Detail)  // 角色详情
		roleRouter.GET("list", roleApi.List)      // 角色列表
	}

	//配置
	configRouter := routerGroup.Group("/system/config")
	{
		configRouter.GET("query", configApi.Query)    // 配置查询
		configRouter.POST("add", configApi.Add)       // 添加配置
		configRouter.POST("update", configApi.Update) // 更新配置
		configRouter.POST("delete", configApi.Delete) // 删除配置
		configRouter.GET("detail", configApi.Detail)  // 配置详情
	}

	//字典
	dictRouter := routerGroup.Group("/system/dict")
	{
		dictRouter.GET("query", dictApi.Query)    // 字典查询
		dictRouter.POST("add", dictApi.Add)       // 添加字典
		dictRouter.POST("update", dictApi.Update) // 更新字典
		dictRouter.POST("delete", dictApi.Delete) // 删除字典
		dictRouter.GET("detail", dictApi.Detail)  // 字典详情
		dictRouter.GET("list", dictApi.List)      // 字典列表
	}

	//字典选项
	dictItemRouter := routerGroup.Group("/system/dictItem")
	{
		dictItemRouter.GET("query", dictItemApi.Query)    // 字典选项查询
		dictItemRouter.POST("add", dictItemApi.Add)       // 添加字典选项
		dictItemRouter.POST("update", dictItemApi.Update) // 更新字典选项
		dictItemRouter.POST("delete", dictItemApi.Delete) // 删除字典选项
		dictItemRouter.GET("detail", dictItemApi.Detail)  // 字典选项详情
		dictItemRouter.GET("list", dictItemApi.List)      // 字典选项列表
	}

	//代码生成（表）
	genTableRouter := routerGroup.Group("/system/genTable")
	{
		genTableRouter.GET("query", genTableApi.Query)             // 代码生成查询
		genTableRouter.POST("add", genTableApi.Add)                // 添加代码生成
		genTableRouter.POST("update", genTableApi.Update)          // 更新代码生成
		genTableRouter.POST("delete", genTableApi.Delete)          // 删除代码生成
		genTableRouter.GET("detail", genTableApi.Detail)           // 代码生成详情
		genTableRouter.GET("tableList", genTableApi.GetTableList)  // 获取表列表
		genTableRouter.GET("previewCode", genTableApi.PreviewCode) // 导入表
	}

	//代码生成（表字段）
	genTableColumnRouter := routerGroup.Group("/system/genTableColumn")
	{
		genTableColumnRouter.POST("update", genTableColumnApi.Update) // 更新代码生成（表字段）
		genTableColumnRouter.GET("list", genTableColumnApi.List)      // 代码生成（表字段）列表
	}

	//系统监控
	monitorRouter := routerGroup.Group("/system/monitor")
	{
		monitorRouter.GET("cacheKeys", monitorApi.GetCacheKeys)   // 获取缓存key
		monitorRouter.GET("cacheValue", monitorApi.GetCacheValue) // 获取缓存value
		monitorRouter.POST("deleteCache", monitorApi.DeleteCache) // 删除缓存
		monitorRouter.GET("serverInfo", monitorApi.GetServerInfo) // 获取服务器信息
	}

	//登录日志
	loginLogRouter := routerGroup.Group("system/loginLog")
	{
		loginLogRouter.GET("query", loginLogApi.Query)   // 登录日志分页查询
		loginLogRouter.GET("detail", loginLogApi.Detail) // 登录日志详情
	}
}
