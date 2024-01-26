import service from '@/utils/request'

// @Tags OutboundLog
// @Summary 创建出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutboundLog true "创建出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /outboundLog/createOutboundLog [post]
export const createOutboundLog = (data) => {
  return service({
    url: '/outboundLog/createOutboundLog',
    method: 'post',
    data
  })
}

// @Tags OutboundLog
// @Summary 删除出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutboundLog true "删除出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outboundLog/deleteOutboundLog [delete]
export const deleteOutboundLog = (params) => {
  return service({
    url: '/outboundLog/deleteOutboundLog',
    method: 'delete',
    params
  })
}

// @Tags OutboundLog
// @Summary 批量删除出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outboundLog/deleteOutboundLog [delete]
export const deleteOutboundLogByIds = (params) => {
  return service({
    url: '/outboundLog/deleteOutboundLogByIds',
    method: 'delete',
    params
  })
}

// @Tags OutboundLog
// @Summary 更新出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutboundLog true "更新出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /outboundLog/updateOutboundLog [put]
export const updateOutboundLog = (data) => {
  return service({
    url: '/outboundLog/updateOutboundLog',
    method: 'put',
    data
  })
}

// @Tags OutboundLog
// @Summary 用id查询出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OutboundLog true "用id查询出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /outboundLog/findOutboundLog [get]
export const findOutboundLog = (params) => {
  return service({
    url: '/outboundLog/findOutboundLog',
    method: 'get',
    params
  })
}

// @Tags OutboundLog
// @Summary 分页获取出库记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取出库记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outboundLog/getOutboundLogList [get]
export const getOutboundLogList = (params) => {
  return service({
    url: '/outboundLog/getOutboundLogList',
    method: 'get',
    params
  })
}
