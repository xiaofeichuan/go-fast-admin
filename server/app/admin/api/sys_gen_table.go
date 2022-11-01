package api

import (
	"fmt"
	"gin-fast-admin/server/common/dto/response"
	"github.com/gin-gonic/gin"
)

type SysTableApi struct{}

// GetDBTableInfos
// @Summary 获取当前数据库所有表信息
// @Tags 代码生成器
// @Success 200 {object} response.JsonResult{data=[]dto.TableInfoResponse}
// @Security ApiKeyAuth
// @Router /sysTable/getDBTableInfos [get]
func (sysTableApi *SysTableApi) GetDBTableInfos(c *gin.Context) {
	tableInfos, err := genTableService.GetDBTableInfos()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.SuccessWithData(tableInfos, c)
	}
}

// ImportTables
// @Summary 导入表
// @Tags 代码生成器
// @param tableNames body []string true "表名"
// @Success 200 {object} response.JsonResult
// @Security ApiKeyAuth
// @Router /sysTable/importTables [post]
func (sysTableApi *SysTableApi) ImportTables(c *gin.Context) {
	var tableNames []string
	c.ShouldBindJSON(&tableNames)
	fmt.Println(tableNames)
	// 生成表
	tables, err := genTableService.ImportGentTable(tableNames)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 生成表字段
	err = genTableColumnService.ImportGenTableColumn(tableNames, tables)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Success(c)
	}
}

// PreviewCode
// @Summary 预览代码
// @Tags 代码生成器
// @param tableId query []string true "表名"
// @Success 200 {object} response.JsonResult
// @Security ApiKeyAuth
// @Router /sysTable/previewCode [get]
func (sysTableApi *SysTableApi) PreviewCode(c *gin.Context) {
	err := genTableService.PreviewCode(1)
	response.Complete(err, c)
}
