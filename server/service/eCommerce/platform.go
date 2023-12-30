package eCommerce

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    eCommerceReq "github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce/request"
)

type PlatformService struct {
}

// CreatePlatform 创建电商平台记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformService *PlatformService) CreatePlatform(platform *eCommerce.Platform) (err error) {
	err = global.GVA_DB.Create(platform).Error
	return err
}

// DeletePlatform 删除电商平台记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformService *PlatformService)DeletePlatform(platform eCommerce.Platform) (err error) {
	err = global.GVA_DB.Delete(&platform).Error
	return err
}

// DeletePlatformByIds 批量删除电商平台记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformService *PlatformService)DeletePlatformByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]eCommerce.Platform{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePlatform 更新电商平台记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformService *PlatformService)UpdatePlatform(platform eCommerce.Platform) (err error) {
	err = global.GVA_DB.Save(&platform).Error
	return err
}

// GetPlatform 根据id获取电商平台记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformService *PlatformService)GetPlatform(id uint) (platform eCommerce.Platform, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&platform).Error
	return
}

// GetPlatformInfoList 分页获取电商平台记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformService *PlatformService)GetPlatformInfoList(info eCommerceReq.PlatformSearch) (list []eCommerce.Platform, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&eCommerce.Platform{})
    var platforms []eCommerce.Platform
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&platforms).Error
	return  platforms, total, err
}
