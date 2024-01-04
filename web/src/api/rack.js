import service from '@/utils/request'

// @Tags Rack
// @Summary 创建货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rack true "创建货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rack/createRack [post]
export const createRack = (data) => {
  return service({
    url: '/rack/createRack',
    method: 'post',
    data
  })
}

// @Tags Rack
// @Summary 删除货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rack true "删除货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rack/deleteRack [delete]
export const deleteRack = (data) => {
  return service({
    url: '/rack/deleteRack',
    method: 'delete',
    data
  })
}

// @Tags Rack
// @Summary 批量删除货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rack/deleteRack [delete]
export const deleteRackByIds = (data) => {
  return service({
    url: '/rack/deleteRackByIds',
    method: 'delete',
    data
  })
}

// @Tags Rack
// @Summary 更新货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rack true "更新货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rack/updateRack [put]
export const updateRack = (data) => {
  return service({
    url: '/rack/updateRack',
    method: 'put',
    data
  })
}

// @Tags Rack
// @Summary 用id查询货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Rack true "用id查询货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rack/findRack [get]
export const findRack = (params) => {
  return service({
    url: '/rack/findRack',
    method: 'get',
    params
  })
}

// @Tags Rack
// @Summary 分页获取货架列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取货架列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rack/getRackList [get]
export const getRackList = (params) => {
  return service({
    url: '/rack/getRackList',
    method: 'get',
    params
  })
}
