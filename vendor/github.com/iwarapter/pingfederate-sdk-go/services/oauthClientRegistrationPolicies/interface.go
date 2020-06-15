package oauthClientRegistrationPolicies

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientRegistrationPoliciesAPI interface {
	GetDynamicClientRegistrationDescriptors() (result *models.ClientRegistrationPolicyDescriptors, resp *http.Response, err error)
	GetDynamicClientRegistrationDescriptor(input *GetDynamicClientRegistrationDescriptorInput) (result *models.ClientRegistrationPolicyDescriptor, resp *http.Response, err error)
	GetDynamicClientRegistrationPolicies() (result *models.ClientRegistrationPolicies, resp *http.Response, err error)
	CreateDynamicClientRegistrationPolicy(input *CreateDynamicClientRegistrationPolicyInput) (result *models.ClientRegistrationPolicy, resp *http.Response, err error)
	GetDynamicClientRegistrationPolicy(input *GetDynamicClientRegistrationPolicyInput) (result *models.ClientRegistrationPolicy, resp *http.Response, err error)
	UpdateDynamicClientRegistrationPolicy(input *UpdateDynamicClientRegistrationPolicyInput) (result *models.ClientRegistrationPolicy, resp *http.Response, err error)
	DeleteDynamicClientRegistrationPolicy(input *DeleteDynamicClientRegistrationPolicyInput) (result *models.ApiResult, resp *http.Response, err error)
}
