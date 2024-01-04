// 自动生成模板Rack
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 货架 结构体  Rack
type Rack struct {
      global.GVA_MODEL
      WarehouseId  *int `json:"warehouseId" form:"warehouseId" gorm:"column:warehouse_id;comment:仓库ID;size:20;"`  //仓库ID 
      ZoneId  *int `json:"zoneId" form:"zoneId" gorm:"column:zone_id;comment:分区ID;size:20;"`  //分区ID 
      RackCode  string `json:"rackCode" form:"rackCode" gorm:"column:rack_code;comment:货架编号;size:50;"`  //货架编号 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:货架状态;"`  //货架状态 
      Priority  *int `json:"priority" form:"priority" gorm:"column:priority;comment:拣货权重;size:10;"`  //拣货权重 
      Remarks  string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注说明;"`  //备注说明 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 货架 Rack自定义表名 t_wms_rack
func (Rack) TableName() string {
  return "t_wms_rack"
}

