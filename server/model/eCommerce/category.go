// 自动生成模板Category
package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// 品类 结构体  Category
type Category struct {
	global.GVA_MODEL
	Domain    string     `json:"domain" form:"domain" gorm:"<-:create;column:domain;comment:类型场景;"`         //品类场景
	Name      string     `json:"name" form:"name" gorm:"column:name;comment:类别名称;size:100;"`                //类别名称
	ParentId  int        `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父类别的标识符;size:10;"` //父类别的标识符
	Children  []Category `json:"children" gorm:"-"`
	CreatedBy uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 品类 Category自定义表名 t_category
func (Category) TableName() string {
	return "t_category"
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if category.ParentId != 0 {
		var parent Category
		err := tx.First(&parent, category.ParentId).Error
		if err != nil {
			return err
		}
		category.Domain = parent.Domain
	}
	return nil
}
