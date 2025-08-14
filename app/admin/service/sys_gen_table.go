package service

import (
	"bytes"
	"go-fast-admin/server/app/admin/consts"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/common/tools"
	"go-fast-admin/server/common/utils"
	"go-fast-admin/server/global"
	"path"
	"strings"
	"text/template"
)

type SysGenTableService struct{}

// Query 业务表查询
func (s *SysGenTableService) Query(query dto.SysGenTableQuery) (list []dto.SysGenTableVo, total int64, err error) {
	db := global.DB.Model(&model.SysGenTable{})
	//查询条件
	if query.TableName != "" {
		db = db.Where("`table_name` LIKE ?", "%"+query.TableName+"%")
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

// Add 添加业务表
func (s *SysGenTableService) Add(addDto dto.SysGenTableAddDto) error {
	var genTable = &model.SysGenTable{
		TableName:        addDto.TableName,
		TableDescription: addDto.TableDescription,
		ModelName:        addDto.ModelName,
		BusinessName:     addDto.BusinessName,
		ModuleName:       addDto.ModuleName,
		MenuParentId:     addDto.MenuParentId,
	}
	err := global.DB.Create(genTable).Error

	//添加表字段
	s.AddColumn(addDto.TableName, genTable.Id)

	return err
}

// Update 更新业务表
func (s *SysGenTableService) Update(updateDto dto.SysGenTableUpdateDto) error {
	err := global.DB.Model(&model.SysGenTable{}).Where("id = ?", updateDto.Id).Updates(map[string]interface{}{
		"TableName":        updateDto.TableName,
		"TableDescription": updateDto.TableDescription,
		"ModelName":        updateDto.ModelName,
		"BusinessName":     updateDto.BusinessName,
		"ModuleName":       updateDto.ModuleName,
		"MenuParentId":     updateDto.MenuParentId,
	}).Error
	return err
}

// Delete 删除业务表
func (s *SysGenTableService) Delete(id int64) error {
	err := global.DB.Delete(&model.SysGenTable{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	//删除字段
	err = global.DB.Delete(&model.SysGenTableColumn{}, "table_id = ?", id).Error
	return err
}

// Detail 获取业务表详情
func (s *SysGenTableService) Detail(id int64) (obj dto.SysGenTableVo, err error) {
	err = global.DB.Model(&model.SysGenTable{}).Where("id = ?", id).Scan(&obj).Error

	//模块别名
	obj.ModuleAlias = strings.ToUpper(obj.ModuleName[:1]) + obj.ModuleName[1:]
	var columnList []dto.SysGenTableColumnVo

	//字段列表
	global.DB.Model(&model.SysGenTableColumn{}).Where("table_id = ?", id).Scan(&columnList)
	obj.ColumnList = columnList
	return obj, err
}

// GetTableList 获取表列表
func (s *SysGenTableService) GetTableList() (genTables []dto.TableInfoVo, err error) {

	//获取表信息
	tableInfos, err := tools.Gorm.Database().GetTableInfoList()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tableInfos); i++ {

		//业务名称
		var businessName = strings.Split(tableInfos[i].TableName, "_")[0]
		if businessName == "sys" {
			businessName = "system"
		}

		nameList := strings.Split(tableInfos[i].TableName, "_")[1:]
		snakeCase := strings.Join(nameList, "_")

		//表名称
		var tableName = tableInfos[i].TableName
		if tableInfos[i].SchemaName != "" {
			//pgsql 表名方式
			tableName = tableInfos[i].SchemaName + "." + tableInfos[i].TableName
		}
		genTables = append(genTables, dto.TableInfoVo{
			TableName:        tableName,
			TableDescription: tableInfos[i].TableDescription,
			ModelName:        utils.SnakeToUpperCamelCase(tableInfos[i].TableName),
			BusinessName:     businessName,
			ModuleName:       utils.SnakeToLowerCamelCase(snakeCase),
		})
	}
	return genTables, err
}

// GetFunctionName 获取功能名称 例如：用户信息表=>用户信息
func (s *SysGenTableService) GetFunctionName(tableDescription string) string {
	if strings.Contains(tableDescription, "表") {
		tableDescription = strings.Replace(tableDescription, "表", "", -1)
	}
	return tableDescription
}

// AddColumn 导入表字段
func (s *SysGenTableService) AddColumn(tableName string, tableId int64) error {

	columnInfos, err := tools.Gorm.Database().GetColumnInfoList(tableName)
	if err != nil {
		return err
	}

	//导入表数据
	var columns []model.SysGenTableColumn
	for _, column := range columnInfos {

		//是否编辑
		isEdit := true
		if column.ColumnName == "id" || column.ColumnName == "created_at" || column.ColumnName == "updated_at" || column.ColumnName == "deleted_at" {
			isEdit = false
		}

		//代码类型
		var codeType = ""
		if strings.Contains(column.ColumnName, "is_") {
			codeType = "bool"
		} else {
			codeType = utils.GetGoType(column.ColumnType)
		}

		//组件类型
		var componentType = consts.ComponentTypeInput
		var queryMethod = "=="
		if codeType == "bool" {
			componentType = consts.ComponentTypeSwitch
		} else if codeType == "time.Time" {
			componentType = consts.ComponentTypeDateTimePicker
			queryMethod = "BETWEEN"
		} else if codeType == "int" {
			componentType = consts.ComponentTypeInputNumber
		}

		columns = append(columns, model.SysGenTableColumn{
			TableId:           tableId,
			ColumnName:        column.ColumnName,
			ColumnDescription: column.ColumnDescription,
			ColumnType:        column.ColumnType,
			ParamName:         utils.SnakeToLowerCamelCase(column.ColumnName),
			CodeField:         utils.SnakeToUpperCamelCase(column.ColumnName),
			CodeType:          codeType,
			IsPk:              column.IsPk,
			IsEdit:            isEdit,
			IsList:            true,
			IsQuery:           false,
			QueryMethod:       queryMethod,
			ComponentType:     componentType,
		})
	}
	err = global.DB.Create(&columns).Error
	return err
}

// PreviewCode 预览代码
func (s *SysGenTableService) PreviewCode(tableId int64) (vos []dto.PreviewCodeVo, err error) {

	table, err := s.Detail(tableId)

	templatePaths := GetTemplatePath()

	for _, templatePath := range templatePaths {
		temp, err := template.ParseFiles(templatePath)
		if err != nil {
			return nil, err
		}
		var apiContent bytes.Buffer
		err = temp.Execute(&apiContent, table)
		var fileName = path.Base(templatePath)

		vos = append(vos, dto.PreviewCodeVo{FileName: strings.Replace(fileName, ".template", "", -1), FileContent: apiContent.String()})
	}

	return vos, err
}

// GetTemplatePath 获取模板文件路径
func GetTemplatePath() []string {
	templatePath := []string{
		"resource/template/admin/model.go.template",
		"resource/template/admin/api.go.template",
		"resource/template/admin/service.go.template",
		"resource/template/admin/router.go.template",
		"resource/template/admin/dto.go.template",
		"resource/template/admin/menu.sql.template",
		"resource/template/web/index.vue.template",
		"resource/template/web/editDialog.vue.template",
		"resource/template/web/api.ts.template",
	}
	return templatePath
}
