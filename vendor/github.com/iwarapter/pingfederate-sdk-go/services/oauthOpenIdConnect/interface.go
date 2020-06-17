package oauthOpenIdConnect

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthOpenIdConnectAPI interface {
	GetSettings() (output *models.OpenIdConnectSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (output *models.OpenIdConnectSettings, resp *http.Response, err error)
	GetPolicies() (output *models.OpenIdConnectPolicies, resp *http.Response, err error)
	CreatePolicy(input *CreatePolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)
	GetPolicy(input *GetPolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)
	UpdatePolicy(input *UpdatePolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)
	DeletePolicy(input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
}
