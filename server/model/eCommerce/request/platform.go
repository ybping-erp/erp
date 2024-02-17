package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eCommerce"
)

type PlatformSearch struct {
	eCommerce.Platform
	request.PageInfo
}
