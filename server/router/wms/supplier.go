package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SupplierRouter struct {
}

// InitSupplierRouter 初始化 供应商 路由信息
func (s *SupplierRouter) InitSupplierRouter(Router *gin.RouterGroup) {
	supplierRouter := Router.Group("supplier").Use(middleware.OperationRecord())
	supplierRouterWithoutRecord := Router.Group("supplier")
	var supplierApi = v1.ApiGroupApp.WmsApiGroup.SupplierApi
	{
		supplierRouter.POST("createSupplier", supplierApi.CreateSupplier)   // 新建供应商
		supplierRouter.DELETE("deleteSupplier", supplierApi.DeleteSupplier) // 删除供应商
		supplierRouter.DELETE("deleteSupplierByIds", supplierApi.DeleteSupplierByIds) // 批量删除供应商
		supplierRouter.PUT("updateSupplier", supplierApi.UpdateSupplier)    // 更新供应商
	}
	{
		supplierRouterWithoutRecord.GET("findSupplier", supplierApi.FindSupplier)        // 根据ID获取供应商
		supplierRouterWithoutRecord.GET("getSupplierList", supplierApi.GetSupplierList)  // 获取供应商列表
	}
}
