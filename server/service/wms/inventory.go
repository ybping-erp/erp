package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
)

type InventoryService struct {
}

// CreateInventory 创建商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) CreateInventory(inventory *wms.Inventory) (err error) {
	err = global.GVA_DB.Create(inventory).Error
	return err
}

// DeleteInventory 删除商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService)DeleteInventory(inventory wms.Inventory) (err error) {
	err = global.GVA_DB.Delete(&inventory).Error
	return err
}

// DeleteInventoryByIds 批量删除商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService)DeleteInventoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wms.Inventory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateInventory 更新商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService)UpdateInventory(inventory wms.Inventory) (err error) {
	err = global.GVA_DB.Save(&inventory).Error
	return err
}

// GetInventory 根据id获取商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService)GetInventory(id uint) (inventory wms.Inventory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&inventory).Error
	return
}

// GetInventoryInfoList 分页获取商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService)GetInventoryInfoList(info wmsReq.InventorySearch) (list []wms.Inventory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.Inventory{})
    var inventorys []wms.Inventory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.GoodsSku != "" {
        db = db.Where("goods_sku = ?",info.GoodsSku)
    }
    if info.WarehouseId != nil {
        db = db.Where("warehouse_id = ?",info.WarehouseId)
    }
    if info.StockStatus != nil {
        db = db.Where("stock_status = ?",info.StockStatus)
    }
	if info.RackId != nil {
		db = db.Where("rack_id = ?",info.RackId)
	}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&inventorys).Error
	return  inventorys, total, err
}
