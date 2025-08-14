package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/common/tools"
	"go-fast-admin/server/common/utils"
	"go-fast-admin/server/global"
	"gopkg.in/errgo.v2/fmt/errors"
	"strconv"
	"time"
)

var store = base64Captcha.DefaultMemStore

type AuthService struct{}

// GenerateCaptcha 生成验证码
func (s *AuthService) GenerateCaptcha() (id string, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}

// Login 用户登录
func (s *AuthService) Login(loginDto dto.LoginDto) (token string, userId int64, err error) {

	if loginDto.CaptchaId == "" || loginDto.CaptchaCode == "" {
		return "", 0, errors.New("验证码错误")
	}

	if !store.Verify(loginDto.CaptchaId, loginDto.CaptchaCode, true) {
		return "", 0, errors.New("验证码错误")
	}

	var user model.SysUser

	err = global.DB.Where("user_name = ?", loginDto.UserName).First(&user).Error
	if err != nil {
		return "", 0, errors.New("账号不存在")
	}

	if user.Status == consts.UserStatusDisable {
		return "", 0, errors.New("账号已被禁用")
	}

	pwd := utils.Md5(loginDto.Password + user.Salt)
	if pwd != user.Password {
		return "", 0, errors.New("密码错误")
	}

	//生成token
	var claims = utils.UserAuthClaims{
		UserId:   user.Id,
		UserName: user.UserName,
		NickName: user.NickName,
		UserType: user.UserType,
	}
	token, err = utils.GenerateToken(claims, time.Now().AddDate(0, 0, 1))

	return token, user.Id, err
}

// GetUserInfo 获取用户信息
func (s *AuthService) GetUserInfo(c *gin.Context) (userInfo dto.UserInfoVo, err error) {
	currentUserId := utils.GetUserId(c)

	//查询用户信息
	err = global.DB.Model(&model.SysUser{}).Where("id = ?", currentUserId).Scan(&userInfo).Error
	if err != nil {
		return userInfo, err
	}

	//是否超级管理员
	if utils.IsSuperAdmin(c) {
		//所有权限
		userInfo.Permission = []string{"*_*_*"}
	} else {
		//用户权限
		userPermissions := GetUserPermission(currentUserId)
		for i := 0; i < len(userPermissions); i++ {
			userInfo.Permission = append(userInfo.Permission, userPermissions[i])
		}
	}
	return userInfo, err
}

// GetAllPermission 获取所有权限
func GetAllPermission() (permissions []string) {

	var cacheKey = consts.CacheKeySysAllPermission

	tools.Redis.Get(cacheKey, &permissions)
	if permissions == nil || len(permissions) == 0 {

		//所有权限
		global.DB.Model(&model.SysMenu{}).
			Select("permission").
			Where("menu_type = ? AND status = ? ", consts.MenuTypeBtn, consts.MenuStatusEnable).
			Scan(&permissions)

		//设置缓存
		tools.Redis.Set(cacheKey, permissions, consts.CacheTime)
	}
	return permissions
}

// GetUserPermission 获取用户权限
func GetUserPermission(userId int64) (permissions []string) {

	var cacheKey = consts.CacheKeySysUserPermission + strconv.FormatInt(userId, 10)

	tools.Redis.Get(cacheKey, &permissions)
	if permissions == nil || len(permissions) == 0 {

		var menuIds []int64

		//用户拥有菜单id
		global.DB.Table("sys_user_role userRole").Select("roleMenu.menu_id").
			Joins("LEFT JOIN sys_role_menu roleMenu ON userRole.role_id = roleMenu.role_id").
			Where("userRole.user_id = ?", userId).
			Group("roleMenu.menu_id").Find(&menuIds)

		//所有权限
		global.DB.Model(&model.SysMenu{}).
			Select("permission").
			Where("id IN ? AND menu_type = ? AND status = ? ", menuIds, consts.MenuTypeBtn, consts.MenuStatusEnable).
			Scan(&permissions)

		//设置缓存
		tools.Redis.Set(cacheKey, permissions, consts.CacheTime)
	}
	return permissions
}

