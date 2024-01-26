package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "gorm.io/gorm"
)

type OutboundLogService struct {
}

// CreateOutboundLog 创建出库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (outboundLogService *OutboundLogService) CreateOutboundLog(outboundLog *wms.OutboundLog) (err error) {
	err = global.GVA_DB.Create(outboundLog).Error
	return err
}

// DeleteOutboundLog 删除出库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (outboundLogService *OutboundLogService)DeleteOutboundLog(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.OutboundLog{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wms.OutboundLog{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteOutboundLogByIds 批量删除出库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (outboundLogService *OutboundLogService)DeleteOutboundLogByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.OutboundLog{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wms.OutboundLog{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateOutboundLog 更新出库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (outboundLogService *OutboundLogService)UpdateOutboundLog(outboundLog wms.OutboundLog) (err error) {
	err = global.GVA_DB.Save(&outboundLog).Error
	return err
}

// GetOutboundLog 根据ID获取出库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (outboundLogService *OutboundLogService)GetOutboundLog(ID string) (outboundLog wms.OutboundLog, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&outboundLog).Error
	return
}

// GetOutboundLogInfoList 分页获取出库记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (outboundLogService *OutboundLogService)GetOutboundLogInfoList(info wmsReq.OutboundLogSearch) (list []wms.OutboundLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.OutboundLog{})
    var outboundLogs []wms.OutboundLog
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.GoodsSku != "" {
        db = db.Where("goods_sku = ?",info.GoodsSku)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&outboundLogs).Error
	return  outboundLogs, total, err
}
