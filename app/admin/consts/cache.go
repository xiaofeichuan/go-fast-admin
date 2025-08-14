package consts

import "time"

const (
	CacheTime                 = time.Minute * 30       //缓存时间（可应用大部分业务缓存）
	CacheKeySysUserMenu       = "sys_user_menu_"       //用户菜单
	CacheKeySysAllPermission  = "sys_all_permission"   //所有权限标识
	CacheKeySysUserPermission = "sys_user_permission_" //用户权限标识
)
