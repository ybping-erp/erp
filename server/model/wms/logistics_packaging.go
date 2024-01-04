// 自动生成模板LogisticsPackaging
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 物流包装 结构体  LogisticsPackaging
type LogisticsPackaging struct {
      global.GVA_MODEL
      GoodsSku  string `json:"goodsSku" form:"goodsSku" gorm:"column:goods_sku;comment:商品SKU;size:50;"`  //商品SKU 
      DeclarationChineseName  string `json:"declarationChineseName" form:"declarationChineseName" gorm:"column:declaration_chinese_name;comment:报关中文名;size:50;"`  //报关中文名 
      DeclarationEnglishName  string `json:"declarationEnglishName" form:"declarationEnglishName" gorm:"column:declaration_english_name;comment:报关英文名;size:50;"`  //报关英文名 
      DeclarationPriceCurrency  string `json:"declarationPriceCurrency" form:"declarationPriceCurrency" gorm:"column:declaration_price_currency;comment:报关价格币种;size:10;"`  //报关价格币种 
      DeclarationWeightUnit  string `json:"declarationWeightUnit" form:"declarationWeightUnit" gorm:"column:declaration_weight_unit;comment:报关重量单位;size:1;"`  //报关重量单位 
      Material  string `json:"material" form:"material" gorm:"column:material;comment:材质;size:255;"`  //材质 
      Purpose  string `json:"purpose" form:"purpose" gorm:"column:purpose;comment:用途;size:255;"`  //用途 
      CustomsCode  string `json:"customsCode" form:"customsCode" gorm:"column:customs_code;comment:海关编码;size:50;"`  //海关编码 
      SpecialGoods  *int `json:"specialGoods" form:"specialGoods" gorm:"column:special_goods;comment:危险运输品;"`  //危险运输品 
      NetWeight  *float64 `json:"netWeight" form:"netWeight" gorm:"column:net_weight;comment:商品净重 (g);size:10;"`  //商品净重 (g) 
      WeightErrorTolerance  *float64 `json:"weightErrorTolerance" form:"weightErrorTolerance" gorm:"column:weight_error_tolerance;comment:允许称重误差 (g);size:10;"`  //允许称重误差 (g) 
      LengthCm  *float64 `json:"lengthCm" form:"lengthCm" gorm:"column:length_cm;comment:商品尺寸 - 长 (cm);size:10;"`  //商品尺寸 - 长 (cm) 
      WidthCm  *float64 `json:"widthCm" form:"widthCm" gorm:"column:width_cm;comment:商品尺寸 - 宽 (cm);size:10;"`  //商品尺寸 - 宽 (cm) 
      HeightCm  *float64 `json:"heightCm" form:"heightCm" gorm:"column:height_cm;comment:商品尺寸 - 高 (cm);size:10;"`  //商品尺寸 - 高 (cm) 
      Packaging  string `json:"packaging" form:"packaging" gorm:"column:packaging;comment:包装信息;"`  //包装信息 
}


// TableName 物流包装 LogisticsPackaging自定义表名 t_wms_logistics_packaging
func (LogisticsPackaging) TableName() string {
  return "t_wms_logistics_packaging"
}

