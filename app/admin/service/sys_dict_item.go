package service

import (
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
)

type SysDictItemService struct{}

// Query 字典选项查询
func (s *SysDictItemService) Query(query dto.SysDictItemQuery) (list []dto.SysDictItemVo, total int64, err error) {
	db := global.DB.Model(&model.SysDictItem{})

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

// Add 添加字典选项
func (s *SysDictItemService) Add(addDto dto.SysDictItemAddDto) error {
	var role = &model.SysDictItem{
		DictId:        addDto.DictId,
		DictItemName:  addDto.DictItemName,
		DictItemValue: addDto.DictItemValue,
		Status:        addDto.Status,
		Remark:        addDto.Remark,
	}
	err := global.DB.Create(role).Error
	return err
}

// Update 更新字典选项
func (s *SysDictItemService) Update(updateDto dto.SysDictItemUpdateDto) error {
	err := global.DB.Model(&model.SysDictItem{}).Where("id = ?", updateDto.Id).Updates(map[string]interface{}{
		"DictId":        updateDto.DictId,
		"DictItemName":  updateDto.DictItemName,
		"DictItemValue": updateDto.DictItemValue,
		"Status":        updateDto.Status,
		"remark":        updateDto.Remark,
	}).Error
	return err
}

// Detail 获取字典选项详情
func (s *SysDictItemService) Detail(id int64) (obj dto.SysDictItemVo, err error) {
	err = global.DB.Model(&model.SysDictItem{}).Where("id = ?", id).Find(&obj).Error
	return obj, err
}

// Delete 删除字典选项
func (s *SysDictItemService) Delete(id int64) error {
	err := global.DB.Delete(&model.SysDictItem{}, "id = ?", id).Error
	return err
}

// List 字典选项列表
func (s *SysDictItemService) List(dictId int) (objs []dto.SysDictItemVo, err error) {
	err = global.DB.Model(&model.SysDictItem{}).Where("dict_id = ?", dictId).Find(&objs).Error
	return objs, err
}
