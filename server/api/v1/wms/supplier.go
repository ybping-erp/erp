package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/wms"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SupplierApi struct {
}

var supplierService = service.ServiceGroupApp.WmsServiceGroup.SupplierService


// CreateSupplier 创建供应商
// @Tags Supplier
// @Summary 创建供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Supplier true "创建供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /supplier/createSupplier [post]
func (supplierApi *SupplierApi) CreateSupplier(c *gin.Context) {
	var supplier wms.Supplier
	err := c.ShouldBindJSON(&supplier)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    supplier.CreatedBy = utils.GetUserID(c)

	if err := supplierService.CreateSupplier(&supplier); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSupplier 删除供应商
// @Tags Supplier
// @Summary 删除供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Supplier true "删除供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /supplier/deleteSupplier [delete]
func (supplierApi *SupplierApi) DeleteSupplier(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := supplierService.DeleteSupplier(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSupplierByIds 批量删除供应商
// @Tags Supplier
// @Summary 批量删除供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /supplier/deleteSupplierByIds [delete]
func (supplierApi *SupplierApi) DeleteSupplierByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := supplierService.DeleteSupplierByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSupplier 更新供应商
// @Tags Supplier
// @Summary 更新供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Supplier true "更新供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /supplier/updateSupplier [put]
func (supplierApi *SupplierApi) UpdateSupplier(c *gin.Context) {
	var supplier wms.Supplier
	err := c.ShouldBindJSON(&supplier)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    supplier.UpdatedBy = utils.GetUserID(c)

	if err := supplierService.UpdateSupplier(supplier); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSupplier 用id查询供应商
// @Tags Supplier
// @Summary 用id查询供应商
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.Supplier true "用id查询供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /supplier/findSupplier [get]
func (supplierApi *SupplierApi) FindSupplier(c *gin.Context) {
	ID := c.Query("ID")
	if resupplier, err := supplierService.GetSupplier(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resupplier": resupplier}, c)
	}
}

// GetSupplierList 分页获取供应商列表
// @Tags Supplier
// @Summary 分页获取供应商列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.SupplierSearch true "分页获取供应商列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /supplier/getSupplierList [get]
func (supplierApi *SupplierApi) GetSupplierList(c *gin.Context) {
	var pageInfo wmsReq.SupplierSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := supplierService.GetSupplierInfoList(pageInfo); err != nil {
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
