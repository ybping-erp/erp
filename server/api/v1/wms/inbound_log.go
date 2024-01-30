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

type InboundLogApi struct {
}

var inboundLogService = service.ServiceGroupApp.WmsServiceGroup.InboundLogService

// CreateInboundLog 创建入库记录
// @Tags InboundLog
// @Summary 创建入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.InboundLog true "创建入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /inboundLog/createInboundLog [post]
func (inboundLogApi *InboundLogApi) CreateInboundLog(c *gin.Context) {
	var inboundLog wms.InboundLog
	err := c.ShouldBindJSON(&inboundLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	inboundLog.CreatedBy = utils.GetUserID(c)

	if err := inboundLogService.CreateInboundLog(&inboundLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInboundLog 删除入库记录
// @Tags InboundLog
// @Summary 删除入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.InboundLog true "删除入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inboundLog/deleteInboundLog [delete]
func (inboundLogApi *InboundLogApi) DeleteInboundLog(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := inboundLogService.DeleteInboundLog(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInboundLogByIds 批量删除入库记录
// @Tags InboundLog
// @Summary 批量删除入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /inboundLog/deleteInboundLogByIds [delete]
func (inboundLogApi *InboundLogApi) DeleteInboundLogByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := inboundLogService.DeleteInboundLogByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInboundLog 更新入库记录
// @Tags InboundLog
// @Summary 更新入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wms.InboundLog true "更新入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inboundLog/updateInboundLog [put]
func (inboundLogApi *InboundLogApi) UpdateInboundLog(c *gin.Context) {
	var inboundLog wms.InboundLog
	err := c.ShouldBindJSON(&inboundLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	inboundLog.UpdatedBy = utils.GetUserID(c)

	if err := inboundLogService.UpdateInboundLog(inboundLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInboundLog 用id查询入库记录
// @Tags InboundLog
// @Summary 用id查询入库记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wms.InboundLog true "用id查询入库记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inboundLog/findInboundLog [get]
func (inboundLogApi *InboundLogApi) FindInboundLog(c *gin.Context) {
	ID := c.Query("ID")
	if reinboundLog, err := inboundLogService.GetInboundLog(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinboundLog": reinboundLog}, c)
	}
}

// GetInboundLogList 分页获取入库记录列表
// @Tags InboundLog
// @Summary 分页获取入库记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wmsReq.InboundLogSearch true "分页获取入库记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inboundLog/getInboundLogList [get]
func (inboundLogApi *InboundLogApi) GetInboundLogList(c *gin.Context) {
	var pageInfo wmsReq.InboundLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := inboundLogService.GetInboundLogInfoList(pageInfo); err != nil {
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
