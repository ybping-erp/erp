package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type InboundLogSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Status         *int       `json:"status" form:"status" `
	GoodsSku       string     `json:"goodsSku" form:"goodsSku" `
	request.PageInfo
}