// GetAuthMenu 获取用户菜单路由
func (s *AuthService) GetAuthMenu(c *gin.Context) (authMenu []dto.AuthMenuVo, err error) {

	var currentUserId = utils.GetUserId(c)

	var cacheKey = consts.CacheKeySysUserMenu + strconv.FormatInt(currentUserId, 10)

	tools.Redis.Get(cacheKey, &authMenu)

	var menuTypes = []consts.MenuType{consts.MenuTypeCatalog, consts.MenuTypeMenu}

	if authMenu == nil || len(authMenu) == 0 {
		var menuList []model.SysMenu
		//是否超级管理员
		if utils.IsSuperAdmin(c) {
			//超级管理员，拥有最高权限
			global.DB.Model(&model.SysMenu{}).Where("menu_type IN ? AND status = ? ",
				menuTypes,
				consts.MenuStatusEnable).Order("sort,id ASC").Scan(&menuList)
		} else {

			var menuIds []int64
			var userId = utils.GetUserId(c)
			//用户拥有菜单id
			global.DB.Table("sys_user_role userRole").Select("roleMenu.menu_id").
				Joins("LEFT JOIN sys_role_menu roleMenu ON userRole.role_id = roleMenu.role_id").
				Where("userRole.user_id = ?", userId).
				Group("roleMenu.menu_id").Find(&menuIds)

			global.DB.Model(&model.SysMenu{}).Where("id IN ? AND menu_type IN ? AND status = ? ",
				menuIds,
				menuTypes,
				consts.MenuStatusEnable).Order("sort,id ASC").Scan(&menuList)
		}
		authMenu = BuildAuthMenuTree(menuList, 0)

		//设置缓存
		tools.Redis.Set(cacheKey, authMenu, consts.CacheTime)
	}

	return authMenu, err
}

// BuildAuthMenuTree 构建菜单树状
func BuildAuthMenuTree(menuList []model.SysMenu, parentId int64) (menus []dto.AuthMenuVo) {

	var menuRouters []dto.AuthMenuVo
	for i := 0; i < len(menuList); i++ {
		var menu = menuList[i]
		if menu.ParentId == parentId {
			var menuRouter dto.AuthMenuVo
			menuRouter.Path = "/" + menu.Path
			menuRouter.Name = menu.Path
			menuRouter.Component = menu.Component
			menuRouter.Meta = dto.MenuRouterMetaVo{
				Title:       menu.MenuName,
				Icon:        menu.Icon,
				IsLink:      menu.Link,
				IsIframe:    menu.IsIframe,
				IsHide:      menu.IsHide,
				IsKeepAlive: menu.IsCache,
				IsAffix:     menu.IsAffix,
			}
			menuRouter.Children = BuildAuthMenuTree(menuList, menu.Id)
			menuRouters = append(menuRouters, menuRouter)
		}
	}
	return menuRouters
}

// UpdatePwd 修改密码
func (s *AuthService) UpdatePwd(c *gin.Context, updatePwdDto dto.UpdatePwdDto) error {
	var user model.SysUser
	currentUserId := utils.GetUserId(c)
	err := global.DB.Model(&model.SysUser{}).Where("id = ?", currentUserId).Scan(&user).Error
	if err != nil {
		return err
	}
	encryptPassword := utils.Md5(updatePwdDto.Password + user.Salt)
	if encryptPassword != user.Password {
		return errors.New("密码错误")
	}

	//密码盐
	salt := utils.GetRoundNumber(15)

	//加密 => MD5（密码+密码盐）
	pwd := utils.Md5(updatePwdDto.NewPassword + salt)

	//更新密码
	err = global.DB.Model(&model.SysUser{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"salt":     salt,
		"password": pwd,
	}).Error
	return err
}
