package oauthClientRegistrationPolicies

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientRegistrationPoliciesAPI interface {
	GetDynamicClientRegistrationDescriptors() (output *models.ClientRegistrationPolicyDescriptors, resp *http.Response, err error)
	GetDynamicClientRegistrationDescriptor(input *GetDynamicClientRegistrationDescriptorInput) (output *models.ClientRegistrationPolicyDescriptor, resp *http.Response, err error)
	GetDynamicClientRegistrationPolicies() (output *models.ClientRegistrationPolicies, resp *http.Response, err error)
	CreateDynamicClientRegistrationPolicy(input *CreateDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)
	GetDynamicClientRegistrationPolicy(input *GetDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)
	UpdateDynamicClientRegistrationPolicy(input *UpdateDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)
	DeleteDynamicClientRegistrationPolicy(input *DeleteDynamicClientRegistrationPolicyInput) (output *models.ApiResult, resp *http.Response, err error)
}
