package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
	"strconv"
)

type SysGenTableApi struct{}

// Query
// @Tags 代码生成（表）
// @Summary 代码生成（表）查询
// @Security ApiKeyAuth
// @Param data query dto.SysGenTableQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysGenTableVo}}
// @Router  /system/genTable/query [get]
func (s *SysGenTableApi) Query(c *gin.Context) {
	var query dto.SysGenTableQuery
	c.ShouldBindQuery(&query)
	list, total, err := genTableService.Query(query)
	response.Complete(response.PageResult{List: list, TotalCount: total}, err, c)
}

// Add
// @Tags 代码生成（表）
// @Summary 添加代码生成（表）
// @Security ApiKeyAuth
// @Param data body dto.SysGenTableAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router  /system/genTable/add [post]
func (s *SysGenTableApi) Add(c *gin.Context) {
	var addDto dto.SysGenTableAddDto
	c.ShouldBindJSON(&addDto)
	err := genTableService.Add(addDto)
	response.CompleteWithMessage(err, c)
}

// Update
// @Tags 代码生成（表）
// @Summary 更新代码生成（表）
// @Security ApiKeyAuth
// @Param data body dto.SysGenTableUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router  /system/genTable/update [post]
func (s *SysGenTableApi) Update(c *gin.Context) {
	var updateDto dto.SysGenTableUpdateDto
	c.ShouldBindJSON(&updateDto)
	err := genTableService.Update(updateDto)
	response.CompleteWithMessage(err, c)
}

// Delete
// @Tags 代码生成（表）
// @Summary 删除代码生成（表）
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router  /system/genTable/delete [post]
func (s *SysGenTableApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)
	err := genTableService.Delete(idInfoDto.Id)
	response.CompleteWithMessage(err, c)
}

// Detail
// @Tags 代码生成（表）
// @Summary 获取代码生成（表）详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "代码生成（表）id"
// @Success 200 {object} response.JsonResult{data=dto.SysGenTableVo}
// @Router  /system/genTable/detail [get]
func (s *SysGenTableApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := genTableService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}

// GetTableList
// @Summary 获取表列表
// @Tags 代码生成（表）
// @Success 200 {object} response.JsonResult{data=[]dto.TableInfoVo}
// @Security ApiKeyAuth
// @Router /system/genTable/tableList [get]
func (s *SysGenTableApi) GetTableList(c *gin.Context) {
	tableList, err := genTableService.GetTableList()
	response.Complete(tableList, err, c)
}

// PreviewCode
// @Summary 预览代码
// @Tags 代码生成（表）
// @Param tableId query string true "表id"
// @Success 200 {object} response.JsonResult
// @Security ApiKeyAuth
// @Router /system/genTable/previewCode [get]
func (s *SysGenTableApi) PreviewCode(c *gin.Context) {
	tableId, err := strconv.ParseInt(c.Query("tableId"), 10, 64)
	vos, err := genTableService.PreviewCode(tableId)
	response.Complete(vos, err, c)
}
