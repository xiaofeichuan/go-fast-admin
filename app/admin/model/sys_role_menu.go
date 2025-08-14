package model

type SysRoleMenu struct {
	Id     int64 `json:"id" gorm:"primaryKey;comment:角色和菜单关联id"`
	RoleId int64 `json:"roleId" gorm:"comment:角色id"`
	MenuId int64 `json:"menuId" gorm:"comment:菜单id"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
