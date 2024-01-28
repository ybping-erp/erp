// 自动生成模板Goods
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type GoodsSPU struct {
	global.GVA_MODEL

	SkuList []Goods `json:"sku_list" form:"sku_list" gorm:"-"` // sku 列表

	CreatedBy uint `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 商品SPU t_wms_goods_spu
func (GoodsSPU) TableName() string {
	return "t_wms_goods_spu"
}

// 商品 结构体  Goods
type Goods struct {
	global.GVA_MODEL
	Code           string `json:"code" form:"code" gorm:"column:code;comment:商品编码;size:20;"`                               //商品编码
	SpuID          uint   `json:"spuId" form:"spuId" gorm:"column:spu_id;comment:商品SPU;"`                                  //商品SPU
	Sku            string `json:"sku" form:"sku" gorm:"column:sku;comment:商品SKU;size:50;"`                                 //商品SKU
	SalesMethod    *int   `json:"salesMethod" form:"salesMethod" gorm:"column:sales_method;comment:销售方式;"`                 //销售方式
	Status         *int   `json:"status" form:"status" gorm:"column:status;comment:商品状态;"`                                 //商品状态
	CategoryId     *int   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:商品分类;size:10;"`            //商品分类
	ChineseName    string `json:"chineseName" form:"chineseName" gorm:"column:chinese_name;comment:中文名称;size:1024;"`       //中文名称
	EnglishName    string `json:"englishName" form:"englishName" gorm:"column:english_name;comment:英文名称;size:1024;"`       //英文名称
	IdentifierCode string `json:"identifierCode" form:"identifierCode" gorm:"column:identifier_code;comment:识别码;size:50;"` //识别码
	SourceUrls     string `json:"sourceUrls" form:"sourceUrls" gorm:"column:source_urls;comment:来源URL列表;"`                 //来源URL列表
	CreatedBy      uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint   `gorm:"column:deleted_by;comment:删除者"`

	Children    []Goods          `json:"children" gorm:"-"`                     // 组合商品
	Attributes  []GoodsAttribute `json:"attributes" form:"attributes" gorm:"-"` //属性信息
	SkuMappings []SkuMapping     `json:"sku_mappings" gorm:"foreignKey:GoodsSku;references:Sku"`
}

// TableName 商品 Goods自定义表名 t_wms_goods
func (Goods) TableName() string {
	return "t_wms_goods"
}

type GoodsAttribute struct {
	global.GVA_MODEL
	GoodsID   uint   `gorm:"column:goods_id;comment:商品ID"`
	Key       string `json:"key" form:"key" gorm:"column:key;comment:属性名称;size:20;"`      //商品编码
	Value     string `json:"value" form:"value" gorm:"column:value;comment:属性值;size:20;"` //商品编码
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 商品属性 GoodsAttribute t_wms_goods_attribute
func (GoodsAttribute) TableName() string {
	return "t_wms_goods_attribute"
}
