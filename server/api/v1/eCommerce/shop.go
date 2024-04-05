package eCommerce

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/eBay"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopApi struct {
}

var shopService = service.ServiceGroupApp.ECommerceServiceGroup.ShopService

// CreateShop 创建店铺
// @Tags Shop
// @Summary 创建店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Shop true "创建店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shop/createShop [post]
func (shopApi *ShopApi) CreateShop(c *gin.Context) {
	var shop eCommerce.Shop
	err := c.ShouldBindJSON(&shop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shop.CreatedBy = utils.GetUserID(c)
	if err := shopService.CreateShop(&shop); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShop 删除店铺
// @Tags Shop
// @Summary 删除店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Shop true "删除店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shop/deleteShop [delete]
func (shopApi *ShopApi) DeleteShop(c *gin.Context) {
	var shop eCommerce.Shop
	err := c.ShouldBindJSON(&shop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopService.DeleteShopByIds([]int{int(shop.ID)}, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShopByIds 批量删除店铺
// @Tags Shop
// @Summary 批量删除店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shop/deleteShopByIds [delete]
func (shopApi *ShopApi) DeleteShopByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := shopService.DeleteShopByIds(req.Ids, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShop 更新店铺
// @Tags Shop
// @Summary 更新店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Shop true "更新店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shop/updateShop [put]
func (shopApi *ShopApi) UpdateShop(c *gin.Context) {
	var shop eCommerce.Shop
	err := c.ShouldBindJSON(&shop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shop.UpdatedBy = utils.GetUserID(c)
	if err := shopService.UpdateShop(shop); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShop 用id查询店铺
// @Tags Shop
// @Summary 用id查询店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eCommerce.Shop true "用id查询店铺"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shop/findShop [get]
func (shopApi *ShopApi) FindShop(c *gin.Context) {
	var shop eCommerce.Shop
	err := c.ShouldBindQuery(&shop)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reshop, err := shopService.GetShop(shop.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshop": reshop}, c)
	}
}

// GetShopList 分页获取店铺列表
// @Tags Shop
// @Summary 分页获取店铺列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eCommerceReq.ShopSearch true "分页获取店铺列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shop/getShopList [get]
func (shopApi *ShopApi) GetShopList(c *gin.Context) {
	var pageInfo eCommerceReq.ShopSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopService.GetShopInfoList(pageInfo); err != nil {
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

// GetShopAccessToken 店铺授权
// @Tags Shop
// @Summary 店铺授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Shop true "店铺授权"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"授权成功"}"
// @Router /shop/authorize [post]
func (shopApi *ShopApi) AuthorizeShop(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var shop eCommerce.Shop
	if shop, err = shopService.GetShop(req.Uint()); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	// 生成SCRF Token并存储，用来在oAuth回调时进行授权凭证对比, 对比后重置该Token
	err = shop.GenerateOAuthState()
	if err != nil {
		global.GVA_LOG.Error("生成CSRFToken失败,请重试!", zap.Error(err))
		response.FailWithMessage("授权失败,请重试!", c)
		return
	}

	if err := shopService.UpdateShop(shop); err != nil {
		global.GVA_LOG.Error("更新CSRFToken失败!", zap.Error(err))
		response.FailWithMessage("授权失败，请重试!", c)
		return
	}

	if shop.PlatformName == "eBay" {
		client := eBay.GetAPIClient()

		// Test Refresh Code
		if shop.RefreshToken != "" {
			token, oauthClient, err := client.GetAccessTokenAndClient(c, shop.RefreshToken, client.GetDefaultScopes())
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}

			// Test list orders
			// err = client.GetOrders(oauthClient)

			// Test create product

			err = client.CreateOrReplaceInventoryItem(oauthClient)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}

			response.FailWithMessage(token.AccessToken, c)
			return
		}
		// Test Refresh Code End

		if oAuthUrl, err := client.GenerateUserAuthorizationURL(c, shop.CSRFToken, client.GetDefaultScopes()); err != nil {
			global.GVA_LOG.Error("eBay.GenerateUserAuthorizationURL 失败!", zap.Error(err))
			response.FailWithMessage("授权失败!", c)
		} else {
			response.OkWithData(gin.H{"oAuthUrl": oAuthUrl}, c)
		}
		return
	}

	response.FailWithMessage("不支持该平台授权", c)
}

// EBayOAuthCallback 店铺授权回调
// @Tags Shop
// @Summary 店铺授权
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"授权成功"}"
// @Router /shop/authorize/callback [get]
func (shopApi *ShopApi) OAuthCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	// 解析state参数，防止CSRF攻击
	shopID, err := eCommerce.Shop{}.GetShopIDFromOAuthState(state)
	if err != nil {
		global.GVA_LOG.Error("非法授权回调！", zap.Error(err))
		response.FailWithMessage("非法请求", c)
		return
	}

	// 应该有且只有一个合法的授权店铺存在
	shop, err := shopService.GetShop(uint(shopID))
	if err != nil {
		global.GVA_LOG.Error("非法授权回调！", zap.Error(err))
		response.FailWithMessage("非法请求", c)
		return
	}

	if shop.CSRFToken != state {
		global.GVA_LOG.Error("非法授权回调！")
		response.FailWithMessage("非法请求", c)
		return
	}

	// 用access code 换取 oauth token
	client := eBay.GetAPIClient()
	if token, _, err := client.ExchangeCodeForAccessTokenAndClient(c, code); err != nil {
		global.GVA_LOG.Error("ExchangeCodeForAccessTokenAndClient失败!", zap.Error(err))
		response.FailWithMessage("授权失败，请重试!", c)
		return
	} else {
		shop.ResetToken(token)
		if err := shopService.UpdateShop(shop); err != nil {
			global.GVA_LOG.Error("重置Token失败!", zap.Error(err))
			response.FailWithMessage("授权失败，请重试!", c)
			return
		}
	}

	// 授权成功重定向到店铺页面
	c.Redirect(http.StatusFound, "http://localhost:8080/#/layout/eCommerce/shop")
}
