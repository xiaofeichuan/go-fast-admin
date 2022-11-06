package dto

import "gin-fast-admin/server/common/dto/request"

type LoginDto struct {
	UserName  string `json:"userName"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// SysUserCreateDto 创建
type SysUserCreateDto struct {
	Account  string  `json:"account" ` //账号
	NickName string  `json:"nickName"` //用户昵称
	UserName string  `json:"userName"` // 用户名
	UserType int     `json:"userType"` //用户类型（0普通账号，1超级管理员）
	Email    string  `json:"email"`    //用户邮箱
	Phone    *string `json:"phone"`    //手机号码
	Password string  `json:"password"` //密码
	Sex      int     `json:"sex"`      //用户性别（0未知，1男，2女）
	Avatar   *string `json:"avatar"`   //头像地址
	Status   int     `json:"status"`   //帐号状态（0正常 1停用）
	Remark   *string `json:"remark"`   //备注
}

// SysUserUpdateDto 更新
type SysUserUpdateDto struct {
	Id uint64 `json:"id"` //编号
	SysUserCreateDto
}

// SysUserQuery 查询
type SysUserQuery struct {
	request.PageQuery
	NickName string `json:"nickName"form:"nickName"` //用户昵称
}

// SysUserVo 输出对象
type SysUserVo struct {
	Id       uint64  `json:"id"`        //编号
	Account  string  `json:"account"`   //用户账号
	NickName string  `json:"nickName"`  //用户昵称
	UserType int     `json:"userType"`  //用户类型（0普通账号，1超级管理员）
	Email    *string `json:"email" `    //用户邮箱
	Phone    *string `json:"phone"`     //手机号码
	Password string  `json:"password" ` //密码
	Sex      int     `json:"sex"`       //用户性别（0未知，1男，2女）
	Avatar   string  `json:"avatar"`    //头像地址
	Status   int     `json:"status"`    //帐号状态（0正常 1停用）
	Remark   *string `json:"remark"`    //备注
}
