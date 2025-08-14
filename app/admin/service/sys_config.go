package service

import (
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
)

type SysConfigService struct{}

// Query 配置查询
func (s *SysConfigService) Query(query dto.SysConfigQuery) (list []dto.SysConfigVo, total int64, err error) {
	db := global.DB.Model(&model.SysConfig{})
	//查询条件
	if query.ConfigName != "" {
		db = db.Where("`config_name` LIKE ?", "%"+query.ConfigName+"%")
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

// Add 添加配置
func (s *SysConfigService) Add(addDto dto.SysConfigAddDto) error {
	var role = &model.SysConfig{
		ConfigName:  addDto.ConfigName,
		ConfigKey:   addDto.ConfigKey,
		ConfigValue: addDto.ConfigValue,
		Status:      addDto.Status,
		Remark:      addDto.Remark,
	}
	err := global.DB.Create(role).Error
	return err
}

// Update 更新配置
func (s *SysConfigService) Update(updateDto dto.SysConfigUpdateDto) error {
	err := global.DB.Model(&model.SysConfig{}).Where("id = ?", updateDto.Id).Updates(map[string]interface{}{
		"config_name":  updateDto.ConfigName,
		"config_key":   updateDto.ConfigKey,
		"config_value": updateDto.ConfigValue,
		"status":       updateDto.Status,
		"remark":       updateDto.Remark,
	}).Error
	return err
}

// Detail 获取配置详情
func (s *SysConfigService) Detail(id int64) (obj dto.SysConfigVo, err error) {
	err = global.DB.Model(&model.SysConfig{}).Where("id = ?", id).Scan(&obj).Error
	return obj, err
}

// Delete 删除配置
func (s *SysConfigService) Delete(id int64) error {
	err := global.DB.Delete(&model.SysConfig{}, "id = ?", id).Error
	return err
}
