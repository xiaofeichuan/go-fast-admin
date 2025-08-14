package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/request"
	"go-fast-admin/server/common/dto/response"
)

type SysLoginLogApi struct{}

// Query
// @Tags 登录日志
// @Summary 登录日志查询
// @Security ApiKeyAuth
// @Param data query dto.SysLoginLogQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysLoginLogVo}}
// @Router /system/loginLog/query [get]
func (sysUserApi *SysLoginLogApi) Query(c *gin.Context) {
	var query dto.SysLoginLogQuery
	c.ShouldBindQuery(&query)
	list, total, err := loginLogService.Query(query)
	response.Complete(response.PageResult{List: list, TotalCount: total}, err, c)
}

// Detail
// @Tags 登录日志
// @Summary 获取登录日志详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "登录日志id"
// @Success 200 {object} response.JsonResult{data=dto.SysLoginLogVo}
// @Router /system/loginLog/detail [get]
func (sysUserApi *SysLoginLogApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := loginLogService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}
