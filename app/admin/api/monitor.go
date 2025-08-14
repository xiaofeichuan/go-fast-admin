package api

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/common/dto/response"
)

type SysMonitorApi struct{}

// GetCacheKeys
// @Tags 监控
// @Summary 获取缓存key
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]string}
// @Router /system/monitor/cacheKeys [get]
func (monitorApi *SysMonitorApi) GetCacheKeys(c *gin.Context) {
	objs := monitorService.GetCacheKeys()
	response.Complete(objs, nil, c)
}

// GetCacheValue
// @Tags 监控
// @Summary 获取缓存value
// @Security ApiKeyAuth
// @Param key query string true "缓存Key"
// @Success 200 {object} response.JsonResult{data=string}
// @Router /system/monitor/cacheValue [get]
func (monitorApi *SysMonitorApi) GetCacheValue(c *gin.Context) {
	key := c.Query("key")
	value := monitorService.GetCacheValue(key)
	response.Complete(value, nil, c)
}

// DeleteCache
// @Tags 监控
// @Summary 删除缓存
// @Security ApiKeyAuth
// @Param key query string true "缓存Key"
// @Success 200 {object} response.JsonResult{}
// @Router /system/monitor/deleteCache [post]
func (monitorApi *SysMonitorApi) DeleteCache(c *gin.Context) {
	key := c.Query("key")
	monitorService.DeleteCache(key)
	response.Complete(nil, nil, c)
}

// GetServerInfo
// @Tags 监控
// @Summary 获取服务器信息（建议：采用更专业的服务器监控软件）
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=string}
// @Router /system/monitor/serverInfo [get]
func (monitorApi *SysMonitorApi) GetServerInfo(c *gin.Context) {
	var server = monitorService.GetServerInfo()
	response.Complete(server, nil, c)
}
