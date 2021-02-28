package passwordCredentialValidators

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
	ServiceName = "PasswordCredentialValidators"
)

type PasswordCredentialValidatorsService struct {
	*client.PfClient
}

// New creates a new instance of the PasswordCredentialValidatorsService client.
func New(cfg *config.Config) *PasswordCredentialValidatorsService {

	return &PasswordCredentialValidatorsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a PasswordCredentialValidators operation
func (c *PasswordCredentialValidatorsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetPasswordCredentialValidatorDescriptors - Get a list of available password credential validator descriptors.
//RequestType: GET
//Input:
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorDescriptors() (output *models.PasswordCredentialValidatorDescriptors, resp *http.Response, err error) {
	return s.GetPasswordCredentialValidatorDescriptorsWithContext(context.Background())
}

//GetPasswordCredentialValidatorDescriptorsWithContext - Get a list of available password credential validator descriptors.
//RequestType: GET
//Input: ctx context.Context,
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorDescriptorsWithContext(ctx context.Context) (output *models.PasswordCredentialValidatorDescriptors, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/descriptors"
	op := &request.Operation{
		Name:       "GetPasswordCredentialValidatorDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PasswordCredentialValidatorDescriptors{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPasswordCredentialValidatorDescriptor - Get the description of a password credential validator by ID.
//RequestType: GET
//Input: input *GetPasswordCredentialValidatorDescriptorInput
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorDescriptor(input *GetPasswordCredentialValidatorDescriptorInput) (output *models.PasswordCredentialValidatorDescriptor, resp *http.Response, err error) {
	return s.GetPasswordCredentialValidatorDescriptorWithContext(context.Background(), input)
}

//GetPasswordCredentialValidatorDescriptorWithContext - Get the description of a password credential validator by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetPasswordCredentialValidatorDescriptorInput
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorDescriptorWithContext(ctx context.Context, input *GetPasswordCredentialValidatorDescriptorInput) (output *models.PasswordCredentialValidatorDescriptor, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPasswordCredentialValidatorDescriptor",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PasswordCredentialValidatorDescriptor{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPasswordCredentialValidators - Get the list of available password credential validators
//RequestType: GET
//Input:
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidators() (output *models.PasswordCredentialValidators, resp *http.Response, err error) {
	return s.GetPasswordCredentialValidatorsWithContext(context.Background())
}

//GetPasswordCredentialValidatorsWithContext - Get the list of available password credential validators
//RequestType: GET
//Input: ctx context.Context,
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorsWithContext(ctx context.Context) (output *models.PasswordCredentialValidators, resp *http.Response, err error) {
	path := "/passwordCredentialValidators"
	op := &request.Operation{
		Name:       "GetPasswordCredentialValidators",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PasswordCredentialValidators{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreatePasswordCredentialValidator - Create a new password credential validator instance
//RequestType: POST
//Input: input *CreatePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) CreatePasswordCredentialValidator(input *CreatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error) {
	return s.CreatePasswordCredentialValidatorWithContext(context.Background(), input)
}

//CreatePasswordCredentialValidatorWithContext - Create a new password credential validator instance
//RequestType: POST
//Input: ctx context.Context, input *CreatePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) CreatePasswordCredentialValidatorWithContext(ctx context.Context, input *CreatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators"
	op := &request.Operation{
		Name:       "CreatePasswordCredentialValidator",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.PasswordCredentialValidator{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPasswordCredentialValidator - Find a password credential validator by ID.
//RequestType: GET
//Input: input *GetPasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidator(input *GetPasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error) {
	return s.GetPasswordCredentialValidatorWithContext(context.Background(), input)
}

//GetPasswordCredentialValidatorWithContext - Find a password credential validator by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetPasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorWithContext(ctx context.Context, input *GetPasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPasswordCredentialValidator",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PasswordCredentialValidator{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePasswordCredentialValidator - Update a password credential validator instance.
//RequestType: PUT
//Input: input *UpdatePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) UpdatePasswordCredentialValidator(input *UpdatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error) {
	return s.UpdatePasswordCredentialValidatorWithContext(context.Background(), input)
}

//UpdatePasswordCredentialValidatorWithContext - Update a password credential validator instance.
//RequestType: PUT
//Input: ctx context.Context, input *UpdatePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) UpdatePasswordCredentialValidatorWithContext(ctx context.Context, input *UpdatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdatePasswordCredentialValidator",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.PasswordCredentialValidator{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeletePasswordCredentialValidator - Delete a password credential validator instance.
//RequestType: DELETE
//Input: input *DeletePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) DeletePasswordCredentialValidator(input *DeletePasswordCredentialValidatorInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeletePasswordCredentialValidatorWithContext(context.Background(), input)
}

//DeletePasswordCredentialValidatorWithContext - Delete a password credential validator instance.
//RequestType: DELETE
//Input: ctx context.Context, input *DeletePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) DeletePasswordCredentialValidatorWithContext(ctx context.Context, input *DeletePasswordCredentialValidatorInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeletePasswordCredentialValidator",
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

type CreatePasswordCredentialValidatorInput struct {
	Body models.PasswordCredentialValidator
}

type DeletePasswordCredentialValidatorInput struct {
	Id string
}

type GetPasswordCredentialValidatorInput struct {
	Id string
}

type GetPasswordCredentialValidatorDescriptorInput struct {
	Id string
}

type UpdatePasswordCredentialValidatorInput struct {
	Body models.PasswordCredentialValidator
	Id   string
}
