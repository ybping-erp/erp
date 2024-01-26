// 自动生成模板InboundLog
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 入库记录 结构体  InboundLog
type InboundLog struct {
 global.GVA_MODEL 
      GoodsSku  string `json:"goodsSku" form:"goodsSku" gorm:"column:goods_sku;comment:商品SKU;size:50;"binding:"required"`  //商品SKU 
      Quantity  *int `json:"quantity" form:"quantity" gorm:"column:quantity;comment:入库数量;size:10;"binding:"required"`  //入库数量 
      RackId  *int `json:"rackId" form:"rackId" gorm:"column:rack_id;comment:货架ID;size:20;"binding:"required"`  //货架 
      WarehouseId  *int `json:"warehouseId" form:"warehouseId" gorm:"column:warehouse_id;comment:仓库ID;size:20;"binding:"required"`  //仓库 
      ZoneId  *int `json:"zoneId" form:"zoneId" gorm:"column:zone_id;comment:库区ID;size:20;"binding:"required"`  //库区 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 入库记录 InboundLog自定义表名 t_wms_inbound_log
func (InboundLog) TableName() string {
  return "t_wms_inbound_log"
}

