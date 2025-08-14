package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
)

type SysDictApi struct{}

// Query
// @Tags 字典
// @Summary 字典查询
// @Security ApiKeyAuth
// @Param data query dto.SysDictQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysDictVo}}
// @Router /system/dict/query [get]
func (sysUserApi *SysDictApi) Query(c *gin.Context) {
	var query dto.SysDictQuery
	c.ShouldBindQuery(&query)
	list, total, err := dictService.Query(query)
	response.Complete(response.PageResult{List: list, TotalCount: total}, err, c)
}

// Add
// @Tags 字典
// @Summary 添加字典
// @Security ApiKeyAuth
// @Param data body dto.SysDictAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/dict/add [post]
func (sysUserApi *SysDictApi) Add(c *gin.Context) {
	var addDto dto.SysDictAddDto
	c.ShouldBindJSON(&addDto)
	err := dictService.Add(addDto)
	response.CompleteWithMessage(err, c)
}

// Update
// @Tags 字典
// @Summary 更新字典
// @Security ApiKeyAuth
// @Param data body dto.SysDictUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/dict/update [post]
func (sysUserApi *SysDictApi) Update(c *gin.Context) {
	var updateDto dto.SysDictUpdateDto
	c.ShouldBindJSON(&updateDto)
	err := dictService.Update(updateDto)
	response.CompleteWithMessage(err, c)
}

// Detail
// @Tags 字典
// @Summary 获取字典详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "字典id"
// @Success 200 {object} response.JsonResult{data=dto.SysDictVo}
// @Router /system/dict/detail [get]
func (sysUserApi *SysDictApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := dictService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}

// Delete
// @Tags 字典
// @Summary 删除字典
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/dict/delete [post]
func (sysUserApi *SysDictApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)
	err := dictService.Delete(idInfoDto.Id)
	response.CompleteWithMessage(err, c)
}

// List
// @Tags 字典
// @Summary 字典列表
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.SysDictVo}
// @Router /system/dict/list [get]
func (sysUserApi *SysDictApi) List(c *gin.Context) {
	objs, err := dictService.List()
	response.Complete(objs, err, c)
}
