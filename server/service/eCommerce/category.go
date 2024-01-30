package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
	"gorm.io/gorm"
)

type CategoryService struct {
}

// CreateCategory 创建品类记录

func (categoryService *CategoryService) CreateCategory(category *eCommerce.Category) (err error) {
	err = global.GVA_DB.Create(category).Error
	return err
}

// DeleteCategoryByIds 批量删除品类记录
func (categoryService *CategoryService) DeleteCategoryByIds(IDs []int, deletedBy uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&eCommerce.Category{}).Where("id in ?", IDs).Update("deleted_by", deletedBy).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&eCommerce.Category{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCategory 更新品类记录

func (categoryService *CategoryService) UpdateCategory(category eCommerce.Category) (err error) {
	err = global.GVA_DB.Save(&category).Error
	return err
}

// GetCategory 根据id获取品类记录

func (categoryService *CategoryService) GetCategory(id uint) (category eCommerce.Category, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryInfoList 分页获取品类记录

func (categoryService *CategoryService) GetCategoryInfoList(info eCommerceReq.CategorySearch) (list []eCommerce.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&eCommerce.Category{})
	if info.Domain != "" {
		db = db.Where("domain = ?", info.Domain)
	}
	var categorys []eCommerce.Category
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Preload("Creator").Find(&categorys).Error
	for k := range categorys {
		err = categoryService.findChildrenCategory(&categorys[k])
	}
	return categorys, total, err
}

// findChildrenCategory 查询子品类
func (categoryService *CategoryService) findChildrenCategory(category *eCommerce.Category) (err error) {
	err = global.GVA_DB.Where("parent_id = ?", category.ID).Preload("Creator").Find(&category.Children).Error
	if len(category.Children) > 0 {
		for k := range category.Children {
			err = categoryService.findChildrenCategory(&category.Children[k])
		}
	}
	return err
}
