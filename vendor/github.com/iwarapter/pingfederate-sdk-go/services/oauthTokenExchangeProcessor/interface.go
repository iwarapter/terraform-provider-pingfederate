package oauthTokenExchangeProcessor

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeProcessorAPI interface {
	GetSettings() (result *models.TokenExchangeProcessorSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.TokenExchangeProcessorSettings, resp *http.Response, err error)
	GetPolicies() (result *models.TokenExchangeProcessorPolicies, resp *http.Response, err error)
	CreatePolicy(input *CreatePolicyInput) (result *models.TokenExchangeProcessorPolicy, resp *http.Response, err error)
	GetPolicy(input *GetPolicyInput) (result *models.TokenExchangeProcessorPolicy, resp *http.Response, err error)
	UpdatePolicy(input *UpdatePolicyInput) (result *models.TokenExchangeProcessorPolicy, resp *http.Response, err error)
	DeletePolicy(input *DeletePolicyInput) (result *models.ApiResult, resp *http.Response, err error)
}
