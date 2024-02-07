package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type PickOrderSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      GoodsName  string `json:"goodsName" form:"goodsName" `
                      GoodsSku  string `json:"goodsSku" form:"goodsSku" `
                      OrderId  *int `json:"orderId" form:"orderId" `
    request.PageInfo
}
