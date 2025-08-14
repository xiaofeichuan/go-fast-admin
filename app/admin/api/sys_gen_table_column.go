package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/dto/response"
)

type SysGenTableColumnApi struct{}

// Update
// @Tags 代码生成（表字段）
// @Summary 更新代码生成（表字段）
// @Security ApiKeyAuth
// @Param data body []dto.SysGenTableUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router  /system/genTableColumn/update [post]
func (s *SysGenTableColumnApi) Update(c *gin.Context) {
	var updateDto []dto.SysGenTableColumnUpdateDto
	c.ShouldBindJSON(&updateDto)
	genTableColumnService.Update(updateDto)
	response.CompleteWithMessage(nil, c)
}

// List
// @Tags 代码生成（表字段）
// @Summary 代码生成（表字段）列表
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.SysGenTableColumnVo}
// @Router  /system/genTableColumn/list [get]
func (s *SysGenTableColumnApi) List(c *gin.Context) {
	var query dto.SysGenTableColumnQuery
	c.ShouldBindQuery(&query)
	objs, err := genTableColumnService.List(query.TableId)
	response.Complete(objs, err, c)
}
