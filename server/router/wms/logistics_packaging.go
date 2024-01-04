package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LogisticsPackagingRouter struct {
}

// InitLogisticsPackagingRouter 初始化 物流包装 路由信息
func (s *LogisticsPackagingRouter) InitLogisticsPackagingRouter(Router *gin.RouterGroup) {
	logisticsPackagingRouter := Router.Group("logisticsPackaging").Use(middleware.OperationRecord())
	logisticsPackagingRouterWithoutRecord := Router.Group("logisticsPackaging")
	var logisticsPackagingApi = v1.ApiGroupApp.WmsApiGroup.LogisticsPackagingApi
	{
		logisticsPackagingRouter.POST("createLogisticsPackaging", logisticsPackagingApi.CreateLogisticsPackaging)   // 新建物流包装
		logisticsPackagingRouter.DELETE("deleteLogisticsPackaging", logisticsPackagingApi.DeleteLogisticsPackaging) // 删除物流包装
		logisticsPackagingRouter.DELETE("deleteLogisticsPackagingByIds", logisticsPackagingApi.DeleteLogisticsPackagingByIds) // 批量删除物流包装
		logisticsPackagingRouter.PUT("updateLogisticsPackaging", logisticsPackagingApi.UpdateLogisticsPackaging)    // 更新物流包装
	}
	{
		logisticsPackagingRouterWithoutRecord.GET("findLogisticsPackaging", logisticsPackagingApi.FindLogisticsPackaging)        // 根据ID获取物流包装
		logisticsPackagingRouterWithoutRecord.GET("getLogisticsPackagingList", logisticsPackagingApi.GetLogisticsPackagingList)  // 获取物流包装列表
	}
}
