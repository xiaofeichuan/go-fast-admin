package api

import (
	adminService "gin-fast-admin/server/app/admin/service"
)

var (
	userService           = &adminService.SysUserService{}
	genTableService       = &adminService.SysGenTableService{}
	genTableColumnService = &adminService.SysGenTableColumnService{}
)
