package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

// InitCategoryRouter 初始化 品类 路由信息
func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category").Use(middleware.OperationRecord())
	categoryRouterWithoutRecord := Router.Group("category")
	var categoryApi = v1.ApiGroupApp.ECommerceApiGroup.CategoryApi
	{
		categoryRouter.POST("createCategory", categoryApi.CreateCategory)   // 新建品类
		categoryRouter.DELETE("deleteCategory", categoryApi.DeleteCategory) // 删除品类
		categoryRouter.DELETE("deleteCategoryByIds", categoryApi.DeleteCategoryByIds) // 批量删除品类
		categoryRouter.PUT("updateCategory", categoryApi.UpdateCategory)    // 更新品类
	}
	{
		categoryRouterWithoutRecord.GET("findCategory", categoryApi.FindCategory)        // 根据ID获取品类
		categoryRouterWithoutRecord.GET("getCategoryList", categoryApi.GetCategoryList)  // 获取品类列表
	}
}
