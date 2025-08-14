package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
	"strconv"
)

type SysDictItemApi struct{}

// Query
// @Tags 字典选项
// @Summary 字典选项查询
// @Security ApiKeyAuth
// @Param data query dto.SysDictItemQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysDictItemVo}}
// @Router /system/dictItem/query [get]
func (sysUserApi *SysDictItemApi) Query(c *gin.Context) {
	var query dto.SysDictItemQuery
	c.ShouldBindQuery(&query)
	list, total, err := dictItemService.Query(query)
	response.Complete(response.PageResult{List: list, TotalCount: total}, err, c)
}

// Add
// @Tags 字典选项
// @Summary 添加字典选项
// @Security ApiKeyAuth
// @Param data body dto.SysDictItemAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/dictItem/add [post]
func (sysUserApi *SysDictItemApi) Add(c *gin.Context) {
	var addDto dto.SysDictItemAddDto
	c.ShouldBindJSON(&addDto)
	err := dictItemService.Add(addDto)
	response.CompleteWithMessage(err, c)
}

// Update
// @Tags 字典选项
// @Summary 更新字典选项
// @Security ApiKeyAuth
// @Param data body dto.SysDictItemUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/dictItem/update [post]
func (sysUserApi *SysDictItemApi) Update(c *gin.Context) {
	var updateDto dto.SysDictItemUpdateDto
	c.ShouldBindJSON(&updateDto)
	err := dictItemService.Update(updateDto)
	response.CompleteWithMessage(err, c)
}

// Detail
// @Tags 字典选项
// @Summary 获取字典选项详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "字典选项id"
// @Success 200 {object} response.JsonResult{data=dto.SysDictItemVo}
// @Router /system/dictItem/detail [get]
func (sysUserApi *SysDictItemApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := dictItemService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}

// Delete
// @Tags 字典选项
// @Summary 删除字典选项
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/dictItem/delete [post]
func (sysUserApi *SysDictItemApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)
	err := dictItemService.Delete(idInfoDto.Id)
	response.CompleteWithMessage(err, c)
}

// List
// @Tags 字典选项
// @Summary 字典选项列表
// @Security ApiKeyAuth
// @Param dictId query int64 true "字典id"
// @Success 200 {object} response.JsonResult{data=[]dto.SysDictItemVo}
// @Router /system/dictItem/list [get]
func (sysUserApi *SysDictItemApi) List(c *gin.Context) {

	dictId, err := strconv.Atoi(c.Query("dictId"))
	if err != nil {
		response.Complete(nil, err, c)
		return
	}
	objs, err := dictItemService.List(dictId)
	response.Complete(objs, err, c)
}
