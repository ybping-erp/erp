package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InboundLogRouter struct {
}

// InitInboundLogRouter 初始化 入库记录 路由信息
func (s *InboundLogRouter) InitInboundLogRouter(Router *gin.RouterGroup) {
	inboundLogRouter := Router.Group("inboundLog").Use(middleware.OperationRecord())
	inboundLogRouterWithoutRecord := Router.Group("inboundLog")
	var inboundLogApi = v1.ApiGroupApp.WmsApiGroup.InboundLogApi
	{
		inboundLogRouter.POST("createInboundLog", inboundLogApi.CreateInboundLog)   // 新建入库记录
		inboundLogRouter.DELETE("deleteInboundLog", inboundLogApi.DeleteInboundLog) // 删除入库记录
		inboundLogRouter.DELETE("deleteInboundLogByIds", inboundLogApi.DeleteInboundLogByIds) // 批量删除入库记录
		inboundLogRouter.PUT("updateInboundLog", inboundLogApi.UpdateInboundLog)    // 更新入库记录
	}
	{
		inboundLogRouterWithoutRecord.GET("findInboundLog", inboundLogApi.FindInboundLog)        // 根据ID获取入库记录
		inboundLogRouterWithoutRecord.GET("getInboundLogList", inboundLogApi.GetInboundLogList)  // 获取入库记录列表
	}
}
