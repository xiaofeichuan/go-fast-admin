package api

import (
	"gin-fast-admin/server/app/admin/dto"
	"gin-fast-admin/server/common/dto/request"
	"gin-fast-admin/server/common/dto/response"
	"github.com/gin-gonic/gin"
)

type SysUserApi struct{}

// Page
// @Tags 用户
// @Summary 用户分页查询
// @Security ApiKeyAuth
// @Param data query dto.SysUserQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysUserVo}}
// @Router /sysUser/page [get]
func (sysUserApi *SysUserApi) Page(c *gin.Context) {
	var query dto.SysUserQuery
	c.ShouldBindQuery(&query)
	list, total, err := userService.Page(query)
	response.Complete(response.PageResult{
		List:       list,
		TotalCount: total,
	}, err, c)
}

// Detail
// @Tags 用户
// @Summary 获取用户详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "用户id"
// @Success 200 {object} response.JsonResult{data=dto.SysUserVo}
// @Router /sysUser/detail [get]
func (sysUserApi *SysUserApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := userService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}

// List
// @Tags 用户
// @Summary 用户列表
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.SysUserVo}
// @Router /sysUser/list [get]
func (sysUserApi *SysUserApi) List(c *gin.Context) {
	objs, err := userService.List()
	response.Complete(objs, err, c)
}

// Add
// @Tags 用户
// @Summary 添加用户
// @Security ApiKeyAuth
// @Param data body dto.SysUserCreateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /sysUser/add [post]
func (sysUserApi *SysUserApi) Add(c *gin.Context) {
	var createDto dto.SysUserCreateDto
	c.ShouldBindJSON(&createDto)
	err := userService.Add(createDto)
	response.CompleteWithMessage(err, c)
}

// Update
// @Tags 用户
// @Summary 更新用户
// @Security ApiKeyAuth
// @Param data body dto.SysUserUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /sysUser/update [post]
func (sysUserApi *SysUserApi) Update(c *gin.Context) {
	var updateDto dto.SysUserUpdateDto
	c.ShouldBindJSON(&updateDto)
	err := userService.Update(updateDto)
	response.CompleteWithMessage(err, c)
}

// Delete
// @Tags 用户
// @Summary 删除用户
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /sysUser/delete [post]
func (sysUserApi *SysUserApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)
	err := userService.Delete(idInfoDto.Id)
	response.CompleteWithMessage(err, c)
}

//// UserLogin
//// @Summary
//// @Tags 用户
//// @Success 200 {object} response.JsonResult
//// @Security ApiKeyAuth
//// @Param userId query int true "用户id"
//// @Router /user/login [post]
//func (sysUserApi *SysUserApi) UserLogin(c *gin.Context) {
//	response.Success(c)
//}
