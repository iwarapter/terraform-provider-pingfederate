package oauthClientRegistrationPolicies

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "OauthClientRegistrationPolicies"
)

type OauthClientRegistrationPoliciesService struct {
	*client.PfClient
}

// New creates a new instance of the OauthClientRegistrationPoliciesService client.
func New(cfg *config.Config) *OauthClientRegistrationPoliciesService {

	return &OauthClientRegistrationPoliciesService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthClientRegistrationPolicies operation
func (c *OauthClientRegistrationPoliciesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetDynamicClientRegistrationDescriptors - Get the list of available client registration policy plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationDescriptors() (output *models.ClientRegistrationPolicyDescriptors, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/descriptors"
	op := &request.Operation{
		Name:       "GetDynamicClientRegistrationDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ClientRegistrationPolicyDescriptors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDynamicClientRegistrationDescriptor - Get the description of a client registration policy plugin descriptor.
//RequestType: GET
//Input: input *GetDynamicClientRegistrationDescriptorInput
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationDescriptor(input *GetDynamicClientRegistrationDescriptorInput) (output *models.ClientRegistrationPolicyDescriptor, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetDynamicClientRegistrationDescriptor",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ClientRegistrationPolicyDescriptor{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDynamicClientRegistrationPolicies - Get a list of client registration policy plugin instances.
//RequestType: GET
//Input:
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationPolicies() (output *models.ClientRegistrationPolicies, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies"
	op := &request.Operation{
		Name:       "GetDynamicClientRegistrationPolicies",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ClientRegistrationPolicies{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateDynamicClientRegistrationPolicy - Create a client registration policy plugin instance.
//RequestType: POST
//Input: input *CreateDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) CreateDynamicClientRegistrationPolicy(input *CreateDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies"
	op := &request.Operation{
		Name:       "CreateDynamicClientRegistrationPolicy",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ClientRegistrationPolicy{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDynamicClientRegistrationPolicy - Get a specific client registration policy plugin instance.
//RequestType: GET
//Input: input *GetDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationPolicy(input *GetDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetDynamicClientRegistrationPolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ClientRegistrationPolicy{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateDynamicClientRegistrationPolicy - Update a client registration policy plugin instance.
//RequestType: PUT
//Input: input *UpdateDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) UpdateDynamicClientRegistrationPolicy(input *UpdateDynamicClientRegistrationPolicyInput) (output *models.ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateDynamicClientRegistrationPolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ClientRegistrationPolicy{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteDynamicClientRegistrationPolicy - Delete a client registration policy plugin instance.
//RequestType: DELETE
//Input: input *DeleteDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) DeleteDynamicClientRegistrationPolicy(input *DeleteDynamicClientRegistrationPolicyInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteDynamicClientRegistrationPolicy",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateDynamicClientRegistrationPolicyInput struct {
	Body models.ClientRegistrationPolicy
}

type DeleteDynamicClientRegistrationPolicyInput struct {
	Id string
}

type GetDynamicClientRegistrationDescriptorInput struct {
	Id string
}

type GetDynamicClientRegistrationPolicyInput struct {
	Id string
}

type UpdateDynamicClientRegistrationPolicyInput struct {
	Body models.ClientRegistrationPolicy
	Id   string
}
