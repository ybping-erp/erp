package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "gorm.io/gorm"
)

type SkuMappingService struct {
}

// CreateSkuMapping 创建商品映射记录
// Author [piexlmax](https://github.com/piexlmax)
func (skuMappingService *SkuMappingService) CreateSkuMapping(skuMapping *wms.SkuMapping) (err error) {
	err = global.GVA_DB.Create(skuMapping).Error
	return err
}

// DeleteSkuMapping 删除商品映射记录
// Author [piexlmax](https://github.com/piexlmax)
func (skuMappingService *SkuMappingService)DeleteSkuMapping(skuMapping wms.SkuMapping) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.SkuMapping{}).Where("id = ?", skuMapping.ID).Update("deleted_by", skuMapping.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&skuMapping).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteSkuMappingByIds 批量删除商品映射记录
// Author [piexlmax](https://github.com/piexlmax)
func (skuMappingService *SkuMappingService)DeleteSkuMappingByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.SkuMapping{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&wms.SkuMapping{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateSkuMapping 更新商品映射记录
// Author [piexlmax](https://github.com/piexlmax)
func (skuMappingService *SkuMappingService)UpdateSkuMapping(skuMapping wms.SkuMapping) (err error) {
	err = global.GVA_DB.Save(&skuMapping).Error
	return err
}

// GetSkuMapping 根据id获取商品映射记录
// Author [piexlmax](https://github.com/piexlmax)
func (skuMappingService *SkuMappingService)GetSkuMapping(id uint) (skuMapping wms.SkuMapping, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&skuMapping).Error
	return
}

// GetSkuMappingInfoList 分页获取商品映射记录
// Author [piexlmax](https://github.com/piexlmax)
func (skuMappingService *SkuMappingService)GetSkuMappingInfoList(info wmsReq.SkuMappingSearch) (list []wms.SkuMapping, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.SkuMapping{})
    var skuMappings []wms.SkuMapping
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&skuMappings).Error
	return  skuMappings, total, err
}
