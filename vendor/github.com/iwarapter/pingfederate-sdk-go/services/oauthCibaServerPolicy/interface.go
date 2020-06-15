package oauthCibaServerPolicy

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthCibaServerPolicyAPI interface {
	GetSettings() (result *models.CibaServerPolicySettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.CibaServerPolicySettings, resp *http.Response, err error)
	GetPolicies() (result *models.RequestPolicies, resp *http.Response, err error)
	CreatePolicy(input *CreatePolicyInput) (result *models.RequestPolicy, resp *http.Response, err error)
	GetPolicy(input *GetPolicyInput) (result *models.RequestPolicy, resp *http.Response, err error)
	UpdatePolicy(input *UpdatePolicyInput) (result *models.RequestPolicy, resp *http.Response, err error)
	DeletePolicy(input *DeletePolicyInput) (result *models.ApiResult, resp *http.Response, err error)
}
