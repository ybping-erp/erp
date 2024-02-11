package wms

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
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
func (pickOrderService *PickOrderService) DeletePickOrder(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wms.PickOrder{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&wms.PickOrder{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePickOrderByIds 批量删除拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService) DeletePickOrderByIds(IDs []string, deleted_by uint) (err error) {
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
func (pickOrderService *PickOrderService) UpdatePickOrder(pickOrder wms.PickOrder) (err error) {
	err = global.GVA_DB.Save(&pickOrder).Error
	return err
}

// GetPickOrder 根据ID获取拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService) GetPickOrder(ID string) (pickOrder wms.PickOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pickOrder).Error
	return
}

// GetPickOrderInfoList 分页获取拣货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (pickOrderService *PickOrderService) GetPickOrderInfoList(info wmsReq.PickOrderSearch) (list []wms.PickOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wms.PickOrder{})
	var pickOrders []wms.PickOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GoodsName != "" {
		db = db.Where("goods_name LIKE ?", "%"+info.GoodsName+"%")
	}
	if info.GoodsSku != "" {
		db = db.Where("goods_sku = ?", info.GoodsSku)
	}
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&pickOrders).Error
	return pickOrders, total, err
}

func (pickOrderService *PickOrderService) AllocatePickOrders(order eCommerce.Order, inventoryMap map[string][]wms.Inventory) (pickOrders []wms.PickOrder, err error) {
	// 根据库存为每个订单明细分配拣货单
	for _, item := range order.Items {
		if item.GoodsSku == nil {
			err = fmt.Errorf("%v 不存在配对的商品", item.ProductSku)
			return
		}
		goodsSKU := *item.GoodsSku
		inventoryList, exists := inventoryMap[goodsSKU]
		if !exists {
			err = fmt.Errorf("%v 库存不存在", goodsSKU)
			return
		}

		for _, inventory := range inventoryList {
			// 说明已经全部分配成功
			if *item.Quantity == 0 {
				break
			}

			// 开始分配
			quantity := utils.MinInt(*item.Quantity, *inventory.Quantity)
			*inventory.Quantity -= quantity
			*item.Quantity -= quantity

			pickOrder := wms.PickOrder{
				// ybping: 修改为GoodsName
				GoodsName:   goodsSKU,
				GoodsSku:    goodsSKU,
				OrderId:     &item.OrderId,
				Quantity:    &quantity,
				RackId:      inventory.RackId,
				WarehouseId: inventory.WarehouseId,
				ZoneId:      inventory.ZoneId,
			}
			pickOrders = append(pickOrders, pickOrder)
		}

		// 还有剩余的数量未分配
		if *item.Quantity > 0 {
			err = fmt.Errorf("%v 库存不足", item.ProductSku)
			return
		}
	}

	// 更新数据库状态
	inventorys := utils.MapToArray(inventoryMap)
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(inventorys).Error; err != nil {
			return err
		}
		if err := tx.Save(pickOrders).Error; err != nil {
			return err
		}

		return nil
	})
	return
}
