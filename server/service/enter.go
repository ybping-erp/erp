package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/wms"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	ECommerceServiceGroup eCommerce.ServiceGroup
	WmsServiceGroup       wms.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
