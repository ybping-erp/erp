package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PlatformSearch struct{
    eCommerce.Platform
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
