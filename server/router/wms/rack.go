package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RackRouter struct {
}

// InitRackRouter 初始化 货架 路由信息
func (s *RackRouter) InitRackRouter(Router *gin.RouterGroup) {
	rackRouter := Router.Group("rack").Use(middleware.OperationRecord())
	rackRouterWithoutRecord := Router.Group("rack")
	var rackApi = v1.ApiGroupApp.WmsApiGroup.RackApi
	{
		rackRouter.POST("createRack", rackApi.CreateRack)   // 新建货架
		rackRouter.DELETE("deleteRack", rackApi.DeleteRack) // 删除货架
		rackRouter.DELETE("deleteRackByIds", rackApi.DeleteRackByIds) // 批量删除货架
		rackRouter.PUT("updateRack", rackApi.UpdateRack)    // 更新货架
	}
	{
		rackRouterWithoutRecord.GET("findRack", rackApi.FindRack)        // 根据ID获取货架
		rackRouterWithoutRecord.GET("getRackList", rackApi.GetRackList)  // 获取货架列表
	}
}
