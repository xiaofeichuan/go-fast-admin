package api

import (
	adminService "go-fast-admin/server/app/admin/service"
)

var (
	authService           = &adminService.SysAuthService{}
	userService           = &adminService.SysUserService{}
	genTableService       = &adminService.SysGenTableService{}
	genTableColumnService = &adminService.SysGenTableColumnService{}
)
