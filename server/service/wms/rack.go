package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "gorm.io/gorm"
)

type RackService struct {
}

// CreateRack 创建货架记录
// Author [piexlmax](https://github.com/piexlmax)
func (rackService *RackService) CreateRack(rack *wms.Rack) (err error) {
	err = global.GVA_DB.Create(rack).Error
	return err
}

// DeleteRack 删除货架记录
// Author [piexlmax](https://github.com/piexlmax)
func (rackService *RackService)DeleteRack(rack wms.Rack) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.Rack{}).Where("id = ?", rack.ID).Update("deleted_by", rack.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&rack).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteRackByIds 批量删除货架记录
// Author [piexlmax](https://github.com/piexlmax)
func (rackService *RackService)DeleteRackByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.Rack{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&wms.Rack{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateRack 更新货架记录
// Author [piexlmax](https://github.com/piexlmax)
func (rackService *RackService)UpdateRack(rack wms.Rack) (err error) {
	err = global.GVA_DB.Save(&rack).Error
	return err
}

// GetRack 根据id获取货架记录
// Author [piexlmax](https://github.com/piexlmax)
func (rackService *RackService)GetRack(id uint) (rack wms.Rack, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&rack).Error
	return
}

// GetRackInfoList 分页获取货架记录
// Author [piexlmax](https://github.com/piexlmax)
func (rackService *RackService)GetRackInfoList(info wmsReq.RackSearch) (list []wms.Rack, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.Rack{})
    var racks []wms.Rack
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.WarehouseId != nil {
        db = db.Where("warehouse_id = ?",info.WarehouseId)
    }
    if info.ZoneId != nil {
        db = db.Where("zone_id = ?",info.ZoneId)
    }
    if info.RackCode != "" {
        db = db.Where("rack_code = ?",info.RackCode)
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
	
	err = db.Find(&racks).Error
	return  racks, total, err
}
