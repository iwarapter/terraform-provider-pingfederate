package session

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SessionAPI interface {
	GetSessionSettings() (output *models.SessionSettings, resp *http.Response, err error)
	UpdateSessionSettings(input *UpdateSessionSettingsInput) (output *models.SessionSettings, resp *http.Response, err error)
	GetGlobalPolicy() (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error)
	UpdateGlobalPolicy(input *UpdateGlobalPolicyInput) (output *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error)
	GetApplicationPolicy() (output *models.ApplicationSessionPolicy, resp *http.Response, err error)
	UpdateApplicationPolicy(input *UpdateApplicationPolicyInput) (output *models.ApplicationSessionPolicy, resp *http.Response, err error)
	GetSourcePolicies() (output *models.AuthenticationSessionPolicies, resp *http.Response, err error)
	CreateSourcePolicy(input *CreateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	GetSourcePolicy(input *GetSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	UpdateSourcePolicy(input *UpdateSourcePolicyInput) (output *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	DeleteSourcePolicy(input *DeleteSourcePolicyInput) (output *models.ApiResult, resp *http.Response, err error)
}
