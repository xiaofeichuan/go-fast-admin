package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
)

type SysMenuApi struct{}

// Query
// @Tags 菜单
// @Summary 菜单查询
// @Security ApiKeyAuth
// @Param data query dto.SysMenuQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysMenuVo}}
// @Router /system/menu/query [get]
func (*SysMenuApi) Query(c *gin.Context) {
	var query dto.SysMenuQuery
	c.ShouldBindQuery(&query)

	list, total, err := menuService.Query(query)
	response.Complete(response.PageResult{List: list, TotalCount: total}, err, c)
}

// Add
// @Tags 菜单
// @Summary 添加菜单
// @Security ApiKeyAuth
// @Param data body dto.SysMenuAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/menu/add [post]
func (*SysMenuApi) Add(c *gin.Context) {
	var addDto dto.SysMenuAddDto
	c.ShouldBindJSON(&addDto)

	err := menuService.Add(addDto)
	response.CompleteWithMessage(err, c)
}

// Update
// @Tags 菜单
// @Summary 更新菜单
// @Security ApiKeyAuth
// @Param data body dto.SysMenuUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/menu/update [post]
func (*SysMenuApi) Update(c *gin.Context) {
	var updateDto dto.SysMenuUpdateDto
	c.ShouldBindJSON(&updateDto)

	err := menuService.Update(updateDto)
	response.CompleteWithMessage(err, c)
}

// Detail
// @Tags 菜单
// @Summary 获取菜单详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "菜单id"
// @Success 200 {object} response.JsonResult{data=dto.SysMenuVo}
// @Router /system/menu/detail [get]
func (*SysMenuApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)

	obj, err := menuService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}

// Delete
// @Tags 菜单
// @Summary 删除菜单
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/menu/delete [post]
func (*SysMenuApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)

	err := menuService.Delete(idInfoDto.Id)
	response.CompleteWithMessage(err, c)
}

// GetMenuTree
// @Tags 菜单
// @Summary 菜单树状
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.SysMenuVo}
// @Router /system/menu/tree [get]
func (*SysMenuApi) GetMenuTree(c *gin.Context) {
	objs, err := menuService.GetMenuTree()
	response.Complete(objs, err, c)
}
