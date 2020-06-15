package session

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SessionAPI interface {
	GetSessionSettings() (result *models.SessionSettings, resp *http.Response, err error)
	UpdateSessionSettings(input *UpdateSessionSettingsInput) (result *models.SessionSettings, resp *http.Response, err error)
	GetGlobalPolicy() (result *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error)
	UpdateGlobalPolicy(input *UpdateGlobalPolicyInput) (result *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error)
	GetApplicationPolicy() (result *models.ApplicationSessionPolicy, resp *http.Response, err error)
	UpdateApplicationPolicy(input *UpdateApplicationPolicyInput) (result *models.ApplicationSessionPolicy, resp *http.Response, err error)
	GetSourcePolicies() (result *models.AuthenticationSessionPolicies, resp *http.Response, err error)
	CreateSourcePolicy(input *CreateSourcePolicyInput) (result *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	GetSourcePolicy(input *GetSourcePolicyInput) (result *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	UpdateSourcePolicy(input *UpdateSourcePolicyInput) (result *models.AuthenticationSessionPolicy, resp *http.Response, err error)
	DeleteSourcePolicy(input *DeleteSourcePolicyInput) (result *models.ApiResult, resp *http.Response, err error)
}
