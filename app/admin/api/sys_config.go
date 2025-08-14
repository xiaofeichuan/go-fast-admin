package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
)

type SysConfigApi struct{}

// Query
// @Tags 配置
// @Summary 配置查询
// @Security ApiKeyAuth
// @Param data query dto.SysConfigQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysConfigVo}}
// @Router /system/config/query [get]
func (sysUserApi *SysConfigApi) Query(c *gin.Context) {
	var query dto.SysConfigQuery
	c.ShouldBindQuery(&query)
	list, total, err := configService.Query(query)
	response.Complete(response.PageResult{List: list, TotalCount: total}, err, c)
}

// Add
// @Tags 配置
// @Summary 添加配置
// @Security ApiKeyAuth
// @Param data body dto.SysConfigAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/config/add [post]
func (sysUserApi *SysConfigApi) Add(c *gin.Context) {
	var addDto dto.SysConfigAddDto
	c.ShouldBindJSON(&addDto)
	err := configService.Add(addDto)
	response.CompleteWithMessage(err, c)
}

// Update
// @Tags 配置
// @Summary 更新配置
// @Security ApiKeyAuth
// @Param data body dto.SysConfigUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/config/update [post]
func (sysUserApi *SysConfigApi) Update(c *gin.Context) {
	var updateDto dto.SysConfigUpdateDto
	c.ShouldBindJSON(&updateDto)
	err := configService.Update(updateDto)
	response.CompleteWithMessage(err, c)
}

// Detail
// @Tags 配置
// @Summary 获取配置详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "配置id"
// @Success 200 {object} response.JsonResult{data=dto.SysConfigVo}
// @Router /system/config/detail [get]
func (sysUserApi *SysConfigApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := configService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}

// Delete
// @Tags 配置
// @Summary 删除配置
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/config/delete [post]
func (sysUserApi *SysConfigApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)
	err := configService.Delete(idInfoDto.Id)
	response.CompleteWithMessage(err, c)
}
