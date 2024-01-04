package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SkuMappingRouter struct {
}

// InitSkuMappingRouter 初始化 商品映射 路由信息
func (s *SkuMappingRouter) InitSkuMappingRouter(Router *gin.RouterGroup) {
	skuMappingRouter := Router.Group("skuMapping").Use(middleware.OperationRecord())
	skuMappingRouterWithoutRecord := Router.Group("skuMapping")
	var skuMappingApi = v1.ApiGroupApp.WmsApiGroup.SkuMappingApi
	{
		skuMappingRouter.POST("createSkuMapping", skuMappingApi.CreateSkuMapping)   // 新建商品映射
		skuMappingRouter.DELETE("deleteSkuMapping", skuMappingApi.DeleteSkuMapping) // 删除商品映射
		skuMappingRouter.DELETE("deleteSkuMappingByIds", skuMappingApi.DeleteSkuMappingByIds) // 批量删除商品映射
		skuMappingRouter.PUT("updateSkuMapping", skuMappingApi.UpdateSkuMapping)    // 更新商品映射
	}
	{
		skuMappingRouterWithoutRecord.GET("findSkuMapping", skuMappingApi.FindSkuMapping)        // 根据ID获取商品映射
		skuMappingRouterWithoutRecord.GET("getSkuMappingList", skuMappingApi.GetSkuMappingList)  // 获取商品映射列表
	}
}
