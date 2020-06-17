package oauthIdpAdapterMappings

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
	ServiceName = "OauthIdpAdapterMappings"
)

type OauthIdpAdapterMappingsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthIdpAdapterMappingsService client.
func New(cfg *config.Config) *OauthIdpAdapterMappingsService {

	return &OauthIdpAdapterMappingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthIdpAdapterMappings operation
func (c *OauthIdpAdapterMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetIdpAdapterMappings - Get the list of IdP adapter mappings.
//RequestType: GET
//Input:
func (s *OauthIdpAdapterMappingsService) GetIdpAdapterMappings() (output *models.IdpAdapterMappings, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings"
	op := &request.Operation{
		Name:       "GetIdpAdapterMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpAdapterMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateIdpAdapterMapping - Create a new IdP adapter mapping.
//RequestType: POST
//Input: input *CreateIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) CreateIdpAdapterMapping(input *CreateIdpAdapterMappingInput) (output *models.IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings"
	op := &request.Operation{
		Name:       "CreateIdpAdapterMapping",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.IdpAdapterMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetIdpAdapterMapping - Find the IdP adapter mapping by the ID.
//RequestType: GET
//Input: input *GetIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) GetIdpAdapterMapping(input *GetIdpAdapterMappingInput) (output *models.IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetIdpAdapterMapping",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpAdapterMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateIdpAdapterMapping - Update an IdP adapter mapping.
//RequestType: PUT
//Input: input *UpdateIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) UpdateIdpAdapterMapping(input *UpdateIdpAdapterMappingInput) (output *models.IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateIdpAdapterMapping",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.IdpAdapterMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteIdpAdapterMapping - Delete an IdP adapter mapping.
//RequestType: DELETE
//Input: input *DeleteIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) DeleteIdpAdapterMapping(input *DeleteIdpAdapterMappingInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteIdpAdapterMapping",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateIdpAdapterMappingInput struct {
	Body models.IdpAdapterMapping

	BypassExternalValidation *bool
}

type DeleteIdpAdapterMappingInput struct {
	Id string
}

type GetIdpAdapterMappingInput struct {
	Id string
}

type UpdateIdpAdapterMappingInput struct {
	Body models.IdpAdapterMapping
	Id   string

	BypassExternalValidation *bool
}
