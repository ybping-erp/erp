import service from '@/utils/request'

// @Tags SkuMapping
// @Summary 创建商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SkuMapping true "创建商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /skuMapping/createSkuMapping [post]
export const createSkuMapping = (data) => {
  return service({
    url: '/skuMapping/createSkuMapping',
    method: 'post',
    data
  })
}

// @Tags SkuMapping
// @Summary 删除商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SkuMapping true "删除商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /skuMapping/deleteSkuMapping [delete]
export const deleteSkuMapping = (data) => {
  return service({
    url: '/skuMapping/deleteSkuMapping',
    method: 'delete',
    data
  })
}

// @Tags SkuMapping
// @Summary 批量删除商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /skuMapping/deleteSkuMapping [delete]
export const deleteSkuMappingByIds = (data) => {
  return service({
    url: '/skuMapping/deleteSkuMappingByIds',
    method: 'delete',
    data
  })
}

// @Tags SkuMapping
// @Summary 更新商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SkuMapping true "更新商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /skuMapping/updateSkuMapping [put]
export const updateSkuMapping = (data) => {
  return service({
    url: '/skuMapping/updateSkuMapping',
    method: 'put',
    data
  })
}

// @Tags SkuMapping
// @Summary 用id查询商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SkuMapping true "用id查询商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /skuMapping/findSkuMapping [get]
export const findSkuMapping = (params) => {
  return service({
    url: '/skuMapping/findSkuMapping',
    method: 'get',
    params
  })
}

// @Tags SkuMapping
// @Summary 分页获取商品映射列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品映射列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /skuMapping/getSkuMappingList [get]
export const getSkuMappingList = (params) => {
  return service({
    url: '/skuMapping/getSkuMappingList',
    method: 'get',
    params
  })
}
