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
)

type LogisticsPackagingApi struct {
}

var logisticsPackagingService = service.ServiceGroupApp.WmsServiceGroup.LogisticsPackagingService


// CreateLogisticsPackaging 创建物流包装
// @Tags LogisticsPackaging
// @Summary 创建物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.LogisticsPackaging true "创建物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /logisticsPackaging/createLogisticsPackaging [post]
func (logisticsPackagingApi *LogisticsPackagingApi) CreateLogisticsPackaging(c *gin.Context) {
	var logisticsPackaging wms.LogisticsPackaging
	err := c.ShouldBindJSON(&logisticsPackaging)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logisticsPackagingService.CreateLogisticsPackaging(&logisticsPackaging); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLogisticsPackaging 删除物流包装
// @Tags LogisticsPackaging
// @Summary 删除物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.LogisticsPackaging true "删除物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logisticsPackaging/deleteLogisticsPackaging [delete]
func (logisticsPackagingApi *LogisticsPackagingApi) DeleteLogisticsPackaging(c *gin.Context) {
	var logisticsPackaging wms.LogisticsPackaging
	err := c.ShouldBindJSON(&logisticsPackaging)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logisticsPackagingService.DeleteLogisticsPackaging(logisticsPackaging); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLogisticsPackagingByIds 批量删除物流包装
// @Tags LogisticsPackaging
// @Summary 批量删除物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /logisticsPackaging/deleteLogisticsPackagingByIds [delete]
func (logisticsPackagingApi *LogisticsPackagingApi) DeleteLogisticsPackagingByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logisticsPackagingService.DeleteLogisticsPackagingByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLogisticsPackaging 更新物流包装
// @Tags LogisticsPackaging
// @Summary 更新物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.LogisticsPackaging true "更新物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /logisticsPackaging/updateLogisticsPackaging [put]
func (logisticsPackagingApi *LogisticsPackagingApi) UpdateLogisticsPackaging(c *gin.Context) {
	var logisticsPackaging wms.LogisticsPackaging
	err := c.ShouldBindJSON(&logisticsPackaging)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logisticsPackagingService.UpdateLogisticsPackaging(logisticsPackaging); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLogisticsPackaging 用id查询物流包装
// @Tags LogisticsPackaging
// @Summary 用id查询物流包装
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.LogisticsPackaging true "用id查询物流包装"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /logisticsPackaging/findLogisticsPackaging [get]
func (logisticsPackagingApi *LogisticsPackagingApi) FindLogisticsPackaging(c *gin.Context) {
	var logisticsPackaging wms.LogisticsPackaging
	err := c.ShouldBindQuery(&logisticsPackaging)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relogisticsPackaging, err := logisticsPackagingService.GetLogisticsPackaging(logisticsPackaging.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relogisticsPackaging": relogisticsPackaging}, c)
	}
}

// GetLogisticsPackagingList 分页获取物流包装列表
// @Tags LogisticsPackaging
// @Summary 分页获取物流包装列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.LogisticsPackagingSearch true "分页获取物流包装列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logisticsPackaging/getLogisticsPackagingList [get]
func (logisticsPackagingApi *LogisticsPackagingApi) GetLogisticsPackagingList(c *gin.Context) {
	var pageInfo wmsReq.LogisticsPackagingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := logisticsPackagingService.GetLogisticsPackagingInfoList(pageInfo); err != nil {
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
