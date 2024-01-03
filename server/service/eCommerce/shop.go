package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
)

type ShopService struct {
}

// CreateShop 创建店铺记录

func (shopService *ShopService) CreateShop(shop *eCommerce.Shop) (err error) {
	err = global.GVA_DB.Create(shop).Error
	return err
}

// DeleteShop 删除店铺记录

func (shopService *ShopService)DeleteShop(shop eCommerce.Shop) (err error) {
	err = global.GVA_DB.Delete(&shop).Error
	return err
}

// DeleteShopByIds 批量删除店铺记录

func (shopService *ShopService)DeleteShopByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]eCommerce.Shop{},"id in ?",ids.Ids).Error
	return err
}

// UpdateShop 更新店铺记录

func (shopService *ShopService)UpdateShop(shop eCommerce.Shop) (err error) {
	err = global.GVA_DB.Save(&shop).Error
	return err
}

// GetShop 根据id获取店铺记录

func (shopService *ShopService)GetShop(id uint) (shop eCommerce.Shop, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&shop).Error
	return
}

// GetShopInfoList 分页获取店铺记录

func (shopService *ShopService)GetShopInfoList(info eCommerceReq.ShopSearch) (list []eCommerce.Shop, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&eCommerce.Shop{})
    var shops []eCommerce.Shop
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.PlatformName != "" {
        db = db.Where("platform_name = ?",info.PlatformName)
    }
    if info.ShopId != "" {
        db = db.Where("shop_id = ?",info.ShopId)
    }
    if info.ShopName != "" {
        db = db.Where("shop_name LIKE ?","%"+ info.ShopName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Preload("Creator").Find(&shops).Error
	return  shops, total, err
}
