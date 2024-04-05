// 自动生成模板Shop
package eCommerce

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"golang.org/x/oauth2"
)

// 店铺 结构体  Shop
type Shop struct {
	global.GVA_MODEL
	PlatformName    string     `json:"platformName" form:"platformName" gorm:"column:platform_name;comment:平台名称;size:255;"` //平台名称
	ShopId          string     `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:平台店铺唯一标识符;size:255;"`              //平台ID
	ShopName        string     `json:"shopName" form:"shopName" gorm:"column:shop_name;comment:店铺名称;size:255;"`             //店铺名称
	CSRFToken       string     `json:"-" form:"-" gorm:"column:scrf_token;comment:授权CSRF Token;size:255;"`
	AccessTokenType string     `json:"-" form:"-" gorm:"column:access_token_type;comment:oAuth Access Token Type;size:32"`
	AccessToken     string     `json:"-" form:"-" gorm:"column:access_token;comment:oAuth Access Token;size:4096"`
	RefreshToken    string     `json:"-" form:"-" gorm:"column:refresh_token;comment:oAuth Refresh Token;size:1024"`
	Expiry          *time.Time `json:"-" form:"-" gorm:"column:expiry;comment:oAuth Token过期时间"`
	CreatedBy       uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 店铺 Shop自定义表名 t_shop
func (Shop) TableName() string {
	return "t_shop"
}

// 保存新的的oauth token，并重置SCRF Token
func (shop *Shop) ResetToken(token *oauth2.Token) {
	shop.AccessTokenType = token.Type()
	shop.AccessToken = token.AccessToken
	shop.RefreshToken = token.RefreshToken
	shop.Expiry = &token.Expiry
	shop.CSRFToken = ""
}

func (shop *Shop) GenerateOAuthState() error {
	token, err := utils.GenerateCSRFToken()
	shop.CSRFToken = fmt.Sprintf("%s#%v", token, shop.ID)
	return err
}

func (Shop) GetShopIDFromOAuthState(state string) (int, error) {
	parts := strings.Split(state, "#")
	if len(parts) != 2 {
		return 0, fmt.Errorf("非法oAuthState: %s", state)
	}

	shopID, err := strconv.Atoi(parts[1])
	return shopID, err
}
