package wms

type RouterGroup struct {
	GoodsRouter
	InventoryRouter
	LogisticsPackagingRouter
	RackRouter
	WarehouseRouter
	SkuMappingRouter
	InboundLogRouter
	OutboundLogRouter
	SupplierRouter
}
