package service

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/common/tools"
	utils2 "go-fast-admin/server/common/utils"
	"go-fast-admin/server/global"
	"gopkg.in/errgo.v2/fmt/errors"
	"strconv"
)

type SysUserService struct{}

// Query 用户分页查询
func (s *SysUserService) Query(query dto.SysUserQuery) (list []dto.SysUserVo, total int64, err error) {
	db := global.DB.Model(&model.SysUser{})
	//查询条件
	if query.UserName != "" {
		value := "%" + query.UserName + "%"
		db = db.Where("nick_name LIKE ? or user_name like ? or phone like ?", value, value, value)
	}

	//总条数
	db.Count(&total)

	offset := (query.PageNum - 1) * query.PageSize
	//查询数据
	err = db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, err
}

// Add 添加用户
func (s *SysUserService) Add(addDto dto.SysUserAddDto) error {

	if addDto.UserName == "" {
		return errors.New("账号名称不允许为空")
	}

	var count int64
	global.DB.Model(&model.SysUser{}).Where("user_name = ?", addDto.UserName).Count(&count)
	if count != 0 {
		return errors.New("账号名称已存在")
	}

	//密码盐
	var salt = utils2.GetRoundNumber(15)

	//加密 => MD5（密码+密码盐）
	var pwd = utils2.Md5(addDto.Password + salt)

	var user = &model.SysUser{
		UserName: addDto.UserName,
		NickName: addDto.NickName,
		UserType: 0,
		Email:    addDto.Email,
		Phone:    addDto.Phone,
		Salt:     salt,
		Password: pwd,
		Gender:   addDto.Gender,
		Status:   addDto.Status,
		Remark:   addDto.Remark,
	}
	err := global.DB.Create(user).Error
	if err != nil {
		return err
	}
	err = SetUserRole(user.Id, addDto.RoleIds)
	return err
}

// Update 更新用户
func (s *SysUserService) Update(updateDto dto.SysUserUpdateDto) error {
	err := global.DB.Model(&model.SysUser{}).Where("id = ?", updateDto.Id).Updates(map[string]interface{}{
		"nick_name": updateDto.NickName,
		"user_type": 0,
		"email":     updateDto.Email,
		"phone":     updateDto.Phone,
		"gender":    updateDto.Gender,
		"status":    updateDto.Status,
		"remark":    updateDto.Remark,
	}).Error
	if err != nil {
		return err
	}
	err = SetUserRole(updateDto.Id, updateDto.RoleIds)
	return err
}

// Detail 获取用户详情
func (s *SysUserService) Detail(id int64) (obj *dto.SysUserVo, err error) {
	err = global.DB.Model(&model.SysUser{}).Where("id = ?", id).Scan(&obj).Error
	if err != nil {
		return nil, err
	}
	var roleIds []int64
	err = global.DB.Model(&model.SysUserRole{}).Select("role_id").Where("user_id = ?", id).Scan(&roleIds).Error
	if err != nil {
		return nil, err
	}
	obj.RoleIds = roleIds
	return obj, err
}

// Delete 删除用户
func (s *SysUserService) Delete(id int64) error {
	err := global.DB.Delete(&model.SysUser{}, "id = ?", id).Error
	return err
}

// ResetPwd 重置密码
func (s *SysUserService) ResetPwd(reqDto dto.ResetPwdDto) error {

	//密码盐
	salt := utils2.GetRoundNumber(15)

	//加密 => MD5（密码+密码盐）
	pwd := utils2.Md5(reqDto.Password + salt)

	//更新密码
	err := global.DB.Model(&model.SysUser{}).Where("id = ?", reqDto.UserId).Updates(map[string]interface{}{
		"salt":     salt,
		"password": pwd,
	}).Error
	return err
}

// SetUserRole 设置用户角色
func SetUserRole(userId int64, roleIds []int64) error {

	err := global.DB.Delete(&model.SysUserRole{}, "user_id = ?", userId).Error
	if err != nil {
		return err
	}
	if len(roleIds) == 0 {
		return nil
	}
	var userRoles []model.SysUserRole
	for i := 0; i < len(roleIds); i++ {
		userRole := model.SysUserRole{
			UserId: userId,
			RoleId: roleIds[i],
		}
		userRoles = append(userRoles, userRole)
	}
	err = global.DB.Create(userRoles).Error

	//删除缓存
	userIdStr := strconv.FormatInt(userId, 10)
	tools.Redis.Del(consts.CacheKeySysUserMenu + userIdStr)
	tools.Redis.Del(consts.CacheKeySysUserPermission + userIdStr)

	return err

}

// SetUserStatus 设置用户状态
func (s *SysUserService) SetUserStatus(reqDto dto.SetUserStateDto) error {

	//用户状态
	err := global.DB.Model(&model.SysUser{}).Where("id = ?", reqDto.UserId).Update("status", reqDto.Status).Error
	return err

}
