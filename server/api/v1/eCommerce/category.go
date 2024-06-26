package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
	wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryApi struct {
}

var categoryService = service.ServiceGroupApp.ECommerceServiceGroup.CategoryService
var goodsService = service.ServiceGroupApp.WmsServiceGroup.GoodsService

// CreateCategory 创建品类
// @Tags Category
// @Summary 创建品类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Category true "创建品类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /category/createCategory [post]
func (categoryApi *CategoryApi) CreateCategory(c *gin.Context) {
	var category eCommerce.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	category.CreatedBy = utils.GetUserID(c)
	if err := categoryService.CreateCategory(&category); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCategory 删除品类
// @Tags Category
// @Summary 删除品类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Category true "删除品类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteCategory [delete]
func (categoryApi *CategoryApi) DeleteCategory(c *gin.Context) {
	var category eCommerce.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	categoryID, total := (int)(category.ID), int64(0)

	// 有商品绑定了该类别，不能删除
	_, total, err = goodsService.GetGoodsInfoList(wmsReq.GoodsSearch{
		CategoryId: &categoryID,
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 10,
		},
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if total > 0 {
		response.FailWithMessage("有商品绑定了该类别，删除失败", c)
		return
	}

	// 有产品绑定了该类别，不能删除
	_, total, err = productService.GetProductInfoList(eCommerceReq.ProductSearch{
		Product:  eCommerce.Product{CategoryId: &categoryID},
		PageInfo: request.PageInfo{Page: 1, PageSize: 10},
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if total > 0 {
		response.FailWithMessage("有产品绑定了该类别，删除失败", c)
		return
	}

	userID := utils.GetUserID(c)
	if err := categoryService.DeleteCategoryByIds([]int{int(category.ID)}, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCategoryByIds 批量删除品类
// @Tags Category
// @Summary 批量删除品类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除品类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /category/deleteCategoryByIds [delete]
func (categoryApi *CategoryApi) DeleteCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := categoryService.DeleteCategoryByIds(IDS.Ids, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCategory 更新品类
// @Tags Category
// @Summary 更新品类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eCommerce.Category true "更新品类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category/updateCategory [put]
func (categoryApi *CategoryApi) UpdateCategory(c *gin.Context) {
	var category eCommerce.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	category.UpdatedBy = utils.GetUserID(c)
	if err := categoryService.UpdateCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCategory 用id查询品类
// @Tags Category
// @Summary 用id查询品类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eCommerce.Category true "用id查询品类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category/findCategory [get]
func (categoryApi *CategoryApi) FindCategory(c *gin.Context) {
	var category eCommerce.Category
	err := c.ShouldBindQuery(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recategory, err := categoryService.GetCategory(category.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recategory": recategory}, c)
	}
}

// GetCategoryList 分页获取品类列表
// @Tags Category
// @Summary 分页获取品类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eCommerceReq.CategorySearch true "分页获取品类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getCategoryList [get]
func (categoryApi *CategoryApi) GetCategoryList(c *gin.Context) {
	var pageInfo eCommerceReq.CategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := categoryService.GetCategoryInfoList(pageInfo); err != nil {
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
