package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsRouter struct {
}

// InitGoodsRouter 初始化 商品 路由信息
func (s *GoodsRouter) InitGoodsRouter(Router *gin.RouterGroup) {
	goodsRouter := Router.Group("goods").Use(middleware.OperationRecord())
	goodsRouterWithoutRecord := Router.Group("goods")
	var goodsApi = v1.ApiGroupApp.WmsApiGroup.GoodsApi
	{
		goodsRouter.POST("createGoods", goodsApi.CreateGoods)   // 新建商品
		goodsRouter.DELETE("deleteGoods", goodsApi.DeleteGoods) // 删除商品
		goodsRouter.DELETE("deleteGoodsByIds", goodsApi.DeleteGoodsByIds) // 批量删除商品
		goodsRouter.PUT("updateGoods", goodsApi.UpdateGoods)    // 更新商品
	}
	{
		goodsRouterWithoutRecord.GET("findGoods", goodsApi.FindGoods)        // 根据ID获取商品
		goodsRouterWithoutRecord.GET("getGoodsList", goodsApi.GetGoodsList)  // 获取商品列表
	}
}
