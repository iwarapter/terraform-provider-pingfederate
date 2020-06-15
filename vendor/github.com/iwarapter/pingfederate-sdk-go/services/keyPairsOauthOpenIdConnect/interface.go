package keyPairsOauthOpenIdConnect

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsOauthOpenIdConnectAPI interface {
	GetOauthOidcKeysSettings() (result *models.OAuthOidcKeysSettings, resp *http.Response, err error)
	UpdateOAuthOidcKeysSettings(input *UpdateOAuthOidcKeysSettingsInput) (result *models.OAuthOidcKeysSettings, resp *http.Response, err error)
}
