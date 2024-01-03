package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type OrderItemApi struct {
}

var orderItemService = service.ServiceGroupApp.ECommerceServiceGroup.OrderItemService


// CreateOrderItem 创建订单明细
// @Tags OrderItem
// @Summary 创建订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.OrderItem true "创建订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /orderItem/createOrderItem [post]
func (orderItemApi *OrderItemApi) CreateOrderItem(c *gin.Context) {
	var orderItem eCommerce.OrderItem
	err := c.ShouldBindJSON(&orderItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderItemService.CreateOrderItem(&orderItem); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrderItem 删除订单明细
// @Tags OrderItem
// @Summary 删除订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.OrderItem true "删除订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderItem/deleteOrderItem [delete]
func (orderItemApi *OrderItemApi) DeleteOrderItem(c *gin.Context) {
	var orderItem eCommerce.OrderItem
	err := c.ShouldBindJSON(&orderItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderItemService.DeleteOrderItem(orderItem); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderItemByIds 批量删除订单明细
// @Tags OrderItem
// @Summary 批量删除订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderItem/deleteOrderItemByIds [delete]
func (orderItemApi *OrderItemApi) DeleteOrderItemByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderItemService.DeleteOrderItemByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderItem 更新订单明细
// @Tags OrderItem
// @Summary 更新订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.OrderItem true "更新订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderItem/updateOrderItem [put]
func (orderItemApi *OrderItemApi) UpdateOrderItem(c *gin.Context) {
	var orderItem eCommerce.OrderItem
	err := c.ShouldBindJSON(&orderItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderItemService.UpdateOrderItem(orderItem); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderItem 用id查询订单明细
// @Tags OrderItem
// @Summary 用id查询订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eCommerce.OrderItem true "用id查询订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderItem/findOrderItem [get]
func (orderItemApi *OrderItemApi) FindOrderItem(c *gin.Context) {
	var orderItem eCommerce.OrderItem
	err := c.ShouldBindQuery(&orderItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorderItem, err := orderItemService.GetOrderItem(orderItem.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorderItem": reorderItem}, c)
	}
}

// GetOrderItemList 分页获取订单明细列表
// @Tags OrderItem
// @Summary 分页获取订单明细列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eCommerceReq.OrderItemSearch true "分页获取订单明细列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderItem/getOrderItemList [get]
func (orderItemApi *OrderItemApi) GetOrderItemList(c *gin.Context) {
	var pageInfo eCommerceReq.OrderItemSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderItemService.GetOrderItemInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
