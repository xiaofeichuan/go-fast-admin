package service

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/common/tools"
	"go-fast-admin/server/global"
	"gopkg.in/errgo.v2/fmt/errors"
)

type SysRoleService struct{}

// Query 角色分页查询
func (s *SysRoleService) Query(query dto.SysRoleQuery) (list []dto.SysRoleVo, total int64, err error) {
	db := global.DB.Model(&model.SysRole{})
	//查询条件
	if query.RoleName != "" {
		db = db.Where("`role_name` LIKE ?", "%"+query.RoleName+"%")
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

// Add 添加角色
func (s *SysRoleService) Add(addDto dto.SysRoleAddDto) error {
	var role = &model.SysRole{
		RoleName: addDto.RoleName,
		RoleCode: addDto.RoleCode,
		Status:   addDto.Status,
		Sort:     addDto.Sort,
		Remark:   addDto.Remark,
	}
	err := global.DB.Create(role).Error
	if err != nil {
		return err
	}

	//设置角色菜单
	err = SetRoleMenu(role.Id, addDto.MenuIds)

	//删除缓存
	tools.Redis.DelByPattern(consts.CacheKeySysUserMenu)
	tools.Redis.DelByPattern(consts.CacheKeySysUserPermission)

	return err
}

// Update 更新角色
func (s *SysRoleService) Update(updateDto dto.SysRoleUpdateDto) error {
	err := global.DB.Model(&model.SysRole{}).Where("id = ?", updateDto.Id).Updates(map[string]interface{}{
		"role_name": updateDto.RoleName,
		"role_code": updateDto.RoleCode,
		"sort":      updateDto.Sort,
		"status":    updateDto.Status,
		"remark":    updateDto.Remark,
	}).Error
	if err != nil {
		return err
	}

	//设置角色菜单
	err = SetRoleMenu(updateDto.Id, updateDto.MenuIds)

	//删除缓存
	tools.Redis.DelByPattern(consts.CacheKeySysUserMenu)
	tools.Redis.DelByPattern(consts.CacheKeySysUserPermission)

	return err
}

// Detail 获取角色详情
func (s *SysRoleService) Detail(id int64) (obj *dto.SysRoleVo, err error) {
	err = global.DB.Model(&model.SysRole{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	var menuIds []int64
	err = global.DB.Model(&model.SysRoleMenu{}).Select("menu_id").Where("role_id = ?", id).Order("id ASC").Scan(&menuIds).Error
	if err != nil {
		return nil, err
	}
	obj.MenuIds = menuIds
	return obj, err
}

// Delete 删除角色
func (s *SysRoleService) Delete(id int64) error {
	var total int64
	global.DB.Model(&model.SysRoleMenu{}).Where("role_id = ?", id).Count(&total)
	if total > 0 {
		return errors.New("该角色与菜单存在关联，禁止删除")
	}
	global.DB.Model(&model.SysUserRole{}).Where("role_id = ?", id).Count(&total)
	if total > 0 {
		return errors.New("该角色与用户存在关联，禁止删除")
	}

	err := global.DB.Delete(&model.SysRole{}, "id = ?", id).Error

	//删除缓存
	tools.Redis.DelByPattern(consts.CacheKeySysUserMenu)
	tools.Redis.DelByPattern(consts.CacheKeySysUserPermission)

	return err
}

// List 角色列表
func (s *SysRoleService) List() (objs []dto.SysRoleVo, err error) {
	err = global.DB.Model(&model.SysRole{}).Where("status = ?", consts.RoleStatusEnable).Scan(&objs).Error
	return objs, err
}

// SetRoleMenu 设置角色菜单
func SetRoleMenu(roleId int64, menuIds []int64) error {

	err := global.DB.Delete(&model.SysRoleMenu{}, "role_id = ?", roleId).Error
	if err != nil {
		return err
	}
	if len(menuIds) == 0 {
		return nil
	}
	var roleMenus []model.SysRoleMenu
	for i := 0; i < len(menuIds); i++ {
		roleMenu := model.SysRoleMenu{
			RoleId: roleId,
			MenuId: menuIds[i],
		}
		roleMenus = append(roleMenus, roleMenu)
	}
	err = global.DB.Create(roleMenus).Error
	return err
}
