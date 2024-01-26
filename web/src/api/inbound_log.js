import service from '@/utils/request'

// @Tags InboundLog
// @Summary 创建入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InboundLog true "创建入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /inboundLog/createInboundLog [post]
export const createInboundLog = (data) => {
  return service({
    url: '/inboundLog/createInboundLog',
    method: 'post',
    data
  })
}

// @Tags InboundLog
// @Summary 删除入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InboundLog true "删除入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inboundLog/deleteInboundLog [delete]
export const deleteInboundLog = (params) => {
  return service({
    url: '/inboundLog/deleteInboundLog',
    method: 'delete',
    params
  })
}

// @Tags InboundLog
// @Summary 批量删除入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inboundLog/deleteInboundLog [delete]
export const deleteInboundLogByIds = (params) => {
  return service({
    url: '/inboundLog/deleteInboundLogByIds',
    method: 'delete',
    params
  })
}

// @Tags InboundLog
// @Summary 更新入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InboundLog true "更新入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inboundLog/updateInboundLog [put]
export const updateInboundLog = (data) => {
  return service({
    url: '/inboundLog/updateInboundLog',
    method: 'put',
    data
  })
}

// @Tags InboundLog
// @Summary 用id查询入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InboundLog true "用id查询入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inboundLog/findInboundLog [get]
export const findInboundLog = (params) => {
  return service({
    url: '/inboundLog/findInboundLog',
    method: 'get',
    params
  })
}

// @Tags InboundLog
// @Summary 分页获取入库记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取入库记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inboundLog/getInboundLogList [get]
export const getInboundLogList = (params) => {
  return service({
    url: '/inboundLog/getInboundLogList',
    method: 'get',
    params
  })
}
