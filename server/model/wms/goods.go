// 自动生成模板Goods
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 商品 结构体  Goods
type Goods struct {
      global.GVA_MODEL
      Code  string `json:"code" form:"code" gorm:"column:code;comment:商品编码;size:20;"`  //商品编码 
      Spu  string `json:"spu" form:"spu" gorm:"column:spu;comment:商品SPU;size:50;"`  //商品SPU 
      SpuAttributes  string `json:"spuAttributes" form:"spuAttributes" gorm:"column:spu_attributes;comment:属性信息;"`  //属性信息 
      Sku  string `json:"sku" form:"sku" gorm:"column:sku;comment:商品SKU;size:50;"`  //商品SKU 
      ChildrenSku  string `json:"childrenSku" form:"childrenSku" gorm:"column:children_sku;comment:组合商品;"`  //组合商品 
      NeedAdditionalProcess  *int `json:"needAdditionalProcess" form:"needAdditionalProcess" gorm:"column:need_additional_process;comment:需要加工;size:10;"`  //需要加工 
      SalesMethod  *int `json:"salesMethod" form:"salesMethod" gorm:"column:sales_method;comment:销售方式;"`  //销售方式 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:商品状态;"`  //商品状态 
      CategoryId  *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:商品分类;size:10;"`  //商品分类 
      ChineseName  string `json:"chineseName" form:"chineseName" gorm:"column:chinese_name;comment:中文名称;size:1024;"`  //中文名称 
      EnglishName  string `json:"englishName" form:"englishName" gorm:"column:english_name;comment:英文名称;size:1024;"`  //英文名称 
      IdentifierCode  string `json:"identifierCode" form:"identifierCode" gorm:"column:identifier_code;comment:识别码;size:50;"`  //识别码 
      SourceUrls  string `json:"sourceUrls" form:"sourceUrls" gorm:"column:source_urls;comment:来源URL列表;"`  //来源URL列表 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
      
      SkuMappings []SkuMapping `json:"sku_mappings" gorm:"foreignKey:GoodsSku;references:Sku"`
}


// TableName 商品 Goods自定义表名 t_wms_goods
func (Goods) TableName() string {
  return "t_wms_goods"
}

