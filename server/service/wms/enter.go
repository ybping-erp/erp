package wms

type ServiceGroup struct {
	GoodsService
	InventoryService
	LogisticsPackagingService
	RackService
	WarehouseService
	SkuMappingService
	InboundLogService
	OutboundLogService
}
