package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAsset_manage_systemRouter(Router *gin.RouterGroup) {
	Asset_manage_systemRouter := Router.Group("asset_manage_system").Use(middleware.OperationRecord())
	{
		Asset_manage_systemRouter.POST("createAsset_manage_system", v1.CreateAsset_manage_system)   // 新建Asset_manage_system
		Asset_manage_systemRouter.DELETE("deleteAsset_manage_system", v1.DeleteAsset_manage_system) // 删除Asset_manage_system
		Asset_manage_systemRouter.DELETE("deleteAsset_manage_systemByIds", v1.DeleteAsset_manage_systemByIds) // 批量删除Asset_manage_system
		Asset_manage_systemRouter.PUT("updateAsset_manage_system", v1.UpdateAsset_manage_system)    // 更新Asset_manage_system
		Asset_manage_systemRouter.GET("findAsset_manage_system", v1.FindAsset_manage_system)        // 根据ID获取Asset_manage_system
		Asset_manage_systemRouter.GET("getAsset_manage_systemList", v1.GetAsset_manage_systemList)  // 获取Asset_manage_system列表

	}
}
func InitAsset_manage_systemRouter_export(Router *gin.RouterGroup) {
	Asset_manage_systemRouter := Router.Group("asset_manage_system")
	{
		Asset_manage_systemRouter.POST("exportAsset_manage_system_resultsByIds", v1.ExportAsset_manage_system_resultsByIds)  // 根据ID导出Asset_manage_system列表
		Asset_manage_systemRouter.POST("exportAsset_manage_system_resultsByConditions", v1.ExportAsset_manage_system_resultsByConditions) // 根据筛选条件导出Asset_manage_system列表
	}
}