package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "gorm.io/gorm"
)

type PickOrderService struct {
}

// CreatePickOrder 创建拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService) CreatePickOrder(pickOrder *wms.PickOrder) (err error) {
	err = global.GVA_DB.Create(pickOrder).Error
	return err
}

// DeletePickOrder 删除拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService)DeletePickOrder(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.PickOrder{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&wms.PickOrder{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeletePickOrderByIds 批量删除拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService)DeletePickOrderByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.PickOrder{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&wms.PickOrder{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdatePickOrder 更新拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService)UpdatePickOrder(pickOrder wms.PickOrder) (err error) {
	err = global.GVA_DB.Save(&pickOrder).Error
	return err
}

// GetPickOrder 根据ID获取拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService)GetPickOrder(ID string) (pickOrder wms.PickOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pickOrder).Error
	return
}

// GetPickOrderInfoList 分页获取拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService)GetPickOrderInfoList(info wmsReq.PickOrderSearch) (list []wms.PickOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.PickOrder{})
    var pickOrders []wms.PickOrder
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.GoodsName != "" {
        db = db.Where("goods_name LIKE ?","%"+ info.GoodsName+"%")
    }
    if info.GoodsSku != "" {
        db = db.Where("goods_sku = ?",info.GoodsSku)
    }
    if info.OrderId != nil {
        db = db.Where("order_id = ?",info.OrderId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&pickOrders).Error
	return  pickOrders, total, err
}
