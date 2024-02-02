package shared

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shared"
    sharedReq "github.com/flipped-aurora/gin-vue-admin/server/model/shared/request"
    "gorm.io/gorm"
)

type RuleService struct {
}

// CreateRule 创建规则记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleService *RuleService) CreateRule(rule *shared.Rule) (err error) {
	err = global.GVA_DB.Create(rule).Error
	return err
}

// DeleteRule 删除规则记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleService *RuleService)DeleteRule(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&shared.Rule{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&shared.Rule{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteRuleByIds 批量删除规则记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleService *RuleService)DeleteRuleByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&shared.Rule{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&shared.Rule{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateRule 更新规则记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleService *RuleService)UpdateRule(rule shared.Rule) (err error) {
	err = global.GVA_DB.Save(&rule).Error
	return err
}

// GetRule 根据ID获取规则记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleService *RuleService)GetRule(ID string) (rule shared.Rule, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&rule).Error
	return
}

// GetRuleInfoList 分页获取规则记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleService *RuleService)GetRuleInfoList(info sharedReq.RuleSearch) (list []shared.Rule, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&shared.Rule{})
    var rules []shared.Rule
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&rules).Error
	return  rules, total, err
}
