package spTokenGenerators

import (
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
	ServiceName = "SpTokenGenerators"
)

type SpTokenGeneratorsService struct {
	*client.PfClient
}

// New creates a new instance of the SpTokenGeneratorsService client.
func New(cfg *config.Config) *SpTokenGeneratorsService {

	return &SpTokenGeneratorsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a SpTokenGenerators operation
func (c *SpTokenGeneratorsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetTokenGeneratorDescriptors - Get the list of available token generators.
//RequestType: GET
//Input:
func (s *SpTokenGeneratorsService) GetTokenGeneratorDescriptors() (output *models.TokenGeneratorDescriptors, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/descriptors"
	op := &request.Operation{
		Name:       "GetTokenGeneratorDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenGeneratorDescriptors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenGeneratorDescriptorsById - Get the description of a token generator plugin by ID.
//RequestType: GET
//Input: input *GetTokenGeneratorDescriptorsByIdInput
func (s *SpTokenGeneratorsService) GetTokenGeneratorDescriptorsById(input *GetTokenGeneratorDescriptorsByIdInput) (output *models.TokenGeneratorDescriptor, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTokenGeneratorDescriptorsById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenGeneratorDescriptor{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenGenerators - Get the list of token generator instances.
//RequestType: GET
//Input:
func (s *SpTokenGeneratorsService) GetTokenGenerators() (output *models.TokenGenerators, resp *http.Response, err error) {
	path := "/sp/tokenGenerators"
	op := &request.Operation{
		Name:       "GetTokenGenerators",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenGenerators{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateTokenGenerator - Create a new token generator instance.
//RequestType: POST
//Input: input *CreateTokenGeneratorInput
func (s *SpTokenGeneratorsService) CreateTokenGenerator(input *CreateTokenGeneratorInput) (output *models.TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators"
	op := &request.Operation{
		Name:       "CreateTokenGenerator",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.TokenGenerator{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenGenerator - Find a token generator instance by ID.
//RequestType: GET
//Input: input *GetTokenGeneratorInput
func (s *SpTokenGeneratorsService) GetTokenGenerator(input *GetTokenGeneratorInput) (output *models.TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTokenGenerator",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenGenerator{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateTokenGenerator - Update a token generator instance.
//RequestType: PUT
//Input: input *UpdateTokenGeneratorInput
func (s *SpTokenGeneratorsService) UpdateTokenGenerator(input *UpdateTokenGeneratorInput) (output *models.TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateTokenGenerator",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.TokenGenerator{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteTokenGenerator - Delete a token generator instance.
//RequestType: DELETE
//Input: input *DeleteTokenGeneratorInput
func (s *SpTokenGeneratorsService) DeleteTokenGenerator(input *DeleteTokenGeneratorInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteTokenGenerator",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateTokenGeneratorInput struct {
	Body models.TokenGenerator
}

type DeleteTokenGeneratorInput struct {
	Id string
}

type GetTokenGeneratorInput struct {
	Id string
}

type GetTokenGeneratorDescriptorsByIdInput struct {
	Id string
}

type UpdateTokenGeneratorInput struct {
	Body models.TokenGenerator
	Id   string
}
