// 自动生成模板Supplier
package wms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 供应商 结构体  Supplier
type Supplier struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:供应商名称;size:255;"`  //名称 
      CompanyAddress  string `json:"companyAddress" form:"companyAddress" gorm:"column:company_address;comment:公司地址;size:255;"`  //公司地址 
      CompanyWebsite  string `json:"companyWebsite" form:"companyWebsite" gorm:"column:company_website;comment:公司网址;size:255;"`  //公司网址 
      ContactName  string `json:"contactName" form:"contactName" gorm:"column:contact_name;comment:联系人;size:255;"`  //联系人 
      ContactTelephone  string `json:"contactTelephone" form:"contactTelephone" gorm:"column:contact_telephone;comment:联系电话;size:255;"`  //联系电话 
      BankName  string `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:银行名称;size:255;"`  //银行名称 
      BankAccountNumber  string `json:"bankAccountNumber" form:"bankAccountNumber" gorm:"column:bank_account_number;comment:银行卡号;size:255;"`  //银行卡号 
      BankAccountName  string `json:"bankAccountName" form:"bankAccountName" gorm:"column:bank_account_name;comment:收款人;size:255;"`  //收款人 
      PaymentMethod  *int `json:"paymentMethod" form:"paymentMethod" gorm:"column:payment_method;comment:付款方式;size:10;"`  //付款方式 
      SettlementMethod  *int `json:"settlementMethod" form:"settlementMethod" gorm:"column:settlement_method;comment:结算方式;size:10;"`  //结算方式 
      Remarks  string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注信息;size:1024;"`  //备注信息 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 供应商 Supplier自定义表名 t_sd_supplier
func (Supplier) TableName() string {
  return "t_sd_supplier"
}

