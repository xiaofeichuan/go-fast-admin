package dto

import (
	"go-fast-admin/server/app/admin/consts"
)

type LoginDto struct {
	UserName    string `json:"userName"`    //用户账号
	Password    string `json:"password"`    // 密码
	CaptchaCode string `json:"captchaCode"` // 验证码
	CaptchaId   string `json:"captchaId"`   // 验证码ID
}

type CaptchaVo struct {
	CaptchaId     string `json:"captchaId"`
	CaptchaBase64 string `json:"captchaBase64"`
}

type UserInfoVo struct {
	Id         int64           `json:"id"`                  //编号
	UserName   string          `json:"userName"`            //用户账号
	NickName   string          `json:"nickName" `           //用户昵称
	UserType   consts.UserType `json:"userType"`            //用户类型（0普通账号，1超级管理员）
	Email      string          `json:"email"`               //用户邮箱
	Phone      *string         `json:"phone"`               //手机号码
	Gender     int             `json:"gender"`              //用户性别（0未知，1男，2女）
	Avatar     *string         `json:"avatar"`              //头像地址
	Status     int             `json:"status"`              //帐号状态（0正常 1停用）
	Remark     *string         `json:"remark"`              //备注
	Permission []string        `json:"permission" gorm:"-"` //权限
}

type AuthMenuVo struct {
	Path      string           `json:"path"`      //路由地址
	Name      string           `json:"name"`      //路由名称（TODO：暂时不清楚前端这个字段有什么用处）
	Meta      MenuRouterMetaVo `json:"meta"`      //Mate
	Component string           `json:"component"` //组件路径
	Children  []AuthMenuVo     `json:"children"`  //子菜单
}

type MenuRouterMetaVo struct {
	Title       string `json:"title"`       //菜单标题
	Icon        string `json:"icon"`        //菜单图标
	IsLink      string `json:"isLink"`      //外部地址
	IsIframe    bool   `json:"isIframe"`    //是否内嵌
	IsHide      bool   `json:"isHide"`      //是否隐藏
	IsKeepAlive bool   `json:"isKeepAlive"` //是否缓存
	IsAffix     bool   `json:"isAffix"`     //是否固定在 tagsView 栏上
}
