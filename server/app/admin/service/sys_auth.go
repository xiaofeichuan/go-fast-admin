package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
	"go-fast-admin/server/utils"
	"gopkg.in/errgo.v2/fmt/errors"
	"time"
)

var store = base64Captcha.DefaultMemStore

type SysAuthService struct{}

// GenerateCaptcha 生成验证码
func (sysAuthService *SysAuthService) GenerateCaptcha() (id string, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}

// Login 用户登录
func (sysAuthService *SysAuthService) Login(loginDto dto.LoginDto) (token string, err error) {

	if !store.Verify(loginDto.CaptchaId, loginDto.Captcha, true) {
		return "", errors.New("验证码错误")
	}

	var user model.SysUser

	err = global.DB.Where("account = ?", loginDto.Account).First(&user).Error
	if err != nil {
		return "", errors.New("账号不存在")
	}

	pwd := utils.Md5(loginDto.Password + user.Salt)
	if pwd != user.Password {
		return "", errors.New("密码错误")
	}

	//生成token
	var claims = utils.UserAuthClaims{
		UserId:   user.Id,
		Account:  user.Account,
		UserName: user.UserName,
		UserType: user.UserType,
	}
	token, err = utils.GenerateToken(claims, time.Now().AddDate(0, 0, 1))
	return token, err
}

// GetUserInfo 获取用户信息
func (sysAuthService *SysAuthService) GetUserInfo(c *gin.Context) (userInfo dto.UserInfoVo, err error) {
	currentUserId := utils.GetCurrentUserId(c)
	err = global.DB.Model(&model.SysUser{}).Where("id = ?", currentUserId).Scan(&userInfo).Error
	return userInfo, err
}
