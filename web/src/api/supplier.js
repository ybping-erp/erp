import service from '@/utils/request'

// @Tags Supplier
// @Summary 创建供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Supplier true "创建供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /supplier/createSupplier [post]
export const createSupplier = (data) => {
  return service({
    url: '/supplier/createSupplier',
    method: 'post',
    data
  })
}

// @Tags Supplier
// @Summary 删除供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Supplier true "删除供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /supplier/deleteSupplier [delete]
export const deleteSupplier = (params) => {
  return service({
    url: '/supplier/deleteSupplier',
    method: 'delete',
    params
  })
}

// @Tags Supplier
// @Summary 批量删除供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /supplier/deleteSupplier [delete]
export const deleteSupplierByIds = (params) => {
  return service({
    url: '/supplier/deleteSupplierByIds',
    method: 'delete',
    params
  })
}

// @Tags Supplier
// @Summary 更新供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Supplier true "更新供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /supplier/updateSupplier [put]
export const updateSupplier = (data) => {
  return service({
    url: '/supplier/updateSupplier',
    method: 'put',
    data
  })
}

// @Tags Supplier
// @Summary 用id查询供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Supplier true "用id查询供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /supplier/findSupplier [get]
export const findSupplier = (params) => {
  return service({
    url: '/supplier/findSupplier',
    method: 'get',
    params
  })
}

// @Tags Supplier
// @Summary 分页获取供应商列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取供应商列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /supplier/getSupplierList [get]
export const getSupplierList = (params) => {
  return service({
    url: '/supplier/getSupplierList',
    method: 'get',
    params
  })
}
