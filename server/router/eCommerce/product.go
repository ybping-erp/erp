package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

// InitProductRouter 初始化 产品表 路由信息
func (s *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouter := Router.Group("product").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("product")
	var productApi = v1.ApiGroupApp.ECommerceApiGroup.ProductApi
	{
		productRouter.POST("createProduct", productApi.CreateProduct)   // 新建产品表
		productRouter.DELETE("deleteProduct", productApi.DeleteProduct) // 删除产品表
		productRouter.DELETE("deleteProductByIds", productApi.DeleteProductByIds) // 批量删除产品表
		productRouter.PUT("updateProduct", productApi.UpdateProduct)    // 更新产品表
	}
	{
		productRouterWithoutRecord.GET("findProduct", productApi.FindProduct)        // 根据ID获取产品表
		productRouterWithoutRecord.GET("getProductList", productApi.GetProductList)  // 获取产品表列表
	}
}
