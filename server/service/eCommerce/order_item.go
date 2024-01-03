package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
)

type OrderItemService struct {
}

// CreateOrderItem 创建订单明细记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderItemService *OrderItemService) CreateOrderItem(orderItem *eCommerce.OrderItem) (err error) {
	err = global.GVA_DB.Create(orderItem).Error
	return err
}

// DeleteOrderItem 删除订单明细记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderItemService *OrderItemService)DeleteOrderItem(orderItem eCommerce.OrderItem) (err error) {
	err = global.GVA_DB.Delete(&orderItem).Error
	return err
}

// DeleteOrderItemByIds 批量删除订单明细记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderItemService *OrderItemService)DeleteOrderItemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]eCommerce.OrderItem{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOrderItem 更新订单明细记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderItemService *OrderItemService)UpdateOrderItem(orderItem eCommerce.OrderItem) (err error) {
	err = global.GVA_DB.Save(&orderItem).Error
	return err
}

// GetOrderItem 根据id获取订单明细记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderItemService *OrderItemService)GetOrderItem(id uint) (orderItem eCommerce.OrderItem, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&orderItem).Error
	return
}

// GetOrderItemInfoList 分页获取订单明细记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderItemService *OrderItemService)GetOrderItemInfoList(info eCommerceReq.OrderItemSearch) (list []eCommerce.OrderItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&eCommerce.OrderItem{})
    var orderItems []eCommerce.OrderItem
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
	
	err = db.Find(&orderItems).Error
	return  orderItems, total, err
}
