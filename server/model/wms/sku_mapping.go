// 自动生成模板SkuMapping
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 商品映射 结构体  SkuMapping
type SkuMapping struct {
      global.GVA_MODEL
      GoodsSku  string `json:"goodsSku" form:"goodsSku" gorm:"column:goods_sku;comment:商品SKU;size:50;"`  //商品SKU 
      ProductSku  string `json:"productSku" form:"productSku" gorm:"column:product_sku;comment:产品SKU;size:50;"`  //产品SKU 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 商品映射 SkuMapping自定义表名 t_wms_sku_mapping
func (SkuMapping) TableName() string {
  return "t_wms_sku_mapping"
}

