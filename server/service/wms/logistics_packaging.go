package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    wmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/wms/request"
)

type LogisticsPackagingService struct {
}

// CreateLogisticsPackaging 创建物流包装记录
// Author [piexlmax](https://github.com/piexlmax)
func (logisticsPackagingService *LogisticsPackagingService) CreateLogisticsPackaging(logisticsPackaging *wms.LogisticsPackaging) (err error) {
	err = global.GVA_DB.Create(logisticsPackaging).Error
	return err
}

// DeleteLogisticsPackaging 删除物流包装记录
// Author [piexlmax](https://github.com/piexlmax)
func (logisticsPackagingService *LogisticsPackagingService)DeleteLogisticsPackaging(logisticsPackaging wms.LogisticsPackaging) (err error) {
	err = global.GVA_DB.Delete(&logisticsPackaging).Error
	return err
}

// DeleteLogisticsPackagingByIds 批量删除物流包装记录
// Author [piexlmax](https://github.com/piexlmax)
func (logisticsPackagingService *LogisticsPackagingService)DeleteLogisticsPackagingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wms.LogisticsPackaging{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLogisticsPackaging 更新物流包装记录
// Author [piexlmax](https://github.com/piexlmax)
func (logisticsPackagingService *LogisticsPackagingService)UpdateLogisticsPackaging(logisticsPackaging wms.LogisticsPackaging) (err error) {
	err = global.GVA_DB.Save(&logisticsPackaging).Error
	return err
}

// GetLogisticsPackaging 根据id获取物流包装记录
// Author [piexlmax](https://github.com/piexlmax)
func (logisticsPackagingService *LogisticsPackagingService)GetLogisticsPackaging(id uint) (logisticsPackaging wms.LogisticsPackaging, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&logisticsPackaging).Error
	return
}

// GetLogisticsPackagingInfoList 分页获取物流包装记录
// Author [piexlmax](https://github.com/piexlmax)
func (logisticsPackagingService *LogisticsPackagingService)GetLogisticsPackagingInfoList(info wmsReq.LogisticsPackagingSearch) (list []wms.LogisticsPackaging, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&wms.LogisticsPackaging{})
    var logisticsPackagings []wms.LogisticsPackaging
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.GoodsSku != "" {
        db = db.Where("goods_sku = ?",info.GoodsSku)
    }
    if info.SpecialGoods != nil {
        db = db.Where("special_goods = ?",info.SpecialGoods)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&logisticsPackagings).Error
	return  logisticsPackagings, total, err
}
