package model

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common"
)

type SysRole struct {
	Id       int64             `json:"id" gorm:"primaryKey;comment:角色ID"`
	RoleName string            `json:"roleName" gorm:"comment:角色名称"`
	RoleCode string            `json:"roleCode" gorm:"comment:角色代码"`
	Status   consts.RoleStatus `json:"status" gorm:"comment:菜单状态（0启用 1停用）"`
	Sort     int               `json:"sort" gorm:"comment:排序"`
	Remark   string            `json:"remark" gorm:"comment:备注"`
	common.BaseModel
}

func (SysRole) TableName() string {
	return "sys_role"
}
