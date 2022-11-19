package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
)

type SysUserApi struct{}

// Page
// @Tags 用户
// @Summary 用户分页查询
// @Security ApiKeyAuth
// @Param data query dto.SysUserQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysUserVo}}
// @Router /system/user/page [get]
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
// @Router /system/user/detail [get]
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
// @Router /system/user/list [get]
func (sysUserApi *SysUserApi) List(c *gin.Context) {
	objs, err := userService.List()
	response.Complete(objs, err, c)
}

// Add
// @Tags 用户
// @Summary 添加用户
// @Security ApiKeyAuth
// @Param data body dto.SysUserAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/user/add [post]
func (sysUserApi *SysUserApi) Add(c *gin.Context) {
	var addDto dto.SysUserAddDto
	c.ShouldBindJSON(&addDto)
	err := userService.Add(addDto)
	response.CompleteWithMessage(err, c)
}

// Update
// @Tags 用户
// @Summary 更新用户
// @Security ApiKeyAuth
// @Param data body dto.SysUserUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/user/update [post]
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
// @Router /system/user/delete [post]
func (sysUserApi *SysUserApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)
	err := userService.Delete(idInfoDto.Id)
	response.CompleteWithMessage(err, c)
}

//
//// ResetPwd
//// @Tags 用户
//// @Summary 重置密码
//// @Security ApiKeyAuth
//// @Param data body dto.ResetPwdDto true "请求参数"
//// @Success 200 {object} response.JsonResult
//// @Router /system/user/resetPwd [post]
//func (sysUserApi *SysUserApi) ResetPwd(c *gin.Context) {
//	var resetPwdDto dto.ResetPwdDto
//	c.ShouldBindJSON(&resetPwdDto)
//	err := userService.ResetPwd(resetPwdDto)
//	response.CompleteWithMessage(err, c)
//}

//
//// UpdatePwd
//// @Tags 用户
//// @Summary 更新密码
//// @Security ApiKeyAuth
//// @Param data body dto.UpdatePwdDto true "请求参数"
//// @Success 200 {object} response.JsonResult
//// @Router /system/user/updatePwd [post]
//func (sysUserApi *SysUserApi) UpdatePwd(c *gin.Context) {
//	var loginDto dto.LoginDto
//	c.ShouldBindJSON(&loginDto)
//	token, err := userService.Login(loginDto)
//	response.Complete(token, err, c)
//}
