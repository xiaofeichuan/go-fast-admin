package model

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common"
)

type SysUser struct {
	Id       int64             `json:"id" gorm:"primaryKey;comment:Id"`
	UserName string            `json:"userName" gorm:"comment:用户账号"`
	NickName string            `json:"nickName" gorm:"comment:用户昵称"`
	UserType consts.UserType   `json:"userType" gorm:"comment:用户类型（0普通账号，1超级管理员）"`
	Email    string            `json:"email" gorm:"default:null;comment:用户邮箱"`
	Phone    *string           `json:"phone" gorm:"comment:手机号码"`
	Password string            `json:"password" gorm:"comment:密码"`
	Salt     string            `json:"salt" gorm:"comment:密码盐"`
	Gender   consts.UserGender `json:"gender" gorm:"comment:用户性别（0未知，1男，2女）"`
	Avatar   *string           `json:"avatar" gorm:"comment:头像地址"`
	Status   consts.UserStatus `json:"status" gorm:"comment:帐号状态（0正常 1停用）"`
	Remark   *string           `json:"remark" gorm:"comment:备注"`
	common.BaseModel
}

func (SysUser) TableName() string {
	return "sys_user"
}
