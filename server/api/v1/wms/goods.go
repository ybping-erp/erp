package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoodsApi struct {
}

var goodsService = service.ServiceGroupApp.WmsServiceGroup.GoodsService

// CreateGoods 创建商品
// @Tags Goods
// @Summary 创建商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Goods true "创建商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goods/createGoods [post]
func (goodsApi *GoodsApi) CreateGoods(c *gin.Context) {
	var goods wms.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	goods.CreatedBy = utils.GetUserID(c)

	if err := goodsService.CreateGoods(&goods); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoods 删除商品
// @Tags Goods
// @Summary 删除商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Goods true "删除商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goods/deleteGoods [delete]
func (goodsApi *GoodsApi) DeleteGoods(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := goodsService.DeleteGoods(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsByIds 批量删除商品
// @Tags Goods
// @Summary 批量删除商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goods/deleteGoodsByIds [delete]
func (goodsApi *GoodsApi) DeleteGoodsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := goodsService.DeleteGoodsByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoods 更新商品
// @Tags Goods
// @Summary 更新商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Goods true "更新商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goods/updateGoods [put]
func (goodsApi *GoodsApi) UpdateGoods(c *gin.Context) {
	var goods wms.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	goods.UpdatedBy = utils.GetUserID(c)

	if err := goodsService.UpdateGoods(goods); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoods 用id查询商品
// @Tags Goods
// @Summary 用id查询商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.Goods true "用id查询商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goods/findGoods [get]
func (goodsApi *GoodsApi) FindGoods(c *gin.Context) {
	ID := c.Query("ID")
	if regoods, err := goodsService.GetGoods(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoods": regoods}, c)
	}
}

// GetGoodsList 分页获取商品列表
// @Tags Goods
// @Summary 分页获取商品列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.GoodsSearch true "分页获取商品列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/getGoodsList [get]
func (goodsApi *GoodsApi) GetGoodsList(c *gin.Context) {
	var pageInfo wmsReq.GoodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := goodsService.GetGoodsInfoList(pageInfo); err != nil {
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
