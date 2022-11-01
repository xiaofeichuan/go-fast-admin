package router

type AdminRouter struct {
	SysUserRouter
	SysTableRouter
}

var AdminRouterGroup = new(AdminRouter)
