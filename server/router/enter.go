package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/eCommerce"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	ECommerce eCommerce.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
