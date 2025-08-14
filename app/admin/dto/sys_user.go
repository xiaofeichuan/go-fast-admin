package dto

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/common/dto/request"
	"time"
)

// SysUserQuery 查询
type SysUserQuery struct {
	request.PageQuery
	UserName string `json:"nickName"form:"userName"` //用户昵称
}

// SysUserVo 输出对象
type SysUserVo struct {
	Id        int64             `json:"id"`              //编号
	UserName  string            `json:"userName"`        //用户账号
	NickName  string            `json:"nickName" `       //用户昵称
	UserType  int               `json:"userType"`        //用户类型（0普通账号，1超级管理员）
	Email     *string           `json:"email" `          //用户邮箱
	Phone     *string           `json:"phone"`           //手机号码
	Gender    consts.UserGender `json:"gender"`          //用户性别（0未知，1男，2女）
	Avatar    string            `json:"avatar"`          //头像地址
	Status    int               `json:"status"`          //帐号状态（0正常 1停用）
	Remark    *string           `json:"remark"`          //备注
	CreatedAt time.Time         `json:"createdAt"`       //创建时间
	UpdatedAt time.Time         `json:"updatedAt"`       //更新时间
	RoleIds   []int64           `json:"roleIds"gorm:"-"` //角色
}

// SysUserAddDto 创建
type SysUserAddDto struct {
	UserName string            `json:"userName"`  //用户账号
	NickName string            `json:"nickName" ` //用户昵称
	Email    string            `json:"email"`     //用户邮箱
	Phone    *string           `json:"phone"`     //手机号码
	Password string            `json:"password"`  //密码
	Gender   consts.UserGender `json:"gender"`    //用户性别（0未知，1男，2女）
	Status   consts.UserStatus `json:"status"`    //帐号状态（0正常 1停用）
	Remark   *string           `json:"remark"`    //备注
	RoleIds  []int64           `json:"roleIds"`   //角色
}

// SysUserUpdateDto 更新
type SysUserUpdateDto struct {
	Id       int64             `json:"id"`        //编号
	NickName string            `json:"nickName" ` //用户昵称
	Email    string            `json:"email"`     //用户邮箱
	Phone    *string           `json:"phone"`     //手机号码
	Gender   consts.UserGender `json:"gender"`    //用户性别（0未知，1男，2女）
	Status   consts.UserStatus `json:"status"`    //帐号状态（0正常 1停用）
	Remark   *string           `json:"remark"`    //备注
	RoleIds  []int64           `json:"roleIds"`   //角色
}

type ResetPwdDto struct {
	UserId   int64  `json:"userId"`   // 用户Id
	Password string `json:"password"` // 密码
}

type UpdatePwdDto struct {
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type SetUserStateDto struct {
	UserId int64             `json:"userId"` // 用户Id
	Status consts.UserStatus `json:"status"` //帐号状态（0正常 1停用）
}
