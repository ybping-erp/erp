package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PickOrderRouter struct {
}

// InitPickOrderRouter 初始化 拣货单 路由信息
func (s *PickOrderRouter) InitPickOrderRouter(Router *gin.RouterGroup) {
	pickOrderRouter := Router.Group("pickOrder").Use(middleware.OperationRecord())
	pickOrderRouterWithoutRecord := Router.Group("pickOrder")
	var pickOrderApi = v1.ApiGroupApp.WmsApiGroup.PickOrderApi
	{
		pickOrderRouter.POST("createPickOrder", pickOrderApi.CreatePickOrder)   // 新建拣货单
		pickOrderRouter.DELETE("deletePickOrder", pickOrderApi.DeletePickOrder) // 删除拣货单
		pickOrderRouter.DELETE("deletePickOrderByIds", pickOrderApi.DeletePickOrderByIds) // 批量删除拣货单
		pickOrderRouter.PUT("updatePickOrder", pickOrderApi.UpdatePickOrder)    // 更新拣货单
	}
	{
		pickOrderRouterWithoutRecord.GET("findPickOrder", pickOrderApi.FindPickOrder)        // 根据ID获取拣货单
		pickOrderRouterWithoutRecord.GET("getPickOrderList", pickOrderApi.GetPickOrderList)  // 获取拣货单列表
	}
}
