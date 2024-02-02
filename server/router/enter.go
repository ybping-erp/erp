package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/shared"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/wms"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	ECommerce eCommerce.RouterGroup
	Wms       wms.RouterGroup
	Shared    shared.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
