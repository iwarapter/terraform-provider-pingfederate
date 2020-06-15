package oauthAccessTokenManagers

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAccessTokenManagersAPI interface {
	GetSettings() (result *models.AccessTokenManagementSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.AccessTokenManagementSettings, resp *http.Response, err error)
	GetTokenManagerDescriptors() (result *models.AccessTokenManagerDescriptors, resp *http.Response, err error)
	GetTokenManagerDescriptor(input *GetTokenManagerDescriptorInput) (result *models.AccessTokenManagerDescriptor, resp *http.Response, err error)
	GetTokenManagers() (result *models.AccessTokenManagers, resp *http.Response, err error)
	CreateTokenManager(input *CreateTokenManagerInput) (result *models.AccessTokenManager, resp *http.Response, err error)
	GetTokenManager(input *GetTokenManagerInput) (result *models.AccessTokenManager, resp *http.Response, err error)
	UpdateTokenManager(input *UpdateTokenManagerInput) (result *models.AccessTokenManager, resp *http.Response, err error)
	DeleteTokenManager(input *DeleteTokenManagerInput) (result *models.ApiResult, resp *http.Response, err error)
}
