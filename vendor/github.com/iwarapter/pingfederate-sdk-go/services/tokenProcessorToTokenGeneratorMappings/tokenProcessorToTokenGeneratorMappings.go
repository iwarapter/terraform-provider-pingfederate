package tokenProcessorToTokenGeneratorMappings

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
	ServiceName = "TokenProcessorToTokenGeneratorMappings"
)

type TokenProcessorToTokenGeneratorMappingsService struct {
	*client.PfClient
}

// New creates a new instance of the TokenProcessorToTokenGeneratorMappingsService client.
func New(cfg *config.Config) *TokenProcessorToTokenGeneratorMappingsService {

	return &TokenProcessorToTokenGeneratorMappingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a TokenProcessorToTokenGeneratorMappings operation
func (c *TokenProcessorToTokenGeneratorMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetTokenToTokenMappings - Get the list of Token Processor to Token Generator Mappings.
//RequestType: GET
//Input:
func (s *TokenProcessorToTokenGeneratorMappingsService) GetTokenToTokenMappings() (output *models.TokenToTokenMappings, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings"
	op := &request.Operation{
		Name:       "GetTokenToTokenMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenToTokenMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateTokenToTokenMapping - Create a new Token Processor to Token Generator Mapping.
//RequestType: POST
//Input: input *CreateTokenToTokenMappingInput
func (s *TokenProcessorToTokenGeneratorMappingsService) CreateTokenToTokenMapping(input *CreateTokenToTokenMappingInput) (output *models.TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings"
	op := &request.Operation{
		Name:       "CreateTokenToTokenMapping",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.TokenToTokenMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenToTokenMappingById - Get a Token Processor to Token Generator Mapping.
//RequestType: GET
//Input: input *GetTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) GetTokenToTokenMappingById(input *GetTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTokenToTokenMappingById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenToTokenMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateTokenToTokenMappingById - Update a Token Processor to Token Generator Mapping.
//RequestType: PUT
//Input: input *UpdateTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) UpdateTokenToTokenMappingById(input *UpdateTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateTokenToTokenMappingById",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.TokenToTokenMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteTokenToTokenMappingById - Delete a Token Processor to Token Generator Mapping.
//RequestType: DELETE
//Input: input *DeleteTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) DeleteTokenToTokenMappingById(input *DeleteTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteTokenToTokenMappingById",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}
	output = &models.TokenToTokenMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateTokenToTokenMappingInput struct {
	Body models.TokenToTokenMapping

	BypassExternalValidation *bool
}

type DeleteTokenToTokenMappingByIdInput struct {
	Id string
}

type GetTokenToTokenMappingByIdInput struct {
	Id string
}

type UpdateTokenToTokenMappingByIdInput struct {
	Body models.TokenToTokenMapping
	Id   string

	BypassExternalValidation *bool
}
