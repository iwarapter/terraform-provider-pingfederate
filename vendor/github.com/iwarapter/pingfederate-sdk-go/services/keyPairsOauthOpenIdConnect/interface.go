package keyPairsOauthOpenIdConnect

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsOauthOpenIdConnectAPI interface {
	GetOauthOidcKeysSettings() (output *models.OAuthOidcKeysSettings, resp *http.Response, err error)
	GetOauthOidcKeysSettingsWithContext(ctx context.Context) (output *models.OAuthOidcKeysSettings, resp *http.Response, err error)

	UpdateOAuthOidcKeysSettings(input *UpdateOAuthOidcKeysSettingsInput) (output *models.OAuthOidcKeysSettings, resp *http.Response, err error)
	UpdateOAuthOidcKeysSettingsWithContext(ctx context.Context, input *UpdateOAuthOidcKeysSettingsInput) (output *models.OAuthOidcKeysSettings, resp *http.Response, err error)
}
