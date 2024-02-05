// 自动生成模板Inventory
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 商品库存 结构体  Inventory
type Inventory struct {
	global.GVA_MODEL
	GoodsSku           string `json:"goodsSku" form:"goodsSku" gorm:"column:goods_sku;comment:商品SKU;size:50;"`                                //商品SKU
	WarehouseId        *int   `json:"warehouseId" form:"warehouseId" gorm:"column:warehouse_id;comment:关联仓库的标识符;size:10;"`                    //关联仓库的标识符
	ZoneId             *int   `json:"zoneId" form:"zoneId" gorm:"column:zone_id;comment:库区ID;size:10;"`                                       //库区ID
	RackId             *int   `json:"rackId" form:"rackId" gorm:"column:rack_id;comment:货架ID;size:10;"`                                       //货架ID
	Quantity           *int   `json:"quantity" form:"quantity" gorm:"column:quantity;comment:库存数量;size:10;"`                                  //库存数量
	ReservedQuantity   *int   `json:"reservedQuantity" form:"reservedQuantity" gorm:"column:reserved_quantity;comment:预留库存数量;size:10;"`       //预留库存数量
	OnOrderQuantity    *int   `json:"onOrderQuantity" form:"onOrderQuantity" gorm:"column:on_order_quantity;comment:在途库存数量;size:10;"`         //在途库存数量
	SafetyQuantity     *int   `json:"safetyQuantity" form:"safetyQuantity" gorm:"column:safety_quantity;comment:预留库存数量;size:10;"`             //预留库存数量
	UnshippingQuantity *int   `json:"unshippingQuantity" form:"unshippingQuantity" gorm:"column:unshipping_quantity;comment:在途库存数量;size:10;"` //未发库存数量
	PreSaleQuantity    *int   `json:"preSaleQuantity" form:"preSaleQuantity" gorm:"column:pre_sale_quantity;comment:在途库存数量;size:10;"`         //未发库存数量
	StockStatus        *int   `json:"stockStatus" form:"stockStatus" gorm:"column:stock_status;comment:库存状态;"`                                //库存状态
}

// TableName 商品库存 Inventory自定义表名 t_wms_inventory
func (Inventory) TableName() string {
	return "t_wms_inventory"
}
