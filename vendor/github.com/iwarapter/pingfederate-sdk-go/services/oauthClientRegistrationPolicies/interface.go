package oauthClientRegistrationPolicies

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientRegistrationPoliciesAPI interface {
	GetDynamicClientRegistrationDescriptors() (output *models.ClientRegistrationPolicyDescriptors, resp *http.Response, err error)
	GetDynamicClientRegistrationDescriptorsWithContext(ctx context.Context) (output *models.ClientRegistrationPolicyDescriptors, resp *http.Response, err error)

	GetDynamicClientRegistrationDescriptor(input *GetDynamicClientRegistrationDescriptorInput) (output *models.ClientRegistrationPolicyDescriptor, resp *http.Response, err error)
	GetDynamicClientRegistrationDescriptorWithContext(ctx context.Context, input *GetDynamicClientRegistrationDescriptorInput) (output *models.ClientRegistrationPolicyDescriptor, resp *http.Response, err error)

	GetDynamicClientRegistrationPolicies() (output *models.ClientRegistrationPolicies, resp *http.Response, err error)
	GetDynamicClientRegistrationPoliciesWithContext(ctx context.Context) (output *models.ClientRegistrationPolicies, resp *http.Response, err error)

	CreateDynamicClientRegistrationPolicy(input *CreateDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)
	CreateDynamicClientRegistrationPolicyWithContext(ctx context.Context, input *CreateDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)

	GetDynamicClientRegistrationPolicy(input *GetDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)
	GetDynamicClientRegistrationPolicyWithContext(ctx context.Context, input *GetDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)

	UpdateDynamicClientRegistrationPolicy(input *UpdateDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)
	UpdateDynamicClientRegistrationPolicyWithContext(ctx context.Context, input *UpdateDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error)

	DeleteDynamicClientRegistrationPolicy(input *DeleteDynamicClientRegistrationPolicyInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteDynamicClientRegistrationPolicyWithContext(ctx context.Context, input *DeleteDynamicClientRegistrationPolicyInput) (output *models.ApiResult, resp *http.Response, err error)
}
