package authenticationApi

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationApiAPI interface {
	GetAuthenticationApiSettings() (output *models.AuthnApiSettings, resp *http.Response, err error)
	UpdateAuthenticationApiSettings(input *UpdateAuthenticationApiSettingsInput) (output *models.AuthnApiSettings, resp *http.Response, err error)
	GetAuthenticationApiApplications() (output *models.AuthnApiApplications, resp *http.Response, err error)
	CreateApplication(input *CreateApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)
	GetApplication(input *GetApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)
	UpdateApplication(input *UpdateApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error)
	DeleteApplication(input *DeleteApplicationInput) (output *models.ApiResult, resp *http.Response, err error)
}
