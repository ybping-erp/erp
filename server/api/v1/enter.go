package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/shared"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/wms"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	ECommerceApiGroup eCommerce.ApiGroup
	WmsApiGroup       wms.ApiGroup
	SharedApiGroup    shared.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
