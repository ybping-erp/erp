package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
    "gorm.io/gorm"
)

type GoodsService struct {
}

// CreateGoods 创建商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService) CreateGoods(goods *wms.Goods) (err error) {
	err = global.GVA_DB.Create(goods).Error
	return err
}

// DeleteGoods 删除商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)DeleteGoods(goods wms.Goods) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.Goods{}).Where("id = ?", goods.ID).Update("deleted_by", goods.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&goods).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteGoodsByIds 批量删除商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)DeleteGoodsByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&wms.Goods{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&wms.Goods{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateGoods 更新商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)UpdateGoods(goods wms.Goods) (err error) {
	err = global.GVA_DB.Save(&goods).Error
	return err
}

// GetGoods 根据id获取商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)GetGoods(id uint) (goods wms.Goods, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&goods).Error
	return
}

// GetGoodsInfoList 分页获取商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)GetGoodsInfoList(info wmsReq.GoodsSearch) (list []wms.Goods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.Goods{})
    var goodss []wms.Goods
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Sku != "" {
        db = db.Where("sku = ?",info.Sku)
    }
    if info.SalesMethod != nil {
        db = db.Where("sales_method = ?",info.SalesMethod)
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
    if info.ChineseName != "" {
        db = db.Where("chinese_name LIKE ?","%"+ info.ChineseName+"%")
    }
    if info.EnglishName != "" {
        db = db.Where("english_name LIKE ?","%"+ info.EnglishName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&goodss).Error
	return  goodss, total, err
}
