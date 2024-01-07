package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
)

type OrderService struct {
}

// CreateOrder 创建订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) CreateOrder(order *eCommerce.Order) (err error) {
	err = global.GVA_DB.Create(order).Error
	return err
}

// DeleteOrder 删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) DeleteOrder(order eCommerce.Order) (err error) {
	err = global.GVA_DB.Delete(&order).Error
	return err
}

// DeleteOrderByIds 批量删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]eCommerce.Order{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrder 更新订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) UpdateOrder(order eCommerce.Order) (err error) {
	err = global.GVA_DB.Save(&order).Error
	return err
}

// GetOrder 根据id获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) GetOrder(id uint) (order eCommerce.Order, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&order).Error
	return
}

// GetOrderInfoList 分页获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderService *OrderService) GetOrderInfoList(info eCommerceReq.OrderSearch) (list []eCommerce.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&eCommerce.Order{})
	var orders []eCommerce.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrderId != "" {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.PlatformName != "" {
		db = db.Where("platform_name = ?", info.PlatformName)
	}
	if info.ShopId != "" {
		db = db.Where("shop_id = ?", info.ShopId)
	}
	if info.ShopName != "" {
		db = db.Where("shop_name LIKE ?", "%"+info.ShopName+"%")
	}
	if info.StatusId != nil {
		db = db.Where("status_id = ?", info.StatusId)
	}
	if info.ShippingOrderId != "" {
		db = db.Where("shipping_order_id = ?", info.ShippingOrderId)
	}
	if info.ShippingCountry != "" {
		db = db.Where("shipping_country = ?", info.ShippingCountry)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&orders).Error
	return orders, total, err
}
