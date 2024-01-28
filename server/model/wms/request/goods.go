package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GoodsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	CategoryId  *int   `json:"categoryId" form:"categoryId" `
	SalesMethod *int   `json:"salesMethod" form:"salesMethod" `
	Sku         string `json:"sku" form:"sku" `
	Status      *int   `json:"status" form:"status" `
	request.PageInfo
}
