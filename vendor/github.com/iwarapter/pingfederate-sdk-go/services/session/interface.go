package session

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SessionAPI interface {
	GetSessionSettings() (output *models.SessionSettings, resp *http.Response, err error)
	GetSessionSettingsWithContext(ctx context.Context) (output *models.SessionSettings, resp *http.Response, err error)

	UpdateSessionSettings(input *UpdateSessionSettingsInput) (output *models.SessionSettings, resp *http.Response, err error)
	UpdateSessionSettingsWithContext(ctx context.Context, input *UpdateSessionSettingsInput) (output *models.SessionSettings, resp *http.Response, err error)

	GetApplicationPolicy() (output *models.ApplicationSessionPolicy, resp *http.Response, err error)
	GetApplicationPolicyWithContext(ctx context.Context) (output *models.ApplicationSessionPolicy, resp *http.Response, err error)

	UpdateApplicationPolicy(input *UpdateApplicationPolicyInput) (output *models.ApplicationSessionPolicy, resp *http.Response, err error)
	UpdateApplicationPolicyWithContext(ctx context.Context, input *UpdateApplicationPolicyInput) (output *models.ApplicationSessionPolicy, resp *http.Response, err error)

	GetGlobalPolicy() (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error)
	GetGlobalPolicyWithContext(ctx context.Context) (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error)

	UpdateGlobalPolicy(input *UpdateGlobalPolicyInput) (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error)
	UpdateGlobalPolicyWithContext(ctx context.Context, input *UpdateGlobalPolicyInput) (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error)

	GetSourcePolicies() (output *models.AuthenticationSessionPolicies, resp *http.Response, err error)
	GetSourcePoliciesWithContext(ctx context.Context) (output *models.AuthenticationSessionPolicies, resp *http.Response, err error)

	CreateSourcePolicy(input *CreateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	CreateSourcePolicyWithContext(ctx context.Context, input *CreateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)

	GetSourcePolicy(input *GetSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	GetSourcePolicyWithContext(ctx context.Context, input *GetSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)

	UpdateSourcePolicy(input *UpdateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	UpdateSourcePolicyWithContext(ctx context.Context, input *UpdateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)

	DeleteSourcePolicy(input *DeleteSourcePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteSourcePolicyWithContext(ctx context.Context, input *DeleteSourcePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
}
