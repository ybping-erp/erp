import service from '@/utils/request'

// @Tags Shop
// @Summary 创建店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shop true "创建店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shop/createShop [post]
export const createShop = (data) => {
  return service({
    url: '/shop/createShop',
    method: 'post',
    data
  })
}

// @Tags Shop
// @Summary 删除店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shop true "删除店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shop/deleteShop [delete]
export const deleteShop = (data) => {
  return service({
    url: '/shop/deleteShop',
    method: 'delete',
    data
  })
}

// @Tags Shop
// @Summary 批量删除店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shop/deleteShop [delete]
export const deleteShopByIds = (data) => {
  return service({
    url: '/shop/deleteShopByIds',
    method: 'delete',
    data
  })
}

// @Tags Shop
// @Summary 更新店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shop true "更新店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shop/updateShop [put]
export const updateShop = (data) => {
  return service({
    url: '/shop/updateShop',
    method: 'put',
    data
  })
}

// @Tags Shop
// @Summary 用id查询店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Shop true "用id查询店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shop/findShop [get]
export const findShop = (params) => {
  return service({
    url: '/shop/findShop',
    method: 'get',
    params
  })
}

// @Tags Shop
// @Summary 分页获取店铺列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取店铺列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shop/getShopList [get]
export const getShopList = (params) => {
  return service({
    url: '/shop/getShopList',
    method: 'get',
    params
  })
}


// @Tags Shop
// @Summary 店铺授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Shop true "店铺授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"授权成功"}"
// @Router /shop/authorize [get]
export const authorizeShop = (data) => {
  return service({
    url: '/shop/authorize',
    method: 'post',
    data
  })
}