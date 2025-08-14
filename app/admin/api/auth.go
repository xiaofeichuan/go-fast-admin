package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/useragent"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/response"
	"go-fast-admin/server/common/utils"
)

type AuthApi struct{}

// GenerateCaptcha
// @Tags 授权
// @Summary 生成验证码
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=dto.CaptchaVo}
// @Router /system/auth/captcha [get]
func (*AuthApi) GenerateCaptcha(c *gin.Context) {
	id, b64s, err := authService.GenerateCaptcha()
	response.Complete(dto.CaptchaVo{CaptchaId: id, CaptchaBase64: b64s}, err, c)
}

// Login
// @Tags 授权
// @Summary 用户登录
// @Security ApiKeyAuth
// @Param data body dto.LoginDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/auth/login [post]
func (*AuthApi) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	c.ShouldBindJSON(&loginDto)
	token, userId, err := authService.Login(loginDto)
	if err == nil {
		ua := useragent.New(c.Request.UserAgent())
		browserName, browserVersion := ua.Browser()
		ip := c.ClientIP()

		//获取Ip信息
		info := utils.GetIPInfo(ip)

		loginLogService.Add(dto.SysLoginLogAddDto{
			UserId:   userId,
			Ip:       ip,
			Location: info,
			Browser:  browserName + "/" + browserVersion,
			OS:       ua.OS(),
		})
	}

	response.Complete(token, err, c)
}

// GetUserInfo
// @Tags 授权
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=dto.UserInfoVo}
// @Router /system/auth/userInfo [get]
func (*AuthApi) GetUserInfo(c *gin.Context) {
	userInfo, err := authService.GetUserInfo(c)
	response.Complete(userInfo, err, c)

}

// GetAuthMenu
// @Tags 授权
// @Summary 获取用户菜单（树状）
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.AuthMenuVo}
// @Router /system/auth/menu [get]
func (*AuthApi) GetAuthMenu(c *gin.Context) {
	authMenu, err := authService.GetAuthMenu(c)
	response.Complete(authMenu, err, c)
}

// UpdatePwd
// @Tags 授权
// @Summary 更新密码
// @Security ApiKeyAuth
// @Param data body dto.UpdatePwdDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/auth/updatePwd [post]
func (*AuthApi) UpdatePwd(c *gin.Context) {
	var updatePwdDto dto.UpdatePwdDto
	c.ShouldBindJSON(&updatePwdDto)
	err := authService.UpdatePwd(c, updatePwdDto)
	response.Complete(nil, err, c)
}
