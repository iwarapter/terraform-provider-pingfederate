package oauthAccessTokenManagers

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAccessTokenManagersAPI interface {
	GetSettings() (output *models.AccessTokenManagementSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (output *models.AccessTokenManagementSettings, resp *http.Response, err error)
	GetTokenManagerDescriptors() (output *models.AccessTokenManagerDescriptors, resp *http.Response, err error)
	GetTokenManagerDescriptor(input *GetTokenManagerDescriptorInput) (output *models.AccessTokenManagerDescriptor, resp *http.Response, err error)
	GetTokenManagers() (output *models.AccessTokenManagers, resp *http.Response, err error)
	CreateTokenManager(input *CreateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)
	GetTokenManager(input *GetTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)
	UpdateTokenManager(input *UpdateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)
	DeleteTokenManager(input *DeleteTokenManagerInput) (output *models.ApiResult, resp *http.Response, err error)
}
