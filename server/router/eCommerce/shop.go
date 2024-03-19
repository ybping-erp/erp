package eCommerce

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopRouter struct {
}

// InitShopRouter 初始化 店铺 路由信息
func (s *ShopRouter) InitShopRouter(Router *gin.RouterGroup) {
	shopRouter := Router.Group("shop").Use(middleware.OperationRecord())
	shopRouterWithoutRecord := Router.Group("shop")
	var shopApi = v1.ApiGroupApp.ECommerceApiGroup.ShopApi
	{
		shopRouter.POST("createShop", shopApi.CreateShop)             // 新建店铺
		shopRouter.DELETE("deleteShop", shopApi.DeleteShop)           // 删除店铺
		shopRouter.DELETE("deleteShopByIds", shopApi.DeleteShopByIds) // 批量删除店铺
		shopRouter.PUT("updateShop", shopApi.UpdateShop)              // 更新店铺
		shopRouter.POST("authorize", shopApi.AuthorizeShop)           // 店铺第三方授权
	}
	{
		shopRouterWithoutRecord.GET("findShop", shopApi.FindShop)       // 根据ID获取店铺
		shopRouterWithoutRecord.GET("getShopList", shopApi.GetShopList) // 获取店铺列表
	}
}

func (s *ShopRouter) InitShopPublicRouter(Router *gin.RouterGroup) {
	var shopApi = v1.ApiGroupApp.ECommerceApiGroup.ShopApi
	shopRouter := Router.Group("shop").Use(middleware.OperationRecord())
	shopRouter.GET("authorize/callback", shopApi.OAuthCallback) // 店铺第三方授权回调
}
