package service

import (
	"bytes"
	"fmt"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
	"go-fast-admin/server/utils"
	"strings"
	"text/template"
)

type SysGenTableService struct{}

// GetDBTableInfos 获取当前数据库所有表信息
func (sysGenTableService *SysGenTableService) GetDBTableInfos() (data []dto.TableInfoVo, err error) {
	var tables []dto.TableInfoVo
	sql := `
		SELECT 
			table_name AS TableName,
			table_comment AS TableComment,
			create_time AS CreateTime,
			update_time AS UpdateTime 
		FROM 
			information_schema.TABLES 
		WHERE 
			table_schema = (SELECT DATABASE()) AND table_name NOT IN (select table_name from sys_code_gen)`
	err = global.DB.Raw(sql).Scan(&tables).Error
	return tables, err
}

// ImportGentTable 导入表
func (sysGenTableService *SysGenTableService) ImportGentTable(tableNames []string) (tables []model.SysGenTable, err error) {
	var tableInfos []dto.TableInfoVo
	sql := `
		SELECT 
			table_name AS TableName,
			table_comment AS TableComment,
			create_time AS CreateTime,
			update_time AS UpdateTime 
		FROM 
			information_schema.TABLES 
		WHERE 
			table_schema = (SELECT DATABASE()) AND table_name IN ?`
	// 查询指定表
	err = global.DB.Raw(sql, tableNames).Scan(&tableInfos).Error
	if err != nil {
		return nil, err
	}

	for i := range tableInfos {
		tables = append(tables, model.SysGenTable{
			TableName:    tableInfos[i].TableName,
			TableComment: tableInfos[i].TableComment,
			ModelName:    utils.ToPascalCase(tableInfos[i].TableName),
			BusinessName: GetBusinessName(tableInfos[i].TableName),
			FunctionName: GetFunctionName(tableInfos[i].TableComment),
			ParamName:    utils.ToCamelCase(tableInfos[i].TableName),
		})
	}
	// 添加
	err = global.DB.Create(&tables).Error
	return tables, err
}

// PreviewCode 预览代码
func (sysGenTableService *SysGenTableService) PreviewCode(tableId uint64) (vos []dto.PreviewCodeVo, err error) {

	table := GetTableDetail(tableId)

	//api文件
	apiTemp, err := template.ParseFiles("resource/template/api.go.template")
	if err != nil {
		return nil, err
	}
	var apiContent bytes.Buffer
	err = apiTemp.Execute(&apiContent, table)
	vos = append(vos, dto.PreviewCodeVo{FileName: "api.go", FileContent: apiContent.String()})

	//dto文件
	dtoTemp, err := template.ParseFiles("resource/template/api.go.template")
	if err != nil {
		return nil, err
	}
	var dtoContent bytes.Buffer
	err = dtoTemp.Execute(&dtoContent, table)
	vos = append(vos, dto.PreviewCodeVo{FileName: "dto.go", FileContent: dtoContent.String()})

	//model文件
	modelTemp, err := template.ParseFiles("resource/template/model.go.template")
	if err != nil {
		return nil, err
	}
	var modelContent bytes.Buffer
	err = modelTemp.Execute(&modelContent, table)
	vos = append(vos, dto.PreviewCodeVo{FileName: "model.go", FileContent: modelContent.String()})

	//router文件
	routerTemp, err := template.ParseFiles("resource/template/model.go.template")
	if err != nil {
		return nil, err
	}
	var routerContent bytes.Buffer
	err = routerTemp.Execute(&routerContent, table)
	vos = append(vos, dto.PreviewCodeVo{FileName: "router.go", FileContent: routerContent.String()})

	//service文件
	serviceTemp, err := template.ParseFiles("resource/template/model.go.template")
	if err != nil {
		return nil, err
	}
	var serviceContent bytes.Buffer
	err = serviceTemp.Execute(&serviceContent, table)
	vos = append(vos, dto.PreviewCodeVo{FileName: "service.go", FileContent: serviceContent.String()})

	return vos, err
}

// GetTableDetail 表数据详情
func GetTableDetail(tableId uint64) (table dto.TableInfoVo) {
	global.DB.Model(model.SysGenTable{Id: tableId}).Scan(&table)

	var columnList *[]model.SysGenTableColumn
	global.DB.Model(model.SysGenTableColumn{TableId: tableId}).Scan(&columnList)
	table.ParamName = utils.ToCamelCase(table.TableName)
	table.ColumnList = columnList
	fmt.Println(table)
	return table

}

// GetFunctionName 获取功能名称 例如：用户信息表=>用户信息
func GetFunctionName(tableComment string) string {
	if strings.Contains(tableComment, "表") {
		tableComment = strings.Replace(tableComment, "表", "", -1)
	}
	return tableComment
}

// GetBusinessName 获取业务名称 例如：sys_user=>user
func GetBusinessName(tableName string) string {
	var nameArr []string
	nameList := strings.Split(tableName, "_")

	for i := 1; i < len(nameList); i++ {
		nameArr = append(nameArr, nameList[i])
	}
	snakeCase := strings.Join(nameArr, "_")
	return utils.ToCamelCase(snakeCase)
}
