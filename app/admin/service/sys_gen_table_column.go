package service

import (
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
)

type SysGenTableColumnService struct{}

// Update 更新业务表
func (s *SysGenTableColumnService) Update(updateDto []dto.SysGenTableColumnUpdateDto) {
	for i := 0; i < len(updateDto); i++ {
		global.DB.Model(&model.SysGenTableColumn{}).Where("id = ?", updateDto[i].Id).Updates(map[string]interface{}{
			"ColumnName":        updateDto[i].ColumnName,
			"ColumnDescription": updateDto[i].ColumnDescription,
			"ColumnType":        updateDto[i].ColumnType,
			"ParamName":         updateDto[i].ParamName,
			"CodeField":         updateDto[i].CodeField,
			"CodeType":          updateDto[i].CodeType,
			"IsPk":              updateDto[i].IsPk,
			"IsEdit":            updateDto[i].IsEdit,
			"IsList":            updateDto[i].IsList,
			"IsQuery":           updateDto[i].IsQuery,
			"QueryMethod":       updateDto[i].QueryMethod,
			"ComponentType":     updateDto[i].ComponentType,
			"DictCode":          updateDto[i].DictCode,
		})
	}

}

// List 业务表列表
func (s *SysGenTableColumnService) List(tableId int64) (objs []dto.SysGenTableColumnVo, err error) {
	err = global.DB.Model(&model.SysGenTableColumn{}).Where("table_id = ?", tableId).Find(&objs).Error
	return objs, err
}
