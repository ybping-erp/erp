package wms

type ApiGroup struct {
	GoodsApi
	InventoryApi
	LogisticsPackagingApi
	RackApi
	WarehouseApi
	SkuMappingApi
	InboundLogApi
	OutboundLogApi
}
