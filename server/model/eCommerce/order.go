// 自动生成模板Order
package eCommerce

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wms"
)

// 订单 结构体  Order
type Order struct {
	global.GVA_MODEL
	PlatformName          string           `json:"platformName" form:"platformName" gorm:"column:platform_name;comment:电商平台;size:50;"`                              //电商平台
	ShopId                string           `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:店铺唯一标识符;size:255;"`                                            //店铺ID
	ShopName              string           `json:"shopName" form:"shopName" gorm:"column:shop_name;comment:店铺名称;size:50;"`                                          //店铺名称
	PlatormOrderId        string           `json:"platformOrderId" form:"platformOrderId" gorm:"column:platform_order_id;comment:店铺订单号;size:50;"`                   //店铺订单号
	TotalAmount           *float64         `json:"totalAmount" form:"totalAmount" gorm:"column:total_amount;comment:订单总金额;size:12;"`                                //订单总金额
	Discount              *float64         `json:"discount" form:"discount" gorm:"column:discount;comment:订单折扣金额;size:5;"`                                          //订单折扣金额
	Tax                   *float64         `json:"tax" form:"tax" gorm:"column:tax;comment:订单税额;size:8;"`                                                           //订单税额
	StatusId              *int             `json:"statusId" form:"statusId" gorm:"column:status_id;comment:订单当前状态ID;size:20;"`                                      //状态ID
	CustomerId            string           `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户的标识符;size:50;"`                                  //客户ID
	CustomerName          string           `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名称;size:50;"`                              //客户名称
	CustomerTel           string           `json:"customerTel" form:"customerTel" gorm:"column:customer_tel;comment:客户电话;size:50;"`                                 //客户电话
	CustomerEmail         string           `json:"customerEmail" form:"customerEmail" gorm:"column:customer_email;comment:客户邮箱;size:50;"`                           //客户邮箱
	CustomerRemarks       string           `json:"customerRemarks" form:"customerRemarks" gorm:"column:customer_remarks;comment:客户备注;size:2048;"`                   //客户备注
	PaymentMethod         string           `json:"paymentMethod" form:"paymentMethod" gorm:"column:payment_method;comment:支付方式;size:50;"`                           //支付方式
	PaymentAt             *time.Time       `json:"paymentAt" form:"paymentAt" gorm:"column:payment_at;comment:付款时间;"`                                               //付款时间
	ShippingOrderId       string           `json:"shippingOrderId" form:"shippingOrderId" gorm:"column:shipping_order_id;comment:运单号;size:50;"`                     //运单号
	ShippingCost          *float64         `json:"shippingCost" form:"shippingCost" gorm:"column:shipping_cost;comment:订单运费;size:8;"`                               //运费
	ShippingStreetAddress string           `json:"shippingStreetAddress" form:"shippingStreetAddress" gorm:"column:shipping_street_address;comment:街道地址;size:255;"` //街道地址
	ShippingCity          string           `json:"shippingCity" form:"shippingCity" gorm:"column:shipping_city;comment:城市;size:100;"`                               //城市
	ShippingState         string           `json:"shippingState" form:"shippingState" gorm:"column:shipping_state;comment:省/州;size:100;"`                           //省/州
	ShippingPostalCode    string           `json:"shippingPostalCode" form:"shippingPostalCode" gorm:"column:shipping_postal_code;comment:邮政编码;size:20;"`           //邮政编码
	ShippingCountry       string           `json:"shippingCountry" form:"shippingCountry" gorm:"column:shipping_country;comment:国家;size:100;"`                      //国家
	ShipAt                *time.Time       `json:"shipAt" form:"shipAt" gorm:"column:ship_at;comment:发货时间;"`                                                        //发货时间
	RefundAt              *time.Time       `json:"refundAt" form:"refundAt" gorm:"column:refund_at;comment:退款时间;"`                                                  //退款时间
	Items                 []*OrderItem     `json:"orderItems" form:"orderItems" gorm:"foreignKey:OrderId;comment:订单明细"`
	PickOrders            []*wms.PickOrder `json:"pickOrders" form:"pickOrders" gorm:"foreignKey:OrderId;comment:拣选单"`
}

// TableName 订单 Order自定义表名 t_order
func (Order) TableName() string {
	return "t_order"
}
