package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OutboundLogRouter struct {
}

// InitOutboundLogRouter 初始化 出库记录 路由信息
func (s *OutboundLogRouter) InitOutboundLogRouter(Router *gin.RouterGroup) {
	outboundLogRouter := Router.Group("outboundLog").Use(middleware.OperationRecord())
	outboundLogRouterWithoutRecord := Router.Group("outboundLog")
	var outboundLogApi = v1.ApiGroupApp.WmsApiGroup.OutboundLogApi
	{
		outboundLogRouter.POST("createOutboundLog", outboundLogApi.CreateOutboundLog)   // 新建出库记录
		outboundLogRouter.DELETE("deleteOutboundLog", outboundLogApi.DeleteOutboundLog) // 删除出库记录
		outboundLogRouter.DELETE("deleteOutboundLogByIds", outboundLogApi.DeleteOutboundLogByIds) // 批量删除出库记录
		outboundLogRouter.PUT("updateOutboundLog", outboundLogApi.UpdateOutboundLog)    // 更新出库记录
	}
	{
		outboundLogRouterWithoutRecord.GET("findOutboundLog", outboundLogApi.FindOutboundLog)        // 根据ID获取出库记录
		outboundLogRouterWithoutRecord.GET("getOutboundLogList", outboundLogApi.GetOutboundLogList)  // 获取出库记录列表
	}
}
