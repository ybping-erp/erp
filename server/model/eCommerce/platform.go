// 自动生成模板Platform
package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 电商平台 结构体  Platform
type Platform struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:平台名称;size:100;"`  //平台名称 
      LogoUrl  string `json:"logoUrl" form:"logoUrl" gorm:"column:logo_url;comment:平台logo;size:100;"`  //平台logo 
}


// TableName 电商平台 Platform自定义表名 t_platform
func (Platform) TableName() string {
  return "t_platform"
}

