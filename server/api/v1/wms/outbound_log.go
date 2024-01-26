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

type OutboundLogApi struct {
}

var outboundLogService = service.ServiceGroupApp.WmsServiceGroup.OutboundLogService


// CreateOutboundLog 创建出库记录
// @Tags OutboundLog
// @Summary 创建出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.OutboundLog true "创建出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /outboundLog/createOutboundLog [post]
func (outboundLogApi *OutboundLogApi) CreateOutboundLog(c *gin.Context) {
	var outboundLog wms.OutboundLog
	err := c.ShouldBindJSON(&outboundLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    outboundLog.CreatedBy = utils.GetUserID(c)

	if err := outboundLogService.CreateOutboundLog(&outboundLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOutboundLog 删除出库记录
// @Tags OutboundLog
// @Summary 删除出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.OutboundLog true "删除出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outboundLog/deleteOutboundLog [delete]
func (outboundLogApi *OutboundLogApi) DeleteOutboundLog(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := outboundLogService.DeleteOutboundLog(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOutboundLogByIds 批量删除出库记录
// @Tags OutboundLog
// @Summary 批量删除出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /outboundLog/deleteOutboundLogByIds [delete]
func (outboundLogApi *OutboundLogApi) DeleteOutboundLogByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := outboundLogService.DeleteOutboundLogByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOutboundLog 更新出库记录
// @Tags OutboundLog
// @Summary 更新出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.OutboundLog true "更新出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /outboundLog/updateOutboundLog [put]
func (outboundLogApi *OutboundLogApi) UpdateOutboundLog(c *gin.Context) {
	var outboundLog wms.OutboundLog
	err := c.ShouldBindJSON(&outboundLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    outboundLog.UpdatedBy = utils.GetUserID(c)

	if err := outboundLogService.UpdateOutboundLog(outboundLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOutboundLog 用id查询出库记录
// @Tags OutboundLog
// @Summary 用id查询出库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.OutboundLog true "用id查询出库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /outboundLog/findOutboundLog [get]
func (outboundLogApi *OutboundLogApi) FindOutboundLog(c *gin.Context) {
	ID := c.Query("ID")
	if reoutboundLog, err := outboundLogService.GetOutboundLog(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoutboundLog": reoutboundLog}, c)
	}
}

// GetOutboundLogList 分页获取出库记录列表
// @Tags OutboundLog
// @Summary 分页获取出库记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.OutboundLogSearch true "分页获取出库记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outboundLog/getOutboundLogList [get]
func (outboundLogApi *OutboundLogApi) GetOutboundLogList(c *gin.Context) {
	var pageInfo wmsReq.OutboundLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := outboundLogService.GetOutboundLogInfoList(pageInfo); err != nil {
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
