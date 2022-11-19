package service

import (
	"fmt"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/app/admin/model"
	"go-fast-admin/server/global"
	"go-fast-admin/server/utils"
)

type SysGenTableColumnService struct{}

// ImportGenTableColumn 导入表字段
func (sysGenTableColumnService *SysGenTableColumnService) ImportGenTableColumn(tableNames []string, tables []model.SysGenTable) error {

	var columnInfos []dto.TableColumnInfoVo
	columnSql := `		
	SELECT
		table_name AS TableName,
		column_name AS ColumnName,
		column_comment AS ColumnComment,
		column_type AS ColumnType,
		( CASE WHEN column_key = 'PRI' THEN '1' ELSE '0' END ) AS IsPk 
	FROM
		information_schema.COLUMNS 
	WHERE
		table_schema = (SELECT DATABASE()) AND table_name IN ? ORDER BY TABLE_NAME,ORDINAL_POSITION`
	err := global.DB.Raw(columnSql, tableNames).Scan(&columnInfos).Error
	if err != nil {
		return err
	}
	//导入表数据
	var columns []model.SysGenTableColumn
	for i := range columnInfos {
		table := model.SysGenTable{}

		// 循环过滤（TODO：暂无更好的Linq通用库，手写过滤）
		for j := 0; j < len(tables); j++ {
			if tables[j].TableName == columnInfos[i].TableName {
				table = tables[j]
				break
			}
		}
		fmt.Println("开始TABLE:")
		fmt.Println(table)
		columns = append(columns, model.SysGenTableColumn{
			TableId:       table.Id,
			ColumnName:    columnInfos[i].ColumnName,
			ColumnComment: columnInfos[i].ColumnComment,
			ColumnType:    columnInfos[i].ColumnType,
			ParamName:     utils.ToCamelCase(columnInfos[i].ColumnName),
			GoField:       utils.ToPascalCase(columnInfos[i].ColumnName),
			GoType:        utils.GetGoType(columnInfos[i].ColumnName, columnInfos[i].ColumnType),
			IsPk:          columnInfos[i].IsPk,
			IsList:        true,
		})
	}
	err = global.DB.Create(&columns).Error
	return err
}
