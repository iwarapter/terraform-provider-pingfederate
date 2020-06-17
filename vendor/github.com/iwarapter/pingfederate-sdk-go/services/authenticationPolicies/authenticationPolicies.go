package authenticationPolicies

import (
	"fmt"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "AuthenticationPolicies"
)

type AuthenticationPoliciesService struct {
	*client.PfClient
}

// New creates a new instance of the AuthenticationPoliciesService client.
func New(cfg *config.Config) *AuthenticationPoliciesService {

	return &AuthenticationPoliciesService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthenticationPolicies operation
func (c *AuthenticationPoliciesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSettings - Get the authentication policies settings.
//RequestType: GET
//Input:
func (s *AuthenticationPoliciesService) GetSettings() (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error) {
	path := "/authenticationPolicies/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPoliciesSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Set the authentication policies settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *AuthenticationPoliciesService) UpdateSettings(input *UpdateSettingsInput) (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error) {
	path := "/authenticationPolicies/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPoliciesSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDefaultAuthenticationPolicy - Get the default configured authentication policy.
//RequestType: GET
//Input:
func (s *AuthenticationPoliciesService) GetDefaultAuthenticationPolicy() (output *models.AuthenticationPolicy, resp *http.Response, err error) {
	path := "/authenticationPolicies/default"
	op := &request.Operation{
		Name:       "GetDefaultAuthenticationPolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicy{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateDefaultAuthenticationPolicy - Set the default authentication policy.
//RequestType: PUT
//Input: input *UpdateDefaultAuthenticationPolicyInput
func (s *AuthenticationPoliciesService) UpdateDefaultAuthenticationPolicy(input *UpdateDefaultAuthenticationPolicyInput) (output *models.AuthenticationPolicy, resp *http.Response, err error) {
	path := "/authenticationPolicies/default"
	op := &request.Operation{
		Name:       "UpdateDefaultAuthenticationPolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicy{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateDefaultAuthenticationPolicyInput struct {
	Body models.AuthenticationPolicy

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.AuthenticationPoliciesSettings
}
