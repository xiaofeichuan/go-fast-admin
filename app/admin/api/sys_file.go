package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
)

type SysFileApi struct{}

// Query
// @Tags 文件管理
// @Summary 文件管理查询
// @Security ApiKeyAuth
// @Param data query dto.SysFileQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysFileVo}}
// @Router /system/file/query [get]
func (sysUserApi *SysFileApi) Query(c *gin.Context) {
	var query dto.SysFileQuery
	c.ShouldBindQuery(&query)
	list, total, err := fileService.Query(query)
	response.Complete(response.PageResult{List: list, TotalCount: total}, err, c)
}

// Add
// @Tags 文件管理
// @Summary 添加文件管理
// @Security ApiKeyAuth
// @Param data body dto.SysFileAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/file/add [post]
func (sysUserApi *SysFileApi) Add(c *gin.Context) {
	var addDto dto.SysFileAddDto
	c.ShouldBindJSON(&addDto)
	err := fileService.Add(addDto)
	response.CompleteWithMessage(err, c)
}

// Detail
// @Tags 文件管理
// @Summary 获取文件管理详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "文件管理id"
// @Success 200 {object} response.JsonResult{data=dto.SysFileVo}
// @Router /system/file/detail [get]
func (sysUserApi *SysFileApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := fileService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}
