import service from '@/utils/request'

// @Tags LogisticsPackaging
// @Summary 创建物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LogisticsPackaging true "创建物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /logisticsPackaging/createLogisticsPackaging [post]
export const createLogisticsPackaging = (data) => {
  return service({
    url: '/logisticsPackaging/createLogisticsPackaging',
    method: 'post',
    data
  })
}

// @Tags LogisticsPackaging
// @Summary 删除物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LogisticsPackaging true "删除物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logisticsPackaging/deleteLogisticsPackaging [delete]
export const deleteLogisticsPackaging = (data) => {
  return service({
    url: '/logisticsPackaging/deleteLogisticsPackaging',
    method: 'delete',
    data
  })
}

// @Tags LogisticsPackaging
// @Summary 批量删除物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logisticsPackaging/deleteLogisticsPackaging [delete]
export const deleteLogisticsPackagingByIds = (data) => {
  return service({
    url: '/logisticsPackaging/deleteLogisticsPackagingByIds',
    method: 'delete',
    data
  })
}

// @Tags LogisticsPackaging
// @Summary 更新物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LogisticsPackaging true "更新物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /logisticsPackaging/updateLogisticsPackaging [put]
export const updateLogisticsPackaging = (data) => {
  return service({
    url: '/logisticsPackaging/updateLogisticsPackaging',
    method: 'put',
    data
  })
}

// @Tags LogisticsPackaging
// @Summary 用id查询物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LogisticsPackaging true "用id查询物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /logisticsPackaging/findLogisticsPackaging [get]
export const findLogisticsPackaging = (params) => {
  return service({
    url: '/logisticsPackaging/findLogisticsPackaging',
    method: 'get',
    params
  })
}

// @Tags LogisticsPackaging
// @Summary 分页获取物流包装列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取物流包装列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logisticsPackaging/getLogisticsPackagingList [get]
export const getLogisticsPackagingList = (params) => {
  return service({
    url: '/logisticsPackaging/getLogisticsPackagingList',
    method: 'get',
    params
  })
}
