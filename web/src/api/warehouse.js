import service from '@/utils/request'

// @Tags Warehouse
// @Summary 创建仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Warehouse true "创建仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /warehouse/createWarehouse [post]
export const createWarehouse = (data) => {
  return service({
    url: '/warehouse/createWarehouse',
    method: 'post',
    data
  })
}

// @Tags Warehouse
// @Summary 删除仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Warehouse true "删除仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /warehouse/deleteWarehouse [delete]
export const deleteWarehouse = (data) => {
  return service({
    url: '/warehouse/deleteWarehouse',
    method: 'delete',
    data
  })
}

// @Tags Warehouse
// @Summary 批量删除仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /warehouse/deleteWarehouse [delete]
export const deleteWarehouseByIds = (data) => {
  return service({
    url: '/warehouse/deleteWarehouseByIds',
    method: 'delete',
    data
  })
}

// @Tags Warehouse
// @Summary 更新仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Warehouse true "更新仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /warehouse/updateWarehouse [put]
export const updateWarehouse = (data) => {
  return service({
    url: '/warehouse/updateWarehouse',
    method: 'put',
    data
  })
}

// @Tags Warehouse
// @Summary 用id查询仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Warehouse true "用id查询仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /warehouse/findWarehouse [get]
export const findWarehouse = (params) => {
  return service({
    url: '/warehouse/findWarehouse',
    method: 'get',
    params
  })
}

// @Tags Warehouse
// @Summary 分页获取仓库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取仓库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /warehouse/getWarehouseList [get]
export const getWarehouseList = (params) => {
  return service({
    url: '/warehouse/getWarehouseList',
    method: 'get',
    params
  })
}
