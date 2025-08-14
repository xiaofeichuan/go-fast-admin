package service

import (
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
	"strconv"
)

type SysDictService struct{}

// Query 字典查询
func (s *SysDictService) Query(query dto.SysDictQuery) (list []dto.SysDictVo, total int64, err error) {
	db := global.DB.Model(&model.SysDict{})
	//查询条件
	if query.DictName != "" {
		db = db.Where("dict_name LIKE ?", "%"+query.DictName+"%")
	}
	if query.DictName != "" {
		db = db.Where("dict_name LIKE ?", "%"+query.DictName+"%")
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

// Add 添加字典
func (s *SysDictService) Add(addDto dto.SysDictAddDto) error {
	var dict = &model.SysDict{
		DictName: addDto.DictName,
		DictCode: addDto.DictCode,
		DictType: addDto.DictType,
		Status:   addDto.Status,
		Remark:   addDto.Remark,
	}
	err := global.DB.Create(dict).Error
	return err
}

// Update 更新字典
func (s *SysDictService) Update(updateDto dto.SysDictUpdateDto) error {
	err := global.DB.Model(&model.SysDict{}).Where("id = ?", updateDto.Id).Updates(map[string]interface{}{
		"DictName": updateDto.DictName,
		"DictCode": updateDto.DictCode,
		"DictType": updateDto.DictType,
		"Status":   updateDto.Status,
		"Remark":   updateDto.Remark,
	}).Error
	return err
}

// Delete 删除字典
func (s *SysDictService) Delete(id int64) error {
	err := global.DB.Delete(&model.SysDict{}, "id = ?", id).Error
	return err
}

// Detail 获取字典详情
func (s *SysDictService) Detail(id int64) (obj dto.SysDictVo, err error) {
	err = global.DB.Model(&model.SysDict{}).Where("id = ?", id).Find(&obj).Error
	return obj, err
}

// List 字典列表
func (s *SysDictService) List() (objs []dto.SysDictVo, err error) {
	err = global.DB.Model(&model.SysDict{}).Where("status = ?", consts.DictStatusEnable).Find(&objs).Error
	var dictItems []dto.SysDictItemVo
	err = global.DB.Model(&model.SysDictItem{}).Where("status = ?", consts.DictItemStatusEnable).Find(&dictItems).Error

	for i := 0; i < len(objs); i++ {

		for _, dictItem := range dictItems {
			if objs[i].Id == dictItem.DictId {

				if objs[i].DictType == consts.DictTypeInt {
					//数字
					num, _ := strconv.Atoi(dictItem.DictItemValue)
					objs[i].Items = append(objs[i].Items, dto.DictItemVo{
						DictItemName:  dictItem.DictItemName,
						DictItemValue: num,
					})
				} else {
					//字符串
					objs[i].Items = append(objs[i].Items, dto.DictItemVo{
						DictItemName:  dictItem.DictItemName,
						DictItemValue: dictItem.DictItemValue,
					})

				}

			}
		}
	}
	return objs, err
}
