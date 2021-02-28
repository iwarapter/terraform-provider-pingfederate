package oauthOpenIdConnect

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthOpenIdConnectAPI interface {
	GetSettings() (output *models.OpenIdConnectSettings, resp *http.Response, err error)
	GetSettingsWithContext(ctx context.Context) (output *models.OpenIdConnectSettings, resp *http.Response, err error)

	UpdateSettings(input *UpdateSettingsInput) (output *models.OpenIdConnectSettings, resp *http.Response, err error)
	UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.OpenIdConnectSettings, resp *http.Response, err error)

	GetPolicies() (output *models.OpenIdConnectPolicies, resp *http.Response, err error)
	GetPoliciesWithContext(ctx context.Context) (output *models.OpenIdConnectPolicies, resp *http.Response, err error)

	CreatePolicy(input *CreatePolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)
	CreatePolicyWithContext(ctx context.Context, input *CreatePolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)

	GetPolicy(input *GetPolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)
	GetPolicyWithContext(ctx context.Context, input *GetPolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)

	UpdatePolicy(input *UpdatePolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)
	UpdatePolicyWithContext(ctx context.Context, input *UpdatePolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error)

	DeletePolicy(input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
	DeletePolicyWithContext(ctx context.Context, input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
}
