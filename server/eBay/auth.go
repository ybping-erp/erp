package eBay

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// API - oAuth api helper class
type API struct {
	appAccessMap sync.Map
	credential   config.EBayCredential
}

var apiInstance *API
var apiMutex sync.Mutex

func GetAPIClient() *API {
	// Lock the mutex to ensure only one goroutine can access the critical section
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if apiInstance == nil {
		apiInstance = &API{
			appAccessMap: sync.Map{},
			credential:   global.GVA_CONFIG.Ebay,
		}
	}
	return apiInstance
}

type applicationTokenSource struct {
	source oauth2.TokenSource
}

func (app *applicationTokenSource) Token() (*oauth2.Token, error) {

	token, err := app.source.Token()

	if err != nil {
		return nil, err
	}

	token.TokenType = ""

	return token, nil

}

func (api *API) GetDefaultScopes() string {
	// return "https://api.ebay.com/oauth/api_scope https://api.ebay.com/oauth/api_scope/buy.order.readonly https://api.ebay.com/oauth/api_scope/buy.guest.order https://api.ebay.com/oauth/api_scope/sell.marketing.readonly https://api.ebay.com/oauth/api_scope/sell.marketing https://api.ebay.com/oauth/api_scope/sell.inventory.readonly https://api.ebay.com/oauth/api_scope/sell.inventory https://api.ebay.com/oauth/api_scope/sell.account.readonly https://api.ebay.com/oauth/api_scope/sell.account https://api.ebay.com/oauth/api_scope/sell.fulfillment.readonly https://api.ebay.com/oauth/api_scope/sell.fulfillment https://api.ebay.com/oauth/api_scope/sell.analytics.readonly https://api.ebay.com/oauth/api_scope/sell.marketplace.insights.readonly https://api.ebay.com/oauth/api_scope/commerce.catalog.readonly https://api.ebay.com/oauth/api_scope/buy.shopping.cart https://api.ebay.com/oauth/api_scope/buy.offer.auction https://api.ebay.com/oauth/api_scope/commerce.identity.readonly https://api.ebay.com/oauth/api_scope/commerce.identity.email.readonly https://api.ebay.com/oauth/api_scope/commerce.identity.phone.readonly https://api.ebay.com/oauth/api_scope/commerce.identity.address.readonly https://api.ebay.com/oauth/api_scope/commerce.identity.name.readonly https://api.ebay.com/oauth/api_scope/commerce.identity.status.readonly https://api.ebay.com/oauth/api_scope/sell.finances https://api.ebay.com/oauth/api_scope/sell.payment.dispute https://api.ebay.com/oauth/api_scope/sell.item.draft https://api.ebay.com/oauth/api_scope/sell.item https://api.ebay.com/oauth/api_scope/sell.reputation https://api.ebay.com/oauth/api_scope/sell.reputation.readonly https://api.ebay.com/oauth/api_scope/commerce.notification.subscription https://api.ebay.com/oauth/api_scope/commerce.notification.subscription.readonly https://api.ebay.com/oauth/api_scope/sell.stores https://api.ebay.com/oauth/api_scope/sell.stores.readonly"
	return "https://api.ebay.com/oauth/api_scope/sandbox https://api.ebay.com/oauth/api_scope/sandbox/buy.order.readonly https://api.ebay.com/oauth/api_scope/sandbox/buy.guest.order https://api.ebay.com/oauth/api_scope/sandbox/sell.marketing.readonly https://api.ebay.com/oauth/api_scope/sandbox/sell.marketing https://api.ebay.com/oauth/api_scope/sandbox/sell.inventory.readonly https://api.ebay.com/oauth/api_scope/sandbox/sell.inventory https://api.ebay.com/oauth/api_scope/sandbox/sell.account.readonly https://api.ebay.com/oauth/api_scope/sandbox/sell.account https://api.ebay.com/oauth/api_scope/sandbox/sell.fulfillment.readonly https://api.ebay.com/oauth/api_scope/sandbox/sell.fulfillment https://api.ebay.com/oauth/api_scope/sandbox/sell.analytics.readonly https://api.ebay.com/oauth/api_scope/sandbox/sell.marketplace.insights.readonly https://api.ebay.com/oauth/api_scope/sandbox/commerce.catalog.readonly https://api.ebay.com/oauth/api_scope/sandbox/buy.shopping.cart https://api.ebay.com/oauth/api_scope/sandbox/buy.offer.auction https://api.ebay.com/oauth/api_scope/sandbox/commerce.identity.readonly https://api.ebay.com/oauth/api_scope/sandbox/commerce.identity.email.readonly https://api.ebay.com/oauth/api_scope/sandbox/commerce.identity.phone.readonly https://api.ebay.com/oauth/api_scope/sandbox/commerce.identity.address.readonly https://api.ebay.com/oauth/api_scope/sandbox/commerce.identity.name.readonly https://api.ebay.com/oauth/api_scope/sandbox/commerce.identity.status.readonly https://api.ebay.com/oauth/api_scope/sandbox/sell.finances https://api.ebay.com/oauth/api_scope/sandbox/sell.payment.dispute https://api.ebay.com/oauth/api_scope/sandbox/sell.item.draft https://api.ebay.com/oauth/api_scope/sandbox/sell.item https://api.ebay.com/oauth/api_scope/sandbox/sell.reputation https://api.ebay.com/oauth/api_scope/sandbox/sell.reputation.readonly https://api.ebay.com/oauth/api_scope/sandbox/commerce.notification.subscription https://api.ebay.com/oauth/api_scope/sandbox/commerce.notification.subscription.readonly https://api.ebay.com/oauth/api_scope/sandbox/sell.stores https://api.ebay.com/oauth/api_scope/sandbox/sell.stores.readonly"
}

