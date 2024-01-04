package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "gorm.io/gorm"
)

type WarehouseService struct {
}

// CreateWarehouse 创建仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseService) CreateWarehouse(warehouse *wms.Warehouse) (err error) {
	err = global.GVA_DB.Create(warehouse).Error
	return err
}

// DeleteWarehouse 删除仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseService)DeleteWarehouse(warehouse wms.Warehouse) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.Warehouse{}).Where("id = ?", warehouse.ID).Update("deleted_by", warehouse.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&warehouse).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteWarehouseByIds 批量删除仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseService)DeleteWarehouseByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.Warehouse{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&wms.Warehouse{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateWarehouse 更新仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseService)UpdateWarehouse(warehouse wms.Warehouse) (err error) {
	err = global.GVA_DB.Save(&warehouse).Error
	return err
}

// GetWarehouse 根据id获取仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseService)GetWarehouse(id uint) (warehouse wms.Warehouse, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&warehouse).Error
	return
}

// GetWarehouseInfoList 分页获取仓库记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseService)GetWarehouseInfoList(info wmsReq.WarehouseSearch) (list []wms.Warehouse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.Warehouse{})
    var warehouses []wms.Warehouse
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.City != "" {
        db = db.Where("city LIKE ?","%"+ info.City+"%")
    }
    if info.StateProvince != "" {
        db = db.Where("state_province LIKE ?","%"+ info.StateProvince+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&warehouses).Error
	return  warehouses, total, err
}
