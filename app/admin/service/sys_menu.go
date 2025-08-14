package service

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/common/tools"
	"go-fast-admin/server/global"
	"gopkg.in/errgo.v2/fmt/errors"
)

type SysMenuService struct{}

// Query 角色分页查询
func (s *SysMenuService) Query(query dto.SysMenuQuery) (list []dto.SysRoleVo, total int64, err error) {
	db := global.DB.Model(&model.SysMenu{})
	//查询条件
	if query.MenuName != "" {
		db = db.Where("`role_name` LIKE ?", "%"+query.MenuName+"%")
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
func (s *SysMenuService) Add(addDto dto.SysMenuAddDto) error {

	if addDto.MenuType == consts.MenuTypeMenu {
		if addDto.Path == "" {
			return errors.New("路由地址不允许为空")
		}
		var menu = &model.SysMenu{}
		global.DB.Where("path = ?", addDto.Path).First(&menu)
		if menu.Id != 0 {
			return errors.New("路由地址不允许重复")
		}
	}

	var menu = &model.SysMenu{
		ParentId:   addDto.ParentId,
		MenuName:   addDto.MenuName,
		MenuType:   addDto.MenuType,
		Path:       addDto.Path,
		Link:       addDto.Link,
		IsIframe:   addDto.IsIframe,
		IsHide:     addDto.IsHide,
		IsCache:    addDto.IsCache,
		Component:  addDto.Component,
		Icon:       addDto.Icon,
		Permission: addDto.Permission,
		Status:     addDto.Status,
		Sort:       addDto.Sort,
		Remark:     addDto.Remark,
	}
	err := global.DB.Create(menu).Error

	//删除缓存
	tools.Redis.DelByPattern(consts.CacheKeySysUserMenu)
	tools.Redis.DelByPattern(consts.CacheKeySysUserPermission)
	return err
}

// Update 更新角色
func (s *SysMenuService) Update(updateDto dto.SysMenuUpdateDto) error {

	if updateDto.MenuType == consts.MenuTypeMenu {
		if updateDto.Path == "" {
			return errors.New("路由地址不允许为空")
		}
		var menu = model.SysMenu{}
		global.DB.Where("id != ? AND path = ?", updateDto.Id, updateDto.Path).First(&menu)
		if menu.Id != 0 {
			return errors.New("路由地址不允许重复")
		}
	}

	err := global.DB.Model(&model.SysMenu{}).Where("id = ?", updateDto.Id).Updates(map[string]interface{}{
		"parent_id":  updateDto.ParentId,
		"menu_name":  updateDto.MenuName,
		"menu_type":  updateDto.MenuType,
		"path":       updateDto.Path,
		"link":       updateDto.Link,
		"is_iframe":  updateDto.IsIframe,
		"is_hide":    updateDto.IsHide,
		"is_cache":   updateDto.IsCache,
		"is_affix":   updateDto.IsAffix,
		"component":  updateDto.Component,
		"icon":       updateDto.Icon,
		"permission": updateDto.Permission,
		"status":     updateDto.Status,
		"sort":       updateDto.Sort,
		"remark":     updateDto.Remark,
	}).Error

	//删除缓存
	tools.Redis.DelByPattern(consts.CacheKeySysUserMenu)
	tools.Redis.DelByPattern(consts.CacheKeySysUserPermission)
	return err
}

// Detail 获取角色详情
func (s *SysMenuService) Detail(id int64) (obj dto.SysMenuVo, err error) {
	err = global.DB.Model(&model.SysMenu{}).Where("id = ?", id).Find(&obj).Error
	return obj, err
}

// Delete 删除角色
func (s *SysMenuService) Delete(id int64) error {
	err := global.DB.Delete(&model.SysMenu{}, "id = ?", id).Error

	//删除缓存
	tools.Redis.DelByPattern(consts.CacheKeySysUserMenu)
	tools.Redis.DelByPattern(consts.CacheKeySysUserPermission)
	return err
}

// GetMenuTree 获取菜单树状
func (s *SysMenuService) GetMenuTree() (objs []dto.SysMenuVo, err error) {

	err = global.DB.Model(&model.SysMenu{}).Order("sort,id ASC").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	tree := BuildMenuTree(objs, 0)
	return tree, err
}

// BuildMenuTree 构建菜单树状
func BuildMenuTree(menuList []dto.SysMenuVo, parentId int64) (menus []dto.SysMenuVo) {
	var menuTree []dto.SysMenuVo
	for i := 0; i < len(menuList); i++ {
		var menu = menuList[i]
		if menu.ParentId == parentId {
			menu.Children = BuildMenuTree(menuList, menu.Id)
			menuTree = append(menuTree, menu)
		}
	}
	return menuTree
}
