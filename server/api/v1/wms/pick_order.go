package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/wms"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PickOrderApi struct {
}

var pickOrderService = service.ServiceGroupApp.WmsServiceGroup.PickOrderService


// CreatePickOrder 创建拣货单
// @Tags PickOrder
// @Summary 创建拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.PickOrder true "创建拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pickOrder/createPickOrder [post]
func (pickOrderApi *PickOrderApi) CreatePickOrder(c *gin.Context) {
	var pickOrder wms.PickOrder
	err := c.ShouldBindJSON(&pickOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    pickOrder.CreatedBy = utils.GetUserID(c)

	if err := pickOrderService.CreatePickOrder(&pickOrder); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePickOrder 删除拣货单
// @Tags PickOrder
// @Summary 删除拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.PickOrder true "删除拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pickOrder/deletePickOrder [delete]
func (pickOrderApi *PickOrderApi) DeletePickOrder(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := pickOrderService.DeletePickOrder(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePickOrderByIds 批量删除拣货单
// @Tags PickOrder
// @Summary 批量删除拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pickOrder/deletePickOrderByIds [delete]
func (pickOrderApi *PickOrderApi) DeletePickOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := pickOrderService.DeletePickOrderByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePickOrder 更新拣货单
// @Tags PickOrder
// @Summary 更新拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.PickOrder true "更新拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pickOrder/updatePickOrder [put]
func (pickOrderApi *PickOrderApi) UpdatePickOrder(c *gin.Context) {
	var pickOrder wms.PickOrder
	err := c.ShouldBindJSON(&pickOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    pickOrder.UpdatedBy = utils.GetUserID(c)

	if err := pickOrderService.UpdatePickOrder(pickOrder); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPickOrder 用id查询拣货单
// @Tags PickOrder
// @Summary 用id查询拣货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.PickOrder true "用id查询拣货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pickOrder/findPickOrder [get]
func (pickOrderApi *PickOrderApi) FindPickOrder(c *gin.Context) {
	ID := c.Query("ID")
	if repickOrder, err := pickOrderService.GetPickOrder(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repickOrder": repickOrder}, c)
	}
}

// GetPickOrderList 分页获取拣货单列表
// @Tags PickOrder
// @Summary 分页获取拣货单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.PickOrderSearch true "分页获取拣货单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pickOrder/getPickOrderList [get]
func (pickOrderApi *PickOrderApi) GetPickOrderList(c *gin.Context) {
	var pageInfo wmsReq.PickOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := pickOrderService.GetPickOrderInfoList(pageInfo); err != nil {
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
