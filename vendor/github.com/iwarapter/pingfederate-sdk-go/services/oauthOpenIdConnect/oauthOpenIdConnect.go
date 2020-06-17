package oauthOpenIdConnect

import (
	"fmt"
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
	ServiceName = "OauthOpenIdConnect"
)

type OauthOpenIdConnectService struct {
	*client.PfClient
}

// New creates a new instance of the OauthOpenIdConnectService client.
func New(cfg *config.Config) *OauthOpenIdConnectService {

	return &OauthOpenIdConnectService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthOpenIdConnect operation
func (c *OauthOpenIdConnectService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSettings - Get the OpenID Connect Settings.
//RequestType: GET
//Input:
func (s *OauthOpenIdConnectService) GetSettings() (output *models.OpenIdConnectSettings, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OpenIdConnectSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Set the OpenID Connect Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthOpenIdConnectService) UpdateSettings(input *UpdateSettingsInput) (output *models.OpenIdConnectSettings, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.OpenIdConnectSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPolicies - Get list of OpenID Connect Policies.
//RequestType: GET
//Input:
func (s *OauthOpenIdConnectService) GetPolicies() (output *models.OpenIdConnectPolicies, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies"
	op := &request.Operation{
		Name:       "GetPolicies",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OpenIdConnectPolicies{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreatePolicy - Create a new OpenID Connect Policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthOpenIdConnectService) CreatePolicy(input *CreatePolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies"
	op := &request.Operation{
		Name:       "CreatePolicy",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.OpenIdConnectPolicy{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPolicy - Find OpenID Connect Policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthOpenIdConnectService) GetPolicy(input *GetPolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OpenIdConnectPolicy{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePolicy - Update an OpenID Connect Policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthOpenIdConnectService) UpdatePolicy(input *UpdatePolicyInput) (output *models.OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdatePolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.OpenIdConnectPolicy{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeletePolicy - Delete an OpenID Connect Policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthOpenIdConnectService) DeletePolicy(input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeletePolicy",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreatePolicyInput struct {
	Body models.OpenIdConnectPolicy

	BypassExternalValidation *bool
}

type DeletePolicyInput struct {
	Id string
}

type GetPolicyInput struct {
	Id string
}

type UpdatePolicyInput struct {
	Body models.OpenIdConnectPolicy
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.OpenIdConnectSettings
}