// GetApplicationTokenAndClient - returns api client credentials token and authenticated http client
func (api *API) GetApplicationTokenAndClient(ctx context.Context, scopes ...string) (*oauth2.Token, *http.Client, error) {

	if ts, ok := api.appAccessMap.Load(api.credential.Environment); ok {
		log.Debug("application access token returned from cache")
		tokenSource := ts.(oauth2.TokenSource)
		token, err := tokenSource.Token()
		if err != nil {
			return nil, nil, fmt.Errorf("unable to generate token for %s environment: %s", api.credential.Environment, err.Error())
		}

		return token, oauth2.NewClient(ctx, tokenSource), nil
	}

	config := generateClientCredentialsConfig(&api.credential, scopes)

	tokenSource := &applicationTokenSource{source: config.TokenSource(ctx)}

	api.appAccessMap.Store(api.credential.Environment, tokenSource)

	token, err := tokenSource.Token()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to generate token for %s environment: %s", api.credential.Environment, err.Error())
	}

	return token, oauth2.NewClient(ctx, tokenSource), nil
}

// GenerateUserAuthorizationURL - returns authorization URL to send user to
// The state parameter in OAuth serves as a security measure to mitigate against cross-site request forgery (CSRF) attacks and to maintain the state of the application during the OAuth flow.
func (api *API) GenerateUserAuthorizationURL(ctx context.Context, state string, scopes ...string) (string, error) {
	config := generateOAuth2Config(&api.credential, scopes)
	return config.AuthCodeURL(state), nil
}

// ExchangeCodeForAccessTokenAndClient - exchange access token for oAuth token and authenticated http client
func (api *API) ExchangeCodeForAccessTokenAndClient(ctx context.Context, code string) (*oauth2.Token, *http.Client, error) {
	config := generateOAuth2Config(&api.credential, nil)
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to generate token for %s environment: %s", api.credential.Environment, err.Error())
	}
	tokenSource := config.TokenSource(ctx, token)
	return token, oauth2.NewClient(ctx, tokenSource), nil

}

// GetAccessTokenAndClient - get token and authenticated http client with access token from refresh token
func (api *API) GetAccessTokenAndClient(ctx context.Context, refreshToken string, scopes ...string) (*oauth2.Token, *http.Client, error) {
	config := generateOAuth2Config(&api.credential, scopes)
	token := &oauth2.Token{
		RefreshToken: refreshToken,
		Expiry:       time.Now(),
	}

	tokenSource := config.TokenSource(ctx, token)

	token, err := tokenSource.Token()

	// Change the token type to bearer for http.Client Authorization purpose
	token.TokenType = "bearer"
	if err != nil || token == nil {
		return nil, nil, fmt.Errorf("unable to generate token for %s environment: %s", api.credential.Environment, err.Error())
	}

	return token, config.Client(ctx, token), nil

}

func generateClientCredentialsConfig(credential *config.EBayCredential, scopes []string) *clientcredentials.Config {
	return &clientcredentials.Config{
		ClientID:     credential.AppID,
		ClientSecret: credential.CertID,
		TokenURL:     credential.ApiEndpoint,
		Scopes:       scopes,
	}
}

func generateOAuth2Config(credential *config.EBayCredential, scopes []string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     credential.AppID,
		ClientSecret: credential.CertID,
		Endpoint: oauth2.Endpoint{
			AuthURL:  credential.WebEndpoint,
			TokenURL: credential.ApiEndpoint,
		},
		RedirectURL: credential.RedirectURI,
		Scopes:      scopes,
	}
}
