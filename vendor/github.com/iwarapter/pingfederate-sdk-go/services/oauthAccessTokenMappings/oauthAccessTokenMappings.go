package oauthAccessTokenMappings

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
	ServiceName = "OauthAccessTokenMappings"
)

type OauthAccessTokenMappingsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthAccessTokenMappingsService client.
func New(cfg *config.Config) *OauthAccessTokenMappingsService {

	return &OauthAccessTokenMappingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthAccessTokenMappings operation
func (c *OauthAccessTokenMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetMappings - Get the list of Access Token Mappings.
//RequestType: GET
//Input:
func (s *OauthAccessTokenMappingsService) GetMappings() (output *models.AccessTokenMappings, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings"
	op := &request.Operation{
		Name:       "GetMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AccessTokenMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateMapping - Create a new Access Token Mapping.
//RequestType: POST
//Input: input *CreateMappingInput
func (s *OauthAccessTokenMappingsService) CreateMapping(input *CreateMappingInput) (output *models.AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings"
	op := &request.Operation{
		Name:       "CreateMapping",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AccessTokenMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetMapping - Find the Access Token Mapping by its ID.
//RequestType: GET
//Input: input *GetMappingInput
func (s *OauthAccessTokenMappingsService) GetMapping(input *GetMappingInput) (output *models.AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetMapping",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AccessTokenMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateMapping - Update an Access Token Mapping.
//RequestType: PUT
//Input: input *UpdateMappingInput
func (s *OauthAccessTokenMappingsService) UpdateMapping(input *UpdateMappingInput) (output *models.AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateMapping",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AccessTokenMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteMapping - Delete an Access Token Mapping.
//RequestType: DELETE
//Input: input *DeleteMappingInput
func (s *OauthAccessTokenMappingsService) DeleteMapping(input *DeleteMappingInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteMapping",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateMappingInput struct {
	Body models.AccessTokenMapping

	BypassExternalValidation *bool
}

type DeleteMappingInput struct {
	Id string
}

type GetMappingInput struct {
	Id string
}

type UpdateMappingInput struct {
	Body models.AccessTokenMapping
	Id   string

	BypassExternalValidation *bool
}
