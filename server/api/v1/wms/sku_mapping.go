package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/wms"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SkuMappingApi struct {
}

var skuMappingService = service.ServiceGroupApp.WmsServiceGroup.SkuMappingService


// CreateSkuMapping 创建商品映射
// @Tags SkuMapping
// @Summary 创建商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.SkuMapping true "创建商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /skuMapping/createSkuMapping [post]
func (skuMappingApi *SkuMappingApi) CreateSkuMapping(c *gin.Context) {
	var skuMapping wms.SkuMapping
	err := c.ShouldBindJSON(&skuMapping)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    skuMapping.CreatedBy = utils.GetUserID(c)
	if err := skuMappingService.CreateSkuMapping(&skuMapping); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSkuMapping 删除商品映射
// @Tags SkuMapping
// @Summary 删除商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.SkuMapping true "删除商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /skuMapping/deleteSkuMapping [delete]
func (skuMappingApi *SkuMappingApi) DeleteSkuMapping(c *gin.Context) {
	var skuMapping wms.SkuMapping
	err := c.ShouldBindJSON(&skuMapping)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    skuMapping.DeletedBy = utils.GetUserID(c)
	if err := skuMappingService.DeleteSkuMapping(skuMapping); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSkuMappingByIds 批量删除商品映射
// @Tags SkuMapping
// @Summary 批量删除商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /skuMapping/deleteSkuMappingByIds [delete]
func (skuMappingApi *SkuMappingApi) DeleteSkuMappingByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := skuMappingService.DeleteSkuMappingByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSkuMapping 更新商品映射
// @Tags SkuMapping
// @Summary 更新商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.SkuMapping true "更新商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /skuMapping/updateSkuMapping [put]
func (skuMappingApi *SkuMappingApi) UpdateSkuMapping(c *gin.Context) {
	var skuMapping wms.SkuMapping
	err := c.ShouldBindJSON(&skuMapping)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    skuMapping.UpdatedBy = utils.GetUserID(c)
	if err := skuMappingService.UpdateSkuMapping(skuMapping); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSkuMapping 用id查询商品映射
// @Tags SkuMapping
// @Summary 用id查询商品映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.SkuMapping true "用id查询商品映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /skuMapping/findSkuMapping [get]
func (skuMappingApi *SkuMappingApi) FindSkuMapping(c *gin.Context) {
	var skuMapping wms.SkuMapping
	err := c.ShouldBindQuery(&skuMapping)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reskuMapping, err := skuMappingService.GetSkuMapping(skuMapping.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reskuMapping": reskuMapping}, c)
	}
}

// GetSkuMappingList 分页获取商品映射列表
// @Tags SkuMapping
// @Summary 分页获取商品映射列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.SkuMappingSearch true "分页获取商品映射列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /skuMapping/getSkuMappingList [get]
func (skuMappingApi *SkuMappingApi) GetSkuMappingList(c *gin.Context) {
	var pageInfo wmsReq.SkuMappingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := skuMappingService.GetSkuMappingInfoList(pageInfo); err != nil {
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
