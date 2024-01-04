// 自动生成模板Warehouse
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 仓库 结构体  Warehouse
type Warehouse struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:仓库名称;size:255;"`  //仓库名称 
      Owner  string `json:"owner" form:"owner" gorm:"column:owner;comment:负责人;size:255;"`  //负责人 
      Telephone  string `json:"telephone" form:"telephone" gorm:"column:telephone;comment:联系电话;size:255;"`  //联系电话 
      StreetAddress  string `json:"streetAddress" form:"streetAddress" gorm:"column:street_address;comment:仓库街道地址;size:255;"`  //仓库街道地址 
      City  string `json:"city" form:"city" gorm:"column:city;comment:城市;size:100;"`  //城市 
      StateProvince  string `json:"stateProvince" form:"stateProvince" gorm:"column:state_province;comment:州/省;size:100;"`  //州/省 
      PostalCode  string `json:"postalCode" form:"postalCode" gorm:"column:postal_code;comment:邮政编码;size:20;"`  //邮政编码 
      Country  string `json:"country" form:"country" gorm:"column:country;comment:国家;size:100;"`  //国家 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 仓库 Warehouse自定义表名 t_wms_warehouse
func (Warehouse) TableName() string {
  return "t_wms_warehouse"
}

