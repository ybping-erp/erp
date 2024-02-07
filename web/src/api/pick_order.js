import service from '@/utils/request'

// @Tags PickOrder
// @Summary 创建拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PickOrder true "创建拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pickOrder/createPickOrder [post]
export const createPickOrder = (data) => {
  return service({
    url: '/pickOrder/createPickOrder',
    method: 'post',
    data
  })
}

// @Tags PickOrder
// @Summary 删除拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PickOrder true "删除拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pickOrder/deletePickOrder [delete]
export const deletePickOrder = (params) => {
  return service({
    url: '/pickOrder/deletePickOrder',
    method: 'delete',
    params
  })
}

// @Tags PickOrder
// @Summary 批量删除拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pickOrder/deletePickOrder [delete]
export const deletePickOrderByIds = (params) => {
  return service({
    url: '/pickOrder/deletePickOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags PickOrder
// @Summary 更新拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PickOrder true "更新拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pickOrder/updatePickOrder [put]
export const updatePickOrder = (data) => {
  return service({
    url: '/pickOrder/updatePickOrder',
    method: 'put',
    data
  })
}

// @Tags PickOrder
// @Summary 用id查询拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PickOrder true "用id查询拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pickOrder/findPickOrder [get]
export const findPickOrder = (params) => {
  return service({
    url: '/pickOrder/findPickOrder',
    method: 'get',
    params
  })
}

// @Tags PickOrder
// @Summary 分页获取拣货单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取拣货单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pickOrder/getPickOrderList [get]
export const getPickOrderList = (params) => {
  return service({
    url: '/pickOrder/getPickOrderList',
    method: 'get',
    params
  })
}
