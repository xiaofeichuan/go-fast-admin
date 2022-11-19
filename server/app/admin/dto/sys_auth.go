package dto

type CaptchaVo struct {
	CaptchaId     string `json:"captchaId"`
	CaptchaBase64 string `json:"captchaBase64"`
}

type LoginDto struct {
	Account   string `json:"account" `  //账号
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type UserInfoVo struct {
	Id       uint64  `json:"id"`       //编号
	Account  string  `json:"account" ` //账号
	UserName string  `json:"userName"` //用户名称
	UserType int     `json:"userType"` //用户类型（0普通账号，1超级管理员）
	Email    string  `json:"email"`    //用户邮箱
	Phone    *string `json:"phone"`    //手机号码
	Sex      int     `json:"sex"`      //用户性别（0未知，1男，2女）
	Avatar   *string `json:"avatar"`   //头像地址
	Status   int     `json:"status"`   //帐号状态（0正常 1停用）
	Remark   *string `json:"remark"`   //备注
}
