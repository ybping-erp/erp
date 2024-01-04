package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/wms"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type WarehouseApi struct {
}

var warehouseService = service.ServiceGroupApp.WmsServiceGroup.WarehouseService


// CreateWarehouse 创建仓库
// @Tags Warehouse
// @Summary 创建仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Warehouse true "创建仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /warehouse/createWarehouse [post]
func (warehouseApi *WarehouseApi) CreateWarehouse(c *gin.Context) {
	var warehouse wms.Warehouse
	err := c.ShouldBindJSON(&warehouse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    warehouse.CreatedBy = utils.GetUserID(c)
	if err := warehouseService.CreateWarehouse(&warehouse); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWarehouse 删除仓库
// @Tags Warehouse
// @Summary 删除仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Warehouse true "删除仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /warehouse/deleteWarehouse [delete]
func (warehouseApi *WarehouseApi) DeleteWarehouse(c *gin.Context) {
	var warehouse wms.Warehouse
	err := c.ShouldBindJSON(&warehouse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    warehouse.DeletedBy = utils.GetUserID(c)
	if err := warehouseService.DeleteWarehouse(warehouse); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWarehouseByIds 批量删除仓库
// @Tags Warehouse
// @Summary 批量删除仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /warehouse/deleteWarehouseByIds [delete]
func (warehouseApi *WarehouseApi) DeleteWarehouseByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := warehouseService.DeleteWarehouseByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWarehouse 更新仓库
// @Tags Warehouse
// @Summary 更新仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Warehouse true "更新仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /warehouse/updateWarehouse [put]
func (warehouseApi *WarehouseApi) UpdateWarehouse(c *gin.Context) {
	var warehouse wms.Warehouse
	err := c.ShouldBindJSON(&warehouse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    warehouse.UpdatedBy = utils.GetUserID(c)
	if err := warehouseService.UpdateWarehouse(warehouse); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWarehouse 用id查询仓库
// @Tags Warehouse
// @Summary 用id查询仓库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.Warehouse true "用id查询仓库"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /warehouse/findWarehouse [get]
func (warehouseApi *WarehouseApi) FindWarehouse(c *gin.Context) {
	var warehouse wms.Warehouse
	err := c.ShouldBindQuery(&warehouse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewarehouse, err := warehouseService.GetWarehouse(warehouse.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewarehouse": rewarehouse}, c)
	}
}

// GetWarehouseList 分页获取仓库列表
// @Tags Warehouse
// @Summary 分页获取仓库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.WarehouseSearch true "分页获取仓库列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /warehouse/getWarehouseList [get]
func (warehouseApi *WarehouseApi) GetWarehouseList(c *gin.Context) {
	var pageInfo wmsReq.WarehouseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := warehouseService.GetWarehouseInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
