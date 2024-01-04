package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type RackSearch struct{
    wms.Rack
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
