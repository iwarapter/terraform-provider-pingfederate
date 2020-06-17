package oauthTokenExchangeProcessor

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
	ServiceName = "OauthTokenExchangeProcessor"
)

type OauthTokenExchangeProcessorService struct {
	*client.PfClient
}

// New creates a new instance of the OauthTokenExchangeProcessorService client.
func New(cfg *config.Config) *OauthTokenExchangeProcessorService {

	return &OauthTokenExchangeProcessorService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthTokenExchangeProcessor operation
func (c *OauthTokenExchangeProcessorService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSettings - Get general OAuth 2.0 Token Exchange Processor settings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeProcessorService) GetSettings() (output *models.TokenExchangeProcessorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeProcessorSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Update general OAuth 2.0 Token Exchange Processor settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthTokenExchangeProcessorService) UpdateSettings(input *UpdateSettingsInput) (output *models.TokenExchangeProcessorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeProcessorSettings{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPolicies - Get list of OAuth 2.0 Token Exchange Processor policies.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeProcessorService) GetPolicies() (output *models.TokenExchangeProcessorPolicies, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies"
	op := &request.Operation{
		Name:       "GetPolicies",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeProcessorPolicies{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreatePolicy - Create a new OAuth 2.0 Token Exchange Processor policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthTokenExchangeProcessorService) CreatePolicy(input *CreatePolicyInput) (output *models.TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies"
	op := &request.Operation{
		Name:       "CreatePolicy",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeProcessorPolicy{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPolicy - Find an OAuth 2.0 Token Exchange Processor policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthTokenExchangeProcessorService) GetPolicy(input *GetPolicyInput) (output *models.TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeProcessorPolicy{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePolicy - Update an OAuth 2.0 Token Exchange Processor policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthTokenExchangeProcessorService) UpdatePolicy(input *UpdatePolicyInput) (output *models.TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdatePolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeProcessorPolicy{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeletePolicy - Delete an OAuth 2.0 Token Exchange Processor policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthTokenExchangeProcessorService) DeletePolicy(input *DeletePolicyInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
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
	Body models.TokenExchangeProcessorPolicy

	BypassExternalValidation *bool
}

type DeletePolicyInput struct {
	Id string
}

type GetPolicyInput struct {
	Id string
}

type UpdatePolicyInput struct {
	Body models.TokenExchangeProcessorPolicy
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.TokenExchangeProcessorSettings

	BypassExternalValidation *bool
}
