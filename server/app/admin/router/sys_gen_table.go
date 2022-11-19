package router

import (
	"go-fast-admin/server/app/admin/api"

	"github.com/gin-gonic/gin"
)

type SysTableRouter struct{}

// 初始化权限路由
func (s *SysTableRouter) InitSysGenTablePrivateRouter(routerGroup *gin.RouterGroup) {
	var sysTableApi = api.SysTableApi{}
	sysTableRouter := routerGroup.Group("sysTable")
	{
		sysTableRouter.GET("getDBTableInfos", sysTableApi.GetDBTableInfos) // 获取当前数据库所有表信息
		sysTableRouter.POST("importTables", sysTableApi.ImportTables)      // 导入表
		sysTableRouter.GET("previewCode", sysTableApi.PreviewCode)         // 导入表
	}
}
