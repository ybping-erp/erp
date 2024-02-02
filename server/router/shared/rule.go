package shared

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RuleRouter struct {
}

// InitRuleRouter 初始化 规则 路由信息
func (s *RuleRouter) InitRuleRouter(Router *gin.RouterGroup) {
	ruleRouter := Router.Group("rule").Use(middleware.OperationRecord())
	ruleRouterWithoutRecord := Router.Group("rule")
	var ruleApi = v1.ApiGroupApp.SharedApiGroup.RuleApi
	{
		ruleRouter.POST("createRule", ruleApi.CreateRule)   // 新建规则
		ruleRouter.DELETE("deleteRule", ruleApi.DeleteRule) // 删除规则
		ruleRouter.DELETE("deleteRuleByIds", ruleApi.DeleteRuleByIds) // 批量删除规则
		ruleRouter.PUT("updateRule", ruleApi.UpdateRule)    // 更新规则
	}
	{
		ruleRouterWithoutRecord.GET("findRule", ruleApi.FindRule)        // 根据ID获取规则
		ruleRouterWithoutRecord.GET("getRuleList", ruleApi.GetRuleList)  // 获取规则列表
	}
}
