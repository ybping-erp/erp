package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlatformRouter struct {
}

// InitPlatformRouter 初始化 电商平台 路由信息
func (s *PlatformRouter) InitPlatformRouter(Router *gin.RouterGroup) {
	platformRouter := Router.Group("platform").Use(middleware.OperationRecord())
	platformRouterWithoutRecord := Router.Group("platform")
	var platformApi = v1.ApiGroupApp.ECommerceApiGroup.PlatformApi
	{
		platformRouter.POST("createPlatform", platformApi.CreatePlatform)   // 新建电商平台
		platformRouter.DELETE("deletePlatform", platformApi.DeletePlatform) // 删除电商平台
		platformRouter.DELETE("deletePlatformByIds", platformApi.DeletePlatformByIds) // 批量删除电商平台
		platformRouter.PUT("updatePlatform", platformApi.UpdatePlatform)    // 更新电商平台
	}
	{
		platformRouterWithoutRecord.GET("findPlatform", platformApi.FindPlatform)        // 根据ID获取电商平台
		platformRouterWithoutRecord.GET("getPlatformList", platformApi.GetPlatformList)  // 获取电商平台列表
	}
}
