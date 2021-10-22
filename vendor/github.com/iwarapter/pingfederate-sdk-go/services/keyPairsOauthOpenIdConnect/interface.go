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

	GetKeySets() (output *models.AdditionalKeySets, resp *http.Response, err error)
	GetKeySetsWithContext(ctx context.Context) (output *models.AdditionalKeySets, resp *http.Response, err error)

	CreateKeySet(input *CreateKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error)
	CreateKeySetWithContext(ctx context.Context, input *CreateKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error)

	GetKeySet(input *GetKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error)
	GetKeySetWithContext(ctx context.Context, input *GetKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error)

	UpdateKeySet(input *UpdateKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error)
	UpdateKeySetWithContext(ctx context.Context, input *UpdateKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error)

	DeleteKeySet(input *DeleteKeySetInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteKeySetWithContext(ctx context.Context, input *DeleteKeySetInput) (output *models.ApiResult, resp *http.Response, err error)
}
