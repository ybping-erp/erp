// 自动生成模板OrderItem
package eCommerce

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	"gorm.io/gorm"
)

// 订单明细 结构体  OrderItem
type OrderItem struct {
	global.GVA_MODEL
	OrderId        int      `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;"`
	PlatormOrderId string   `json:"platformOrderId" form:"platformOrderId" gorm:"column:platform_order_id;comment:关联订单的标识符;size:50;"` //关联订单的标识符
	ProductSku     string   `json:"productSku" form:"productSku" gorm:"column:product_sku;comment:关联产品的标识符;size:50;"`                 //关联产品的标识符
	ProductUrl     string   `json:"productUrl" form:"productUrl" gorm:"column:product_url;comment:产品图片;size:512;"`                    //产品图片
	Attributes     string   `json:"attributes" form:"attributes" gorm:"column:attributes;comment:产品属性;"`                              //产品属性
	Quantity       *int     `json:"quantity" form:"quantity" gorm:"column:quantity;comment:订单中产品的数量;size:10;"`                        //订单中产品的数量
	UnitPrice      *float64 `json:"unitPrice" form:"unitPrice" gorm:"column:unit_price;comment:产品的单价;size:10;"`                       //产品的单价
	TotalAmount    *float64 `json:"totalAmount" form:"totalAmount" gorm:"column:total_amount;comment:订单项总金额;size:12;"`                //订单项总金额
	GoodsSku       *string  `json:"goodsSku" gorm:"-"`
}

// TableName 订单明细 OrderItem自定义表名 t_order_item
func (OrderItem) TableName() string {
	return "t_order_item"
}

func (item *OrderItem) AfterFind(tx *gorm.DB) (err error) {
	var mapping wms.SkuMapping
	err = tx.Model(&wms.SkuMapping{}).Where("product_sku = ?", item.ProductSku).First(&mapping).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	} else if err != nil {
		return
	} else {
		item.GoodsSku = &mapping.GoodsSku
	}

	return
}
