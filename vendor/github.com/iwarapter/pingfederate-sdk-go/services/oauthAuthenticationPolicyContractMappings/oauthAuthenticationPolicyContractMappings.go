package oauthAuthenticationPolicyContractMappings

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
	ServiceName = "OauthAuthenticationPolicyContractMappings"
)

type OauthAuthenticationPolicyContractMappingsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthAuthenticationPolicyContractMappingsService client.
func New(cfg *config.Config) *OauthAuthenticationPolicyContractMappingsService {

	return &OauthAuthenticationPolicyContractMappingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthAuthenticationPolicyContractMappings operation
func (c *OauthAuthenticationPolicyContractMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetApcMappings - Get the list of authentication policy contract to persistent grant mappings.
//RequestType: GET
//Input:
func (s *OauthAuthenticationPolicyContractMappingsService) GetApcMappings() (output *models.ApcToPersistentGrantMappings, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings"
	op := &request.Operation{
		Name:       "GetApcMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ApcToPersistentGrantMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateApcMapping - Create a new authentication policy contract to persistent grant mapping.
//RequestType: POST
//Input: input *CreateApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) CreateApcMapping(input *CreateApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings"
	op := &request.Operation{
		Name:       "CreateApcMapping",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ApcToPersistentGrantMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetApcMapping - Find the authentication policy contract to persistent grant mapping by ID.
//RequestType: GET
//Input: input *GetApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) GetApcMapping(input *GetApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetApcMapping",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ApcToPersistentGrantMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateApcMapping - Update an authentication policy contract to persistent grant mapping.
//RequestType: PUT
//Input: input *UpdateApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) UpdateApcMapping(input *UpdateApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateApcMapping",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ApcToPersistentGrantMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteApcMapping - Delete an authentication policy contract to persistent grant mapping.
//RequestType: DELETE
//Input: input *DeleteApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) DeleteApcMapping(input *DeleteApcMappingInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteApcMapping",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateApcMappingInput struct {
	Body models.ApcToPersistentGrantMapping

	BypassExternalValidation *bool
}

type DeleteApcMappingInput struct {
	Id string
}

type GetApcMappingInput struct {
	Id string
}

type UpdateApcMappingInput struct {
	Body models.ApcToPersistentGrantMapping
	Id   string

	BypassExternalValidation *bool
}
