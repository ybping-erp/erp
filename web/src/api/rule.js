import service from '@/utils/request'

// @Tags Rule
// @Summary 创建规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rule true "创建规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rule/createRule [post]
export const createRule = (data) => {
  return service({
    url: '/rule/createRule',
    method: 'post',
    data
  })
}

// @Tags Rule
// @Summary 删除规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rule true "删除规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rule/deleteRule [delete]
export const deleteRule = (params) => {
  return service({
    url: '/rule/deleteRule',
    method: 'delete',
    params
  })
}

// @Tags Rule
// @Summary 批量删除规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rule/deleteRule [delete]
export const deleteRuleByIds = (params) => {
  return service({
    url: '/rule/deleteRuleByIds',
    method: 'delete',
    params
  })
}

// @Tags Rule
// @Summary 更新规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rule true "更新规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rule/updateRule [put]
export const updateRule = (data) => {
  return service({
    url: '/rule/updateRule',
    method: 'put',
    data
  })
}

// @Tags Rule
// @Summary 用id查询规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Rule true "用id查询规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rule/findRule [get]
export const findRule = (params) => {
  return service({
    url: '/rule/findRule',
    method: 'get',
    params
  })
}

// @Tags Rule
// @Summary 分页获取规则列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取规则列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rule/getRuleList [get]
export const getRuleList = (params) => {
  return service({
    url: '/rule/getRuleList',
    method: 'get',
    params
  })
}
