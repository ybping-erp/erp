import service from '@/utils/request'

// @Tags OrderItem
// @Summary 创建订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "创建订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /orderItem/createOrderItem [post]
export const createOrderItem = (data) => {
  return service({
    url: '/orderItem/createOrderItem',
    method: 'post',
    data
  })
}

// @Tags OrderItem
// @Summary 删除订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "删除订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderItem/deleteOrderItem [delete]
export const deleteOrderItem = (data) => {
  return service({
    url: '/orderItem/deleteOrderItem',
    method: 'delete',
    data
  })
}

// @Tags OrderItem
// @Summary 批量删除订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderItem/deleteOrderItem [delete]
export const deleteOrderItemByIds = (data) => {
  return service({
    url: '/orderItem/deleteOrderItemByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderItem
// @Summary 更新订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "更新订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderItem/updateOrderItem [put]
export const updateOrderItem = (data) => {
  return service({
    url: '/orderItem/updateOrderItem',
    method: 'put',
    data
  })
}

// @Tags OrderItem
// @Summary 用id查询订单明细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderItem true "用id查询订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderItem/findOrderItem [get]
export const findOrderItem = (params) => {
  return service({
    url: '/orderItem/findOrderItem',
    method: 'get',
    params
  })
}

// @Tags OrderItem
// @Summary 分页获取订单明细列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单明细列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderItem/getOrderItemList [get]
export const getOrderItemList = (params) => {
  return service({
    url: '/orderItem/getOrderItemList',
    method: 'get',
    params
  })
}
