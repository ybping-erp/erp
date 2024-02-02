package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "gorm.io/gorm"
)

type SupplierService struct {
}

// CreateSupplier 创建供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService) CreateSupplier(supplier *wms.Supplier) (err error) {
	err = global.GVA_DB.Create(supplier).Error
	return err
}

// DeleteSupplier 删除供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)DeleteSupplier(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.Supplier{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wms.Supplier{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteSupplierByIds 批量删除供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)DeleteSupplierByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.Supplier{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wms.Supplier{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateSupplier 更新供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)UpdateSupplier(supplier wms.Supplier) (err error) {
	err = global.GVA_DB.Save(&supplier).Error
	return err
}

// GetSupplier 根据ID获取供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)GetSupplier(ID string) (supplier wms.Supplier, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&supplier).Error
	return
}

// GetSupplierInfoList 分页获取供应商记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)GetSupplierInfoList(info wmsReq.SupplierSearch) (list []wms.Supplier, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.Supplier{})
    var suppliers []wms.Supplier
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.ContactName != "" {
        db = db.Where("contact_name LIKE ?","%"+ info.ContactName+"%")
    }
    if info.ContactTelephone != "" {
        db = db.Where("contact_telephone = ?",info.ContactTelephone)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&suppliers).Error
	return  suppliers, total, err
}
