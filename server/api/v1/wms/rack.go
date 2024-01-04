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

type RackApi struct {
}

var rackService = service.ServiceGroupApp.WmsServiceGroup.RackService


// CreateRack 创建货架
// @Tags Rack
// @Summary 创建货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Rack true "创建货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rack/createRack [post]
func (rackApi *RackApi) CreateRack(c *gin.Context) {
	var rack wms.Rack
	err := c.ShouldBindJSON(&rack)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    rack.CreatedBy = utils.GetUserID(c)
	if err := rackService.CreateRack(&rack); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRack 删除货架
// @Tags Rack
// @Summary 删除货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Rack true "删除货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rack/deleteRack [delete]
func (rackApi *RackApi) DeleteRack(c *gin.Context) {
	var rack wms.Rack
	err := c.ShouldBindJSON(&rack)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    rack.DeletedBy = utils.GetUserID(c)
	if err := rackService.DeleteRack(rack); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRackByIds 批量删除货架
// @Tags Rack
// @Summary 批量删除货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rack/deleteRackByIds [delete]
func (rackApi *RackApi) DeleteRackByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := rackService.DeleteRackByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRack 更新货架
// @Tags Rack
// @Summary 更新货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.Rack true "更新货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rack/updateRack [put]
func (rackApi *RackApi) UpdateRack(c *gin.Context) {
	var rack wms.Rack
	err := c.ShouldBindJSON(&rack)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    rack.UpdatedBy = utils.GetUserID(c)
	if err := rackService.UpdateRack(rack); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRack 用id查询货架
// @Tags Rack
// @Summary 用id查询货架
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.Rack true "用id查询货架"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rack/findRack [get]
func (rackApi *RackApi) FindRack(c *gin.Context) {
	var rack wms.Rack
	err := c.ShouldBindQuery(&rack)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerack, err := rackService.GetRack(rack.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerack": rerack}, c)
	}
}

// GetRackList 分页获取货架列表
// @Tags Rack
// @Summary 分页获取货架列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.RackSearch true "分页获取货架列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rack/getRackList [get]
func (rackApi *RackApi) GetRackList(c *gin.Context) {
	var pageInfo wmsReq.RackSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := rackService.GetRackInfoList(pageInfo); err != nil {
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
