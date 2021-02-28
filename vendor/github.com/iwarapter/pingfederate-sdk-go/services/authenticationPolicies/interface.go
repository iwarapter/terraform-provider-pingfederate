package authenticationPolicies

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationPoliciesAPI interface {
	GetSettings() (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error)
	GetSettingsWithContext(ctx context.Context) (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error)

	UpdateSettings(input *UpdateSettingsInput) (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error)
	UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error)

	GetDefaultAuthenticationPolicy() (output *models.AuthenticationPolicy, resp *http.Response, err error)
	GetDefaultAuthenticationPolicyWithContext(ctx context.Context) (output *models.AuthenticationPolicy, resp *http.Response, err error)

	UpdateDefaultAuthenticationPolicy(input *UpdateDefaultAuthenticationPolicyInput) (output *models.AuthenticationPolicy, resp *http.Response, err error)
	UpdateDefaultAuthenticationPolicyWithContext(ctx context.Context, input *UpdateDefaultAuthenticationPolicyInput) (output *models.AuthenticationPolicy, resp *http.Response, err error)
}
