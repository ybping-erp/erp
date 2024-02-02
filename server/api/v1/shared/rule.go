package shared

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shared"
	sharedReq "github.com/flipped-aurora/gin-vue-admin/server/model/shared/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RuleApi struct {
}

var ruleService = service.ServiceGroupApp.SharedServiceGroup.RuleService

// CreateRule 创建规则
// @Tags Rule
// @Summary 创建规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shared.Rule true "创建规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rule/createRule [post]
func (ruleApi *RuleApi) CreateRule(c *gin.Context) {
	var rule shared.Rule
	err := c.ShouldBindJSON(&rule)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rule.CreatedBy = utils.GetUserID(c)
	// 接口创建的永远不是系统规则
	rule.SystemDefault = nil
	if err := ruleService.CreateRule(&rule); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRule 删除规则
// @Tags Rule
// @Summary 删除规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shared.Rule true "删除规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rule/deleteRule [delete]
func (ruleApi *RuleApi) DeleteRule(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := ruleService.DeleteRule(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRuleByIds 批量删除规则
// @Tags Rule
// @Summary 批量删除规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rule/deleteRuleByIds [delete]
func (ruleApi *RuleApi) DeleteRuleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := ruleService.DeleteRuleByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRule 更新规则
// @Tags Rule
// @Summary 更新规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shared.Rule true "更新规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rule/updateRule [put]
func (ruleApi *RuleApi) UpdateRule(c *gin.Context) {
	var rule shared.Rule
	err := c.ShouldBindJSON(&rule)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rule.UpdatedBy = utils.GetUserID(c)

	if err := ruleService.UpdateRule(rule); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRule 用id查询规则
// @Tags Rule
// @Summary 用id查询规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shared.Rule true "用id查询规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rule/findRule [get]
func (ruleApi *RuleApi) FindRule(c *gin.Context) {
	ID := c.Query("ID")
	if rerule, err := ruleService.GetRule(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerule": rerule}, c)
	}
}

// GetRuleList 分页获取规则列表
// @Tags Rule
// @Summary 分页获取规则列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sharedReq.RuleSearch true "分页获取规则列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rule/getRuleList [get]
func (ruleApi *RuleApi) GetRuleList(c *gin.Context) {
	var pageInfo sharedReq.RuleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ruleService.GetRuleInfoList(pageInfo); err != nil {
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
