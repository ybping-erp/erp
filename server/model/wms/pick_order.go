// 自动生成模板PickOrder
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 拣货单 结构体  PickOrder
type PickOrder struct {
	global.GVA_MODEL
	GoodsName   string `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;size:50;"`       //商品
	GoodsSku    string `json:"goodsSku" form:"goodsSku" gorm:"column:goods_sku;comment:商品SKU;size:50;"`         //SKU
	OrderId     *int   `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;size:20;"`             //订单
	Quantity    *int   `json:"quantity" form:"quantity" gorm:"column:quantity;comment:拣货数量;size:10;"`           //拣货数量
	RackId      *int   `json:"rackId" form:"rackId" gorm:"column:rack_id;comment:货架ID;size:20;"`                //货架
	WarehouseId *int   `json:"warehouseId" form:"warehouseId" gorm:"column:warehouse_id;comment:仓库ID;size:20;"` //仓库
	ZoneId      *int   `json:"zoneId" form:"zoneId" gorm:"column:zone_id;comment:库区ID;size:20;"`                //库区
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 拣货单 PickOrder自定义表名 t_wms_pick_order
func (PickOrder) TableName() string {
	return "t_wms_pick_order"
}
