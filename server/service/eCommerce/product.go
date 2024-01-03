package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
)

type ProductService struct {
}

// CreateProduct 创建产品表记录

func (productService *ProductService) CreateProduct(product *eCommerce.Product) (err error) {
	err = global.GVA_DB.Create(product).Error
	return err
}

// DeleteProduct 删除产品表记录

func (productService *ProductService)DeleteProduct(product eCommerce.Product) (err error) {
	err = global.GVA_DB.Delete(&product).Error
	return err
}

// DeleteProductByIds 批量删除产品表记录

func (productService *ProductService)DeleteProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]eCommerce.Product{},"id in ?",ids.Ids).Error
	return err
}

// UpdateProduct 更新产品表记录

func (productService *ProductService)UpdateProduct(product eCommerce.Product) (err error) {
	err = global.GVA_DB.Save(&product).Error
	return err
}

// GetProduct 根据id获取产品表记录

func (productService *ProductService)GetProduct(id uint) (product eCommerce.Product, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&product).Error
	return
}

// GetProductInfoList 分页获取产品表记录

func (productService *ProductService)GetProductInfoList(info eCommerceReq.ProductSearch) (list []eCommerce.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&eCommerce.Product{})
    var products []eCommerce.Product
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CategoryId != nil {
        db = db.Where("category_id = ?",info.CategoryId)
    }
    if info.ProductName != "" {
        db = db.Where("product_name LIKE ?","%"+ info.ProductName+"%")
    }
    if info.Sku != "" {
        db = db.Where("sku = ?",info.Sku)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Preload("Creator").Find(&products).Error
	return  products, total, err
}
