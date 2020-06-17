package idpToSpAdapterMapping

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
	ServiceName = "IdpToSpAdapterMapping"
)

type IdpToSpAdapterMappingService struct {
	*client.PfClient
}

// New creates a new instance of the IdpToSpAdapterMappingService client.
func New(cfg *config.Config) *IdpToSpAdapterMappingService {

	return &IdpToSpAdapterMappingService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IdpToSpAdapterMapping operation
func (c *IdpToSpAdapterMappingService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetIdpToSpAdapterMappings - Get list of IdP-to-SP Adapter Mappings.
//RequestType: GET
//Input:
func (s *IdpToSpAdapterMappingService) GetIdpToSpAdapterMappings() (output *models.IdpToSpAdapterMappings, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping"
	op := &request.Operation{
		Name:       "GetIdpToSpAdapterMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpToSpAdapterMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateIdpToSpAdapterMapping - Create a new IdP-to-SP Adapter mapping.
//RequestType: POST
//Input: input *CreateIdpToSpAdapterMappingInput
func (s *IdpToSpAdapterMappingService) CreateIdpToSpAdapterMapping(input *CreateIdpToSpAdapterMappingInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping"
	op := &request.Operation{
		Name:       "CreateIdpToSpAdapterMapping",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.IdpToSpAdapterMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetIdpToSpAdapterMappingsById - Get an IdP-to-SP Adapter Mapping.
//RequestType: GET
//Input: input *GetIdpToSpAdapterMappingsByIdInput
func (s *IdpToSpAdapterMappingService) GetIdpToSpAdapterMappingsById(input *GetIdpToSpAdapterMappingsByIdInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetIdpToSpAdapterMappingsById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpToSpAdapterMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateIdpToSpAdapterMapping - Update the specified IdP-to-SP Adapter mapping.
//RequestType: PUT
//Input: input *UpdateIdpToSpAdapterMappingInput
func (s *IdpToSpAdapterMappingService) UpdateIdpToSpAdapterMapping(input *UpdateIdpToSpAdapterMappingInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateIdpToSpAdapterMapping",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.IdpToSpAdapterMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteIdpToSpAdapterMappingsById - Delete an Adapter to Adapter Mapping.
//RequestType: DELETE
//Input: input *DeleteIdpToSpAdapterMappingsByIdInput
func (s *IdpToSpAdapterMappingService) DeleteIdpToSpAdapterMappingsById(input *DeleteIdpToSpAdapterMappingsByIdInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteIdpToSpAdapterMappingsById",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateIdpToSpAdapterMappingInput struct {
	Body models.IdpToSpAdapterMapping

	BypassExternalValidation *bool
}

type DeleteIdpToSpAdapterMappingsByIdInput struct {
	Id string
}

type GetIdpToSpAdapterMappingsByIdInput struct {
	Id string
}

type UpdateIdpToSpAdapterMappingInput struct {
	Body models.IdpToSpAdapterMapping
	Id   string

	BypassExternalValidation *bool
}
