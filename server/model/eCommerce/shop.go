// 自动生成模板Shop
package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 店铺 结构体  Shop
type Shop struct {
	global.GVA_MODEL
	PlatformName string `json:"platformName" form:"platformName" gorm:"column:platform_name;comment:平台名称;size:255;"`      //平台名称
	ShopId       string `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:平台店铺唯一标识符;size:255;"`                   //平台ID
	ShopName     string `json:"shopName" form:"shopName" gorm:"column:shop_name;comment:店铺名称;size:255;"`                  //店铺名称
	ClientId     string `json:"clientId" form:"clientId" gorm:"column:client_id;comment:API Client ID;size:255;"`         //API Client ID
	ClientCert   string `json:"clientCert" form:"clientCert" gorm:"column:client_cert;comment:API Client Cert;size:255;"` //API Client Cert
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 店铺 Shop自定义表名 t_shop
func (Shop) TableName() string {
	return "t_shop"
}
