// 自动生成模板Product
package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
  "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 产品表 结构体  Product
type Product struct {
      global.GVA_MODEL
      CategoryId  *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:产品所属的类别标识符;size:10;"`  //产品所属的类别标识符 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:产品描述;"`  //产品描述 
      ProductName  string `json:"productName" form:"productName" gorm:"column:product_name;comment:产品名称;size:255;"`  //产品名称 
      Sku  string `json:"sku" form:"sku" gorm:"column:sku;comment:产品的唯一标识符;size:50;"`  //产品的唯一标识符 
      Summary  string `json:"summary" form:"summary" gorm:"column:summary;comment:产品摘要;"`  //产品摘要 
      UnitPrice  *float64 `json:"unitPrice" form:"unitPrice" gorm:"column:unit_price;comment:产品的单价;size:10;"`  //产品的单价 
      CreatorID       int           `json:"creator_id" form:"creator_id" gorm:"column:creator_id;comment:创建人"`  // 用户id
	    Creator         system.SysUser `json:"creator" gorm:"foreignKey:CreatorID;"`
}


// TableName 产品表 Product自定义表名 t_product
func (Product) TableName() string {
  return "t_product"
}

