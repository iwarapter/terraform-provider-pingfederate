package authenticationApi

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationApiAPI interface {
	GetAuthenticationApiSettings() (result *models.AuthnApiSettings, resp *http.Response, err error)
	UpdateAuthenticationApiSettings(input *UpdateAuthenticationApiSettingsInput) (result *models.AuthnApiSettings, resp *http.Response, err error)
	GetAuthenticationApiApplications() (result *models.AuthnApiApplications, resp *http.Response, err error)
	CreateApplication(input *CreateApplicationInput) (result *models.AuthnApiApplication, resp *http.Response, err error)
	GetApplication(input *GetApplicationInput) (result *models.AuthnApiApplication, resp *http.Response, err error)
	UpdateApplication(input *UpdateApplicationInput) (result *models.AuthnApiApplication, resp *http.Response, err error)
	DeleteApplication(input *DeleteApplicationInput) (result *models.ApiResult, resp *http.Response, err error)
}
