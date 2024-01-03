package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderItemRouter struct {
}

// InitOrderItemRouter 初始化 订单明细 路由信息
func (s *OrderItemRouter) InitOrderItemRouter(Router *gin.RouterGroup) {
	orderItemRouter := Router.Group("orderItem").Use(middleware.OperationRecord())
	orderItemRouterWithoutRecord := Router.Group("orderItem")
	var orderItemApi = v1.ApiGroupApp.ECommerceApiGroup.OrderItemApi
	{
		orderItemRouter.POST("createOrderItem", orderItemApi.CreateOrderItem)   // 新建订单明细
		orderItemRouter.DELETE("deleteOrderItem", orderItemApi.DeleteOrderItem) // 删除订单明细
		orderItemRouter.DELETE("deleteOrderItemByIds", orderItemApi.DeleteOrderItemByIds) // 批量删除订单明细
		orderItemRouter.PUT("updateOrderItem", orderItemApi.UpdateOrderItem)    // 更新订单明细
	}
	{
		orderItemRouterWithoutRecord.GET("findOrderItem", orderItemApi.FindOrderItem)        // 根据ID获取订单明细
		orderItemRouterWithoutRecord.GET("getOrderItemList", orderItemApi.GetOrderItemList)  // 获取订单明细列表
	}
}
