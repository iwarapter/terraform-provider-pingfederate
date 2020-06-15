package oauthOpenIdConnect

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthOpenIdConnectAPI interface {
	GetSettings() (result *models.OpenIdConnectSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.OpenIdConnectSettings, resp *http.Response, err error)
	GetPolicies() (result *models.OpenIdConnectPolicies, resp *http.Response, err error)
	CreatePolicy(input *CreatePolicyInput) (result *models.OpenIdConnectPolicy, resp *http.Response, err error)
	GetPolicy(input *GetPolicyInput) (result *models.OpenIdConnectPolicy, resp *http.Response, err error)
	UpdatePolicy(input *UpdatePolicyInput) (result *models.OpenIdConnectPolicy, resp *http.Response, err error)
	DeletePolicy(input *DeletePolicyInput) (result *models.ApiResult, resp *http.Response, err error)
}
