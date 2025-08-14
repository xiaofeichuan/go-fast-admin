package model

type SysUserRole struct {
	Id     int64 `json:"id" gorm:"primaryKey;comment:用户和角色关联id"`
	UserId int64 `json:"userId" gorm:"comment:用户id"`
	RoleId int64 `json:"menuId" gorm:"comment:角色id"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
