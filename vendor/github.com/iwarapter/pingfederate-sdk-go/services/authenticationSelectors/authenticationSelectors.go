package authenticationSelectors

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
	ServiceName = "AuthenticationSelectors"
)

type AuthenticationSelectorsService struct {
	*client.PfClient
}

// New creates a new instance of the AuthenticationSelectorsService client.
func New(cfg *config.Config) *AuthenticationSelectorsService {

	return &AuthenticationSelectorsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthenticationSelectors operation
func (c *AuthenticationSelectorsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetAuthenticationSelectorDescriptors - Get the list of available Authentication Selector descriptors.
//RequestType: GET
//Input:
func (s *AuthenticationSelectorsService) GetAuthenticationSelectorDescriptors() (output *models.AuthenticationSelectorDescriptors, resp *http.Response, err error) {
	path := "/authenticationSelectors/descriptors"
	op := &request.Operation{
		Name:       "GetAuthenticationSelectorDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSelectorDescriptors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAuthenticationSelectorDescriptorsById - Get the description of an Authentication Selector plugin by ID.
//RequestType: GET
//Input: input *GetAuthenticationSelectorDescriptorsByIdInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelectorDescriptorsById(input *GetAuthenticationSelectorDescriptorsByIdInput) (output *models.AuthenticationSelectorDescriptor, resp *http.Response, err error) {
	path := "/authenticationSelectors/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetAuthenticationSelectorDescriptorsById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSelectorDescriptor{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAuthenticationSelectors - Get the list of configured Authentication Selector instances.
//RequestType: GET
//Input: input *GetAuthenticationSelectorsInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelectors(input *GetAuthenticationSelectorsInput) (output *models.AuthenticationSelectors, resp *http.Response, err error) {
	path := "/authenticationSelectors"
	op := &request.Operation{
		Name:       "GetAuthenticationSelectors",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
		},
	}
	output = &models.AuthenticationSelectors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateAuthenticationSelector - Create a new authentication selector instance.
//RequestType: POST
//Input: input *CreateAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) CreateAuthenticationSelector(input *CreateAuthenticationSelectorInput) (output *models.AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors"
	op := &request.Operation{
		Name:       "CreateAuthenticationSelector",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSelector{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAuthenticationSelector - Get an Authentication Selector instance by ID.
//RequestType: GET
//Input: input *GetAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelector(input *GetAuthenticationSelectorInput) (output *models.AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetAuthenticationSelector",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSelector{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAuthenticationSelector - Update an authentication selector instance.
//RequestType: PUT
//Input: input *UpdateAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) UpdateAuthenticationSelector(input *UpdateAuthenticationSelectorInput) (output *models.AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateAuthenticationSelector",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthenticationSelector{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteAuthenticationSelector - Delete an Authentication Selector instance.
//RequestType: DELETE
//Input: input *DeleteAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) DeleteAuthenticationSelector(input *DeleteAuthenticationSelectorInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteAuthenticationSelector",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateAuthenticationSelectorInput struct {
	Body models.AuthenticationSelector
}

type DeleteAuthenticationSelectorInput struct {
	Id string
}

type GetAuthenticationSelectorInput struct {
	Id string
}

type GetAuthenticationSelectorDescriptorsByIdInput struct {
	Id string
}

type GetAuthenticationSelectorsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type UpdateAuthenticationSelectorInput struct {
	Body models.AuthenticationSelector
	Id   string
}
