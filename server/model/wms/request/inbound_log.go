package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type InboundLogSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      GoodsSku  string `json:"goodsSku" form:"goodsSku" `
    request.PageInfo
}
