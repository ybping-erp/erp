package wms

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InventoryRouter struct {
}

// InitInventoryRouter 初始化 商品库存 路由信息
func (s *InventoryRouter) InitInventoryRouter(Router *gin.RouterGroup) {
	inventoryRouter := Router.Group("inventory").Use(middleware.OperationRecord())
	inventoryRouterWithoutRecord := Router.Group("inventory")
	var inventoryApi = v1.ApiGroupApp.WmsApiGroup.InventoryApi
	{
		inventoryRouter.PUT("addInventoryFromInboundLog", inventoryApi.AddInventoryFromInboundLog) // 入库商品库存
		inventoryRouter.DELETE("deleteInventory", inventoryApi.DeleteInventory)                    // 删除商品库存
		inventoryRouter.DELETE("deleteInventoryByIds", inventoryApi.DeleteInventoryByIds)          // 批量删除商品库存
		inventoryRouter.PUT("updateInventory", inventoryApi.UpdateInventory)                       // 更新商品库存
	}
	{
		inventoryRouterWithoutRecord.GET("findInventory", inventoryApi.FindInventory)       // 根据ID获取商品库存
		inventoryRouterWithoutRecord.GET("getInventoryList", inventoryApi.GetInventoryList) // 获取商品库存列表
	}
}
