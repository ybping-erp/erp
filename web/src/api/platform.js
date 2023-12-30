import service from '@/utils/request'

// @Tags Platform
// @Summary 创建电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platform true "创建电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /platform/createPlatform [post]
export const createPlatform = (data) => {
  return service({
    url: '/platform/createPlatform',
    method: 'post',
    data
  })
}

// @Tags Platform
// @Summary 删除电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platform true "删除电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platform/deletePlatform [delete]
export const deletePlatform = (data) => {
  return service({
    url: '/platform/deletePlatform',
    method: 'delete',
    data
  })
}

// @Tags Platform
// @Summary 批量删除电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platform/deletePlatform [delete]
export const deletePlatformByIds = (data) => {
  return service({
    url: '/platform/deletePlatformByIds',
    method: 'delete',
    data
  })
}

// @Tags Platform
// @Summary 更新电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platform true "更新电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /platform/updatePlatform [put]
export const updatePlatform = (data) => {
  return service({
    url: '/platform/updatePlatform',
    method: 'put',
    data
  })
}

// @Tags Platform
// @Summary 用id查询电商平台
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Platform true "用id查询电商平台"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /platform/findPlatform [get]
export const findPlatform = (params) => {
  return service({
    url: '/platform/findPlatform',
    method: 'get',
    params
  })
}

// @Tags Platform
// @Summary 分页获取电商平台列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取电商平台列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platform/getPlatformList [get]
export const getPlatformList = (params) => {
  return service({
    url: '/platform/getPlatformList',
    method: 'get',
    params
  })
}
