import service from '@/utils/request'

// @Tags Inventory
// @Summary 创建商品库存
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Inventory true "创建商品库存"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /inventory/createInventory [post]
export const createInventory = (data) => {
  return service({
    url: '/inventory/createInventory',
    method: 'post',
    data
  })
}

// @Tags Inventory
// @Summary 删除商品库存
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Inventory true "删除商品库存"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inventory/deleteInventory [delete]
export const deleteInventory = (data) => {
  return service({
    url: '/inventory/deleteInventory',
    method: 'delete',
    data
  })
}

// @Tags Inventory
// @Summary 批量删除商品库存
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品库存"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inventory/deleteInventory [delete]
export const deleteInventoryByIds = (data) => {
  return service({
    url: '/inventory/deleteInventoryByIds',
    method: 'delete',
    data
  })
}

// @Tags Inventory
// @Summary 更新商品库存
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Inventory true "更新商品库存"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inventory/updateInventory [put]
export const updateInventory = (data) => {
  return service({
    url: '/inventory/updateInventory',
    method: 'put',
    data
  })
}

// @Tags Inventory
// @Summary 用id查询商品库存
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Inventory true "用id查询商品库存"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inventory/findInventory [get]
export const findInventory = (params) => {
  return service({
    url: '/inventory/findInventory',
    method: 'get',
    params
  })
}

// @Tags Inventory
// @Summary 分页获取商品库存列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品库存列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inventory/getInventoryList [get]
export const getInventoryList = (params) => {
  return service({
    url: '/inventory/getInventoryList',
    method: 'get',
    params
  })
}


// AddInventoryFromInboundLog 根据入库单入库
// @Tags Inventory
// @Summary 创建商品库存
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Inventory true "创建商品库存"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"入库成功"}"
// @Router /inventory/addInventoryFromInboundLog [put]
export const addInventoryFromInboundLog = (inboundLogID) => {
  return service({
    url: `/inventory/addInventoryFromInboundLog?inboundLogID=${inboundLogID}`,
    method: 'put',
  })
}

