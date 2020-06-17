package oauthTokenExchangeTokenGeneratorMappings

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
	ServiceName = "OauthTokenExchangeTokenGeneratorMappings"
)

type OauthTokenExchangeTokenGeneratorMappingsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthTokenExchangeTokenGeneratorMappingsService client.
func New(cfg *config.Config) *OauthTokenExchangeTokenGeneratorMappingsService {

	return &OauthTokenExchangeTokenGeneratorMappingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthTokenExchangeTokenGeneratorMappings operation
func (c *OauthTokenExchangeTokenGeneratorMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetTokenGeneratorMappings - Get the list of Token Exchange Processor policy to Token Generator Mappings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeTokenGeneratorMappingsService) GetTokenGeneratorMappings() (output *models.ProcessorPolicyToGeneratorMappings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings"
	op := &request.Operation{
		Name:       "GetTokenGeneratorMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ProcessorPolicyToGeneratorMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateTokenGeneratorMapping - Create a new Token Exchange Processor policy to Token Generator Mapping.
//RequestType: POST
//Input: input *CreateTokenGeneratorMappingInput
func (s *OauthTokenExchangeTokenGeneratorMappingsService) CreateTokenGeneratorMapping(input *CreateTokenGeneratorMappingInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings"
	op := &request.Operation{
		Name:       "CreateTokenGeneratorMapping",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ProcessorPolicyToGeneratorMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenGeneratorMappingById - Get a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: GET
//Input: input *GetTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeTokenGeneratorMappingsService) GetTokenGeneratorMappingById(input *GetTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTokenGeneratorMappingById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ProcessorPolicyToGeneratorMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateTokenGeneratorMappingById - Update a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: PUT
//Input: input *UpdateTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeTokenGeneratorMappingsService) UpdateTokenGeneratorMappingById(input *UpdateTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateTokenGeneratorMappingById",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ProcessorPolicyToGeneratorMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteTokenGeneratorMappingById - Delete a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: DELETE
//Input: input *DeleteTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeTokenGeneratorMappingsService) DeleteTokenGeneratorMappingById(input *DeleteTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteTokenGeneratorMappingById",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}
	output = &models.ProcessorPolicyToGeneratorMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateTokenGeneratorMappingInput struct {
	Body models.ProcessorPolicyToGeneratorMapping

	BypassExternalValidation *bool
}

type DeleteTokenGeneratorMappingByIdInput struct {
	Id string
}

type GetTokenGeneratorMappingByIdInput struct {
	Id string
}

type UpdateTokenGeneratorMappingByIdInput struct {
	Body models.ProcessorPolicyToGeneratorMapping
	Id   string

	BypassExternalValidation *bool
}
