package oauthAccessTokenManagers

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAccessTokenManagersAPI interface {
	GetSettings() (output *models.AccessTokenManagementSettings, resp *http.Response, err error)
	GetSettingsWithContext(ctx context.Context) (output *models.AccessTokenManagementSettings, resp *http.Response, err error)

	UpdateSettings(input *UpdateSettingsInput) (output *models.AccessTokenManagementSettings, resp *http.Response, err error)
	UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.AccessTokenManagementSettings, resp *http.Response, err error)

	GetTokenManagerDescriptors() (output *models.AccessTokenManagerDescriptors, resp *http.Response, err error)
	GetTokenManagerDescriptorsWithContext(ctx context.Context) (output *models.AccessTokenManagerDescriptors, resp *http.Response, err error)

	GetTokenManagerDescriptor(input *GetTokenManagerDescriptorInput) (output *models.AccessTokenManagerDescriptor, resp *http.Response, err error)
	GetTokenManagerDescriptorWithContext(ctx context.Context, input *GetTokenManagerDescriptorInput) (output *models.AccessTokenManagerDescriptor, resp *http.Response, err error)

	GetTokenManagers() (output *models.AccessTokenManagers, resp *http.Response, err error)
	GetTokenManagersWithContext(ctx context.Context) (output *models.AccessTokenManagers, resp *http.Response, err error)

	CreateTokenManager(input *CreateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)
	CreateTokenManagerWithContext(ctx context.Context, input *CreateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)

	GetTokenManager(input *GetTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)
	GetTokenManagerWithContext(ctx context.Context, input *GetTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)

	UpdateTokenManager(input *UpdateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)
	UpdateTokenManagerWithContext(ctx context.Context, input *UpdateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error)

	DeleteTokenManager(input *DeleteTokenManagerInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteTokenManagerWithContext(ctx context.Context, input *DeleteTokenManagerInput) (output *models.ApiResult, resp *http.Response, err error)
}
