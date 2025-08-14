package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
)

type SysRoleApi struct{}

// Query
// @Tags 角色
// @Summary 角色查询
// @Security ApiKeyAuth
// @Param data query dto.SysRoleQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysRoleVo}}
// @Router /system/role/query [get]
func (*SysRoleApi) Query(c *gin.Context) {
	var query dto.SysRoleQuery
	c.ShouldBindQuery(&query)
	list, total, err := roleService.Query(query)
	response.Complete(response.PageResult{
		List:       list,
		TotalCount: total,
	}, err, c)
}

// Add
// @Tags 角色
// @Summary 添加角色
// @Security ApiKeyAuth
// @Param data body dto.SysRoleAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/role/add [post]
func (*SysRoleApi) Add(c *gin.Context) {
	var addDto dto.SysRoleAddDto
	c.ShouldBindJSON(&addDto)
	err := roleService.Add(addDto)
	response.CompleteWithMessage(err, c)
}

// Update
// @Tags 角色
// @Summary 更新角色
// @Security ApiKeyAuth
// @Param data body dto.SysRoleUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/role/update [post]
func (*SysRoleApi) Update(c *gin.Context) {
	var updateDto dto.SysRoleUpdateDto
	c.ShouldBindJSON(&updateDto)
	err := roleService.Update(updateDto)
	response.CompleteWithMessage(err, c)
}

// Detail
// @Tags 角色
// @Summary 获取角色详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "用户id"
// @Success 200 {object} response.JsonResult{data=dto.SysRoleVo}
// @Router /system/role/detail [get]
func (*SysRoleApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := roleService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}

// Delete
// @Tags 角色
// @Summary 删除角色
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/role/delete [post]
func (*SysRoleApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)
	err := roleService.Delete(idInfoDto.Id)
	response.CompleteWithMessage(err, c)
}

// List
// @Tags 角色
// @Summary 角色列表
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.SysRoleVo}
// @Router /system/role/list [get]
func (*SysRoleApi) List(c *gin.Context) {
	objs, err := roleService.List()
	response.Complete(objs, err, c)
}
