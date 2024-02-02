// 自动生成模板Rule
package shared

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 规则 结构体  Rule
type Rule struct {
	global.GVA_MODEL
	ActionCode    *int   `json:"actionCode" form:"actionCode" gorm:"column:action_code;comment:动作代码;size:10;"`            //动作代码
	ActionDesc    string `json:"actionDesc" form:"actionDesc" gorm:"column:action_desc;comment:动作描述;size:255;"`           //动作描述
	Domain        string `json:"domain" form:"domain" gorm:"column:domain;comment:领域;size:128;"`                          //领域
	Name          string `json:"name" form:"name" gorm:"column:name;comment:规则名称;size:255;"`                              //规则名称
	RuleStr       string `json:"ruleStr" form:"ruleStr" gorm:"column:rule_str;comment:规则条件;"`                             //规则条件
	Status        *int   `json:"status" form:"status" gorm:"column:status;comment:规则状态;size:10;"`                         //规则状态
	SystemDefault *int   `json:"systemDefault" form:"systemDefault" gorm:"column:system_default;comment:系统默认规则;size:10;"` //系统默认规则
	CreatedBy     uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 规则 Rule自定义表名 t_rules
func (Rule) TableName() string {
	return "t_rules"
}
