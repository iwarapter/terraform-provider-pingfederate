package oauthCibaServerPolicy

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthCibaServerPolicyAPI interface {
	GetSettings() (output *models.CibaServerPolicySettings, resp *http.Response, err error)
	GetSettingsWithContext(ctx context.Context) (output *models.CibaServerPolicySettings, resp *http.Response, err error)

	UpdateSettings(input *UpdateSettingsInput) (output *models.CibaServerPolicySettings, resp *http.Response, err error)
	UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.CibaServerPolicySettings, resp *http.Response, err error)

	GetPolicies() (output *models.RequestPolicies, resp *http.Response, err error)
	GetPoliciesWithContext(ctx context.Context) (output *models.RequestPolicies, resp *http.Response, err error)

	CreatePolicy(input *CreatePolicyInput) (output *models.RequestPolicy, resp *http.Response, err error)
	CreatePolicyWithContext(ctx context.Context, input *CreatePolicyInput) (output *models.RequestPolicy, resp *http.Response, err error)

	GetPolicy(input *GetPolicyInput) (output *models.RequestPolicy, resp *http.Response, err error)
	GetPolicyWithContext(ctx context.Context, input *GetPolicyInput) (output *models.RequestPolicy, resp *http.Response, err error)

	UpdatePolicy(input *UpdatePolicyInput) (output *models.RequestPolicy, resp *http.Response, err error)
	UpdatePolicyWithContext(ctx context.Context, input *UpdatePolicyInput) (output *models.RequestPolicy, resp *http.Response, err error)

	DeletePolicy(input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
	DeletePolicyWithContext(ctx context.Context, input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
}
