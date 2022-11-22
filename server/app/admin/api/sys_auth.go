package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/response"
)

type SysAuthApi struct{}

// GenerateCaptcha
// @Tags 授权
// @Summary 生成验证码
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=dto.CaptchaVo}
// @Router /captcha [get]
func (sysAuthApi *SysAuthApi) GenerateCaptcha(c *gin.Context) {
	id, b64s, err := authService.GenerateCaptcha()
	response.Complete(dto.CaptchaVo{
		CaptchaId:     id,
		CaptchaBase64: b64s,
	}, err, c)
}

// Login
// @Tags 授权
// @Summary 用户登录
// @Security ApiKeyAuth
// @Param data body dto.LoginDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /login [post]
func (sysAuthApi *SysAuthApi) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	c.ShouldBindJSON(&loginDto)
	token, err := authService.Login(loginDto)
	response.Complete(token, err, c)
}

// GetUserInfo
// @Tags 授权
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=dto.UserInfoVo}
// @Router /getUserInfo [get]
func (sysAuthApi *SysAuthApi) GetUserInfo(c *gin.Context) {
	userInfo, err := authService.GetUserInfo(c)
	response.Complete(userInfo, err, c)

}
