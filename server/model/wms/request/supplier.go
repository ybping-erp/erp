package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type SupplierSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Name  string `json:"name" form:"name" `
                      ContactName  string `json:"contactName" form:"contactName" `
                      ContactTelephone  string `json:"contactTelephone" form:"contactTelephone" `
    request.PageInfo
}
