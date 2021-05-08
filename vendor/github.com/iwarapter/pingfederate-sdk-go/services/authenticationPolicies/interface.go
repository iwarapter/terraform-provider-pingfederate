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

	GetFragments(input *GetFragmentsInput) (output *models.AuthenticationPolicyFragments, resp *http.Response, err error)
	GetFragmentsWithContext(ctx context.Context, input *GetFragmentsInput) (output *models.AuthenticationPolicyFragments, resp *http.Response, err error)

	CreateFragment(input *CreateFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error)
	CreateFragmentWithContext(ctx context.Context, input *CreateFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error)

	GetFragment(input *GetFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error)
	GetFragmentWithContext(ctx context.Context, input *GetFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error)

	UpdateFragment(input *UpdateFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error)
	UpdateFragmentWithContext(ctx context.Context, input *UpdateFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error)

	DeleteFragment(input *DeleteFragmentInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteFragmentWithContext(ctx context.Context, input *DeleteFragmentInput) (output *models.ApiResult, resp *http.Response, err error)
}
