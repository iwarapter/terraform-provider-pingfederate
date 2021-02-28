package authenticationApi

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationApiAPI interface {
	GetAuthenticationApiSettings() (output *models.AuthnApiSettings, resp *http.Response, err error)
	GetAuthenticationApiSettingsWithContext(ctx context.Context) (output *models.AuthnApiSettings, resp *http.Response, err error)

	UpdateAuthenticationApiSettings(input *UpdateAuthenticationApiSettingsInput) (output *models.AuthnApiSettings, resp *http.Response, err error)
	UpdateAuthenticationApiSettingsWithContext(ctx context.Context, input *UpdateAuthenticationApiSettingsInput) (output *models.AuthnApiSettings, resp *http.Response, err error)

	GetAuthenticationApiApplications() (output *models.AuthnApiApplications, resp *http.Response, err error)
	GetAuthenticationApiApplicationsWithContext(ctx context.Context) (output *models.AuthnApiApplications, resp *http.Response, err error)

	CreateApplication(input *CreateApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)
	CreateApplicationWithContext(ctx context.Context, input *CreateApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)

	GetApplication(input *GetApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)
	GetApplicationWithContext(ctx context.Context, input *GetApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)

	UpdateApplication(input *UpdateApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)
	UpdateApplicationWithContext(ctx context.Context, input *UpdateApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)

	DeleteApplication(input *DeleteApplicationInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteApplicationWithContext(ctx context.Context, input *DeleteApplicationInput) (output *models.ApiResult, resp *http.Response, err error)
}
