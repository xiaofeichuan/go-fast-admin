package service

import (
	"gin-fast-admin/server/app/admin/dto"
	"gin-fast-admin/server/app/admin/model"
	"gin-fast-admin/server/global"
)

type SysUserService struct{}

// Page 用户分页查询
func (sysUserService *SysUserService) Page(query dto.SysUserQuery) (list []dto.SysUserVo, total int64, err error) {
	db := global.DB.Model(&model.SysUser{})
	//查询条件
	if query.NickName != "" {
		db = db.Where("`nick_name` LIKE ?", "%"+query.NickName+"%")
	}
	offset := (query.PageIndex - 1) * query.PageSize
	//查询数据
	err = db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	//总条数
	db.Count(&total)
	return list, total, err
}

// Detail 获取用户详情
func (sysUserService *SysUserService) Detail(id uint64) (obj dto.SysUserVo, err error) {
	err = global.DB.Model(&model.SysUser{}).Where("id = ?", id).Scan(&obj).Error
	return obj, err
}

// List 用户列表
func (sysUserService *SysUserService) List() (objs []dto.SysUserVo, err error) {
	err = global.DB.Model(&model.SysUser{}).Scan(&objs).Error
	return objs, err
}

// Add 添加用户
func (sysUserService *SysUserService) Add(createDto dto.SysUserCreateDto) error {
	var user = &model.SysUser{
		Account:  createDto.Account,
		NickName: createDto.NickName,
		UserType: createDto.UserType,
		Email:    createDto.Email,
		Phone:    createDto.Phone,
		Password: createDto.Password,
		Sex:      createDto.Sex,
		Avatar:   createDto.Avatar,
		Status:   createDto.Status,
		Remark:   createDto.Remark,
	}
	err := global.DB.Create(user).Error
	return err
}

// Update 更新用户
func (sysUserService *SysUserService) Update(updateDto dto.SysUserUpdateDto) error {
	var user = &model.SysUser{
		Id:       updateDto.Id,
		Account:  updateDto.Account,
		NickName: updateDto.NickName,
		UserType: updateDto.UserType,
		Email:    updateDto.Email,
		Phone:    updateDto.Phone,
		Password: updateDto.Password,
		Sex:      updateDto.Sex,
		Avatar:   updateDto.Avatar,
		Status:   updateDto.Status,
		Remark:   updateDto.Remark,
	}
	err := global.DB.Updates(user).Error
	return err
}

// Delete 删除用户
func (sysUserService *SysUserService) Delete(id uint64) error {
	err := global.DB.Delete(&model.SysUser{}, "id = ?", id).Error
	return err
}
