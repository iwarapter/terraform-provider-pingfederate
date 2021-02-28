package idpTokenProcessors

import (
	"context"
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
	ServiceName = "IdpTokenProcessors"
)

type IdpTokenProcessorsService struct {
	*client.PfClient
}

// New creates a new instance of the IdpTokenProcessorsService client.
func New(cfg *config.Config) *IdpTokenProcessorsService {

	return &IdpTokenProcessorsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IdpTokenProcessors operation
func (c *IdpTokenProcessorsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetTokenProcessorDescriptors - Get the list of available token processors.
//RequestType: GET
//Input:
func (s *IdpTokenProcessorsService) GetTokenProcessorDescriptors() (output *models.TokenProcessorDescriptors, resp *http.Response, err error) {
	return s.GetTokenProcessorDescriptorsWithContext(context.Background())
}

//GetTokenProcessorDescriptorsWithContext - Get the list of available token processors.
//RequestType: GET
//Input: ctx context.Context,
func (s *IdpTokenProcessorsService) GetTokenProcessorDescriptorsWithContext(ctx context.Context) (output *models.TokenProcessorDescriptors, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/descriptors"
	op := &request.Operation{
		Name:       "GetTokenProcessorDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenProcessorDescriptors{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenProcessorDescriptorsById - Get the description of a token processor plugin by ID.
//RequestType: GET
//Input: input *GetTokenProcessorDescriptorsByIdInput
func (s *IdpTokenProcessorsService) GetTokenProcessorDescriptorsById(input *GetTokenProcessorDescriptorsByIdInput) (output *models.TokenProcessorDescriptor, resp *http.Response, err error) {
	return s.GetTokenProcessorDescriptorsByIdWithContext(context.Background(), input)
}

//GetTokenProcessorDescriptorsByIdWithContext - Get the description of a token processor plugin by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetTokenProcessorDescriptorsByIdInput
func (s *IdpTokenProcessorsService) GetTokenProcessorDescriptorsByIdWithContext(ctx context.Context, input *GetTokenProcessorDescriptorsByIdInput) (output *models.TokenProcessorDescriptor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTokenProcessorDescriptorsById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenProcessorDescriptor{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenProcessors - Get the list of token processor instances.
//RequestType: GET
//Input:
func (s *IdpTokenProcessorsService) GetTokenProcessors() (output *models.TokenProcessors, resp *http.Response, err error) {
	return s.GetTokenProcessorsWithContext(context.Background())
}

//GetTokenProcessorsWithContext - Get the list of token processor instances.
//RequestType: GET
//Input: ctx context.Context,
func (s *IdpTokenProcessorsService) GetTokenProcessorsWithContext(ctx context.Context) (output *models.TokenProcessors, resp *http.Response, err error) {
	path := "/idp/tokenProcessors"
	op := &request.Operation{
		Name:       "GetTokenProcessors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenProcessors{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateTokenProcessor - Create a new token processor instance.
//RequestType: POST
//Input: input *CreateTokenProcessorInput
func (s *IdpTokenProcessorsService) CreateTokenProcessor(input *CreateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error) {
	return s.CreateTokenProcessorWithContext(context.Background(), input)
}

//CreateTokenProcessorWithContext - Create a new token processor instance.
//RequestType: POST
//Input: ctx context.Context, input *CreateTokenProcessorInput
func (s *IdpTokenProcessorsService) CreateTokenProcessorWithContext(ctx context.Context, input *CreateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors"
	op := &request.Operation{
		Name:       "CreateTokenProcessor",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.TokenProcessor{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenProcessor - Find a token processor instance by ID.
//RequestType: GET
//Input: input *GetTokenProcessorInput
func (s *IdpTokenProcessorsService) GetTokenProcessor(input *GetTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error) {
	return s.GetTokenProcessorWithContext(context.Background(), input)
}

//GetTokenProcessorWithContext - Find a token processor instance by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetTokenProcessorInput
func (s *IdpTokenProcessorsService) GetTokenProcessorWithContext(ctx context.Context, input *GetTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTokenProcessor",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenProcessor{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateTokenProcessor - Update a token processor instance.
//RequestType: PUT
//Input: input *UpdateTokenProcessorInput
func (s *IdpTokenProcessorsService) UpdateTokenProcessor(input *UpdateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error) {
	return s.UpdateTokenProcessorWithContext(context.Background(), input)
}

//UpdateTokenProcessorWithContext - Update a token processor instance.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateTokenProcessorInput
func (s *IdpTokenProcessorsService) UpdateTokenProcessorWithContext(ctx context.Context, input *UpdateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateTokenProcessor",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.TokenProcessor{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteTokenProcessor - Delete a token processor instance.
//RequestType: DELETE
//Input: input *DeleteTokenProcessorInput
func (s *IdpTokenProcessorsService) DeleteTokenProcessor(input *DeleteTokenProcessorInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteTokenProcessorWithContext(context.Background(), input)
}

//DeleteTokenProcessorWithContext - Delete a token processor instance.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteTokenProcessorInput
func (s *IdpTokenProcessorsService) DeleteTokenProcessorWithContext(ctx context.Context, input *DeleteTokenProcessorInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteTokenProcessor",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateTokenProcessorInput struct {
	Body models.TokenProcessor
}

type DeleteTokenProcessorInput struct {
	Id string
}

type GetTokenProcessorInput struct {
	Id string
}

type GetTokenProcessorDescriptorsByIdInput struct {
	Id string
}

type UpdateTokenProcessorInput struct {
	Body models.TokenProcessor
	Id   string
}
