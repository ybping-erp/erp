package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
	"gorm.io/gorm"
)

type InboundLogService struct {
}

// CreateInboundLog 创建入库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (inboundLogService *InboundLogService) CreateInboundLog(inboundLog *wms.InboundLog) (err error) {
	err = global.GVA_DB.Create(inboundLog).Error
	return err
}

// DeleteInboundLog 删除入库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (inboundLogService *InboundLogService) DeleteInboundLog(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wms.InboundLog{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}

		inboundLog, err := inboundLogService.GetInboundLog(ID)
		if err != nil {
			return err
		}

		if err = tx.Delete(&inboundLog).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// DeleteInboundLogByIds 批量删除入库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (inboundLogService *InboundLogService) DeleteInboundLogByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wms.InboundLog{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&wms.InboundLog{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateInboundLog 更新入库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (inboundLogService *InboundLogService) UpdateInboundLog(inboundLog wms.InboundLog) (err error) {
	err = global.GVA_DB.Save(&inboundLog).Error
	return err
}

// GetInboundLog 根据ID获取入库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (inboundLogService *InboundLogService) GetInboundLog(ID string) (inboundLog wms.InboundLog, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&inboundLog).Error
	return
}

// GetInboundLogInfoList 分页获取入库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (inboundLogService *InboundLogService) GetInboundLogInfoList(info wmsReq.InboundLogSearch) (list []wms.InboundLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wms.InboundLog{})
	var inboundLogs []wms.InboundLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GoodsSku != "" {
		db = db.Where("goods_sku = ?", info.GoodsSku)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&inboundLogs).Error
	return inboundLogs, total, err
}
