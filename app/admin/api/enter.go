package api

import adminService "go-fast-admin/server/app/admin/service"

var (
	authService           = &adminService.AuthService{}
	monitorService        = &adminService.MonitorService{}
	configService         = &adminService.SysConfigService{}
	dictService           = &adminService.SysDictService{}
	dictItemService       = &adminService.SysDictItemService{}
	userService           = &adminService.SysUserService{}
	roleService           = &adminService.SysRoleService{}
	menuService           = &adminService.SysMenuService{}
	genTableService       = &adminService.SysGenTableService{}
	genTableColumnService = &adminService.SysGenTableColumnService{}
	loginLogService       = &adminService.SysLoginLogService{}
	fileService           = &adminService.SysFileService{}
)
