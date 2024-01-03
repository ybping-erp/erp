package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
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
	if err := shopService.DeleteShop(shop); err != nil {
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
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopService.DeleteShopByIds(IDS); err != nil {
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
