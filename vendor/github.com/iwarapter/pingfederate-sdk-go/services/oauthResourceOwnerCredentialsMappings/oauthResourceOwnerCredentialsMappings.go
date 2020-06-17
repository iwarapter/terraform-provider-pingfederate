package oauthResourceOwnerCredentialsMappings

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
	ServiceName = "OauthResourceOwnerCredentialsMappings"
)

type OauthResourceOwnerCredentialsMappingsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthResourceOwnerCredentialsMappingsService client.
func New(cfg *config.Config) *OauthResourceOwnerCredentialsMappingsService {

	return &OauthResourceOwnerCredentialsMappingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthResourceOwnerCredentialsMappings operation
func (c *OauthResourceOwnerCredentialsMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetResourceOwnerCredentialsMappings - Get the list of Resource Owner Credentials mappings.
//RequestType: GET
//Input:
func (s *OauthResourceOwnerCredentialsMappingsService) GetResourceOwnerCredentialsMappings() (output *models.ResourceOwnerCredentialsMappings, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings"
	op := &request.Operation{
		Name:       "GetResourceOwnerCredentialsMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ResourceOwnerCredentialsMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateResourceOwnerCredentialsMapping - Create a new Resource Owner Credentials mapping.
//RequestType: POST
//Input: input *CreateResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) CreateResourceOwnerCredentialsMapping(input *CreateResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings"
	op := &request.Operation{
		Name:       "CreateResourceOwnerCredentialsMapping",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ResourceOwnerCredentialsMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetResourceOwnerCredentialsMapping - Find the Resource Owner Credentials mapping by the ID.
//RequestType: GET
//Input: input *GetResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) GetResourceOwnerCredentialsMapping(input *GetResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetResourceOwnerCredentialsMapping",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ResourceOwnerCredentialsMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateResourceOwnerCredentialsMapping - Update a Resource Owner Credentials mapping.
//RequestType: PUT
//Input: input *UpdateResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) UpdateResourceOwnerCredentialsMapping(input *UpdateResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateResourceOwnerCredentialsMapping",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ResourceOwnerCredentialsMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteResourceOwnerCredentialsMapping - Delete a Resource Owner Credentials mapping.
//RequestType: DELETE
//Input: input *DeleteResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) DeleteResourceOwnerCredentialsMapping(input *DeleteResourceOwnerCredentialsMappingInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteResourceOwnerCredentialsMapping",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateResourceOwnerCredentialsMappingInput struct {
	Body models.ResourceOwnerCredentialsMapping

	BypassExternalValidation *bool
}

type DeleteResourceOwnerCredentialsMappingInput struct {
	Id string
}

type GetResourceOwnerCredentialsMappingInput struct {
	Id string
}

type UpdateResourceOwnerCredentialsMappingInput struct {
	Body models.ResourceOwnerCredentialsMapping
	Id   string

	BypassExternalValidation *bool
}
