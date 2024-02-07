package wms

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
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
func (inventoryService *InventoryService) DeleteInventory(inventory wms.Inventory) (err error) {
	err = global.GVA_DB.Delete(&inventory).Error
	return err
}

// DeleteInventoryByIds 批量删除商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) DeleteInventoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wms.Inventory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateInventory 更新商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) UpdateInventory(inventory wms.Inventory) (err error) {
	err = global.GVA_DB.Save(&inventory).Error
	return err
}

// GetInventory 根据id获取商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) GetInventory(id uint) (inventory wms.Inventory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&inventory).Error
	return
}

// GetInventoryInfoList 分页获取商品库存记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventoryService *InventoryService) GetInventoryInfoList(info wmsReq.InventorySearch) (list []wms.Inventory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wms.Inventory{})
	var inventorys []wms.Inventory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GoodsSku != "" {
		db = db.Where("goods_sku = ?", info.GoodsSku)
	}
	if info.WarehouseId != nil {
		db = db.Where("warehouse_id = ?", info.WarehouseId)
	}
	if info.StockStatus != nil {
		db = db.Where("stock_status = ?", info.StockStatus)
	}
	if info.RackId != nil {
		db = db.Where("rack_id = ?", info.RackId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&inventorys).Error
	return inventorys, total, err
}

func (inventoryService *InventoryService) AddInventoryFromInboundLog(inboundLog wms.InboundLog) (err error) {
	if *inboundLog.Status == InboundLogStatus_FINISHED {
		return errors.New("该入库单已经入库")
	}

	if *inboundLog.Status != InboundLogStatus_PENDING {
		return errors.New("该状态不允许入库")
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var inventory wms.Inventory
		err := tx.Model(&wms.Inventory{}).Where("goods_sku = ? and warehouse_id = ? and zone_id = ? and rack_id = ?", inboundLog.GoodsSku, inboundLog.WarehouseId, inboundLog.ZoneId, inboundLog.RackId).First(&inventory).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 初始化库存
		if errors.Is(err, gorm.ErrRecordNotFound) {
			inventory = wms.Inventory{
				GoodsSku:    inboundLog.GoodsSku,
				WarehouseId: inboundLog.WarehouseId,
				ZoneId:      inboundLog.ZoneId,
				RackId:      inboundLog.RackId,
				Quantity:    utils.Pointer[int](0),
			}
		}

		// 更新入库单状态
		inboundLog.Status = utils.Pointer[int](InboundLogStatus_FINISHED)
		err = tx.Save(&inboundLog).Error
		if err != nil {
			return err
		}

		// 更新库存
		*inventory.Quantity += *inboundLog.Quantity
		return tx.Save(&inventory).Error
	})
	return err

}

func (inventoryService *InventoryService) SubInventoryFromOutboundLog(outBoundLog wms.OutboundLog) (err error) {
	if *outBoundLog.Status == OutboundLogStatus_FINISHED {
		return errors.New("该出库单已经出库")
	}

	if *outBoundLog.Status != InboundLogStatus_PENDING {
		return errors.New("该状态不允许出库")
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var inventory wms.Inventory
		err := tx.Model(&wms.Inventory{}).Where("goods_sku = ? and warehouse_id = ? and zone_id = ? and rack_id = ?", outBoundLog.GoodsSku, outBoundLog.WarehouseId, outBoundLog.ZoneId, outBoundLog.RackId).First(&inventory).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 初始化库存
		if errors.Is(err, gorm.ErrRecordNotFound) {
			inventory = wms.Inventory{
				GoodsSku:    outBoundLog.GoodsSku,
				WarehouseId: outBoundLog.WarehouseId,
				ZoneId:      outBoundLog.ZoneId,
				RackId:      outBoundLog.RackId,
				Quantity:    utils.Pointer[int](0),
			}
		}

		// 更新入库单状态
		outBoundLog.Status = utils.Pointer[int](OutboundLogStatus_FINISHED)
		err = tx.Save(&outBoundLog).Error
		if err != nil {
			return err
		}

		// 更新库存
		*inventory.Quantity -= *outBoundLog.Quantity
		return tx.Save(&inventory).Error
	})
	return err

}
