package oauthTokenExchangeProcessor

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeProcessorAPI interface {
	GetSettings() (output *models.TokenExchangeProcessorSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (output *models.TokenExchangeProcessorSettings, resp *http.Response, err error)
	GetPolicies() (output *models.TokenExchangeProcessorPolicies, resp *http.Response, err error)
	CreatePolicy(input *CreatePolicyInput) (output *models.TokenExchangeProcessorPolicy, resp *http.Response, err error)
	GetPolicy(input *GetPolicyInput) (output *models.TokenExchangeProcessorPolicy, resp *http.Response, err error)
	UpdatePolicy(input *UpdatePolicyInput) (output *models.TokenExchangeProcessorPolicy, resp *http.Response, err error)
	DeletePolicy(input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
}
