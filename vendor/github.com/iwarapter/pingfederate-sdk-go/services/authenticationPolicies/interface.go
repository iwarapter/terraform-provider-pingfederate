package authenticationPolicies

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationPoliciesAPI interface {
	GetSettings() (result *models.AuthenticationPoliciesSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.AuthenticationPoliciesSettings, resp *http.Response, err error)
	GetDefaultAuthenticationPolicy() (result *models.AuthenticationPolicy, resp *http.Response, err error)
	UpdateDefaultAuthenticationPolicy(input *UpdateDefaultAuthenticationPolicyInput) (result *models.AuthenticationPolicy, resp *http.Response, err error)
}
