package authenticationApi

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
	ServiceName = "AuthenticationApi"
)

type AuthenticationApiService struct {
	*client.PfClient
}

// New creates a new instance of the AuthenticationApiService client.
func New(cfg *config.Config) *AuthenticationApiService {

	return &AuthenticationApiService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthenticationApi operation
func (c *AuthenticationApiService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetAuthenticationApiSettings - Get the Authentication API settings.
//RequestType: GET
//Input:
func (s *AuthenticationApiService) GetAuthenticationApiSettings() (output *models.AuthnApiSettings, resp *http.Response, err error) {
	path := "/authenticationApi/settings"
	op := &request.Operation{
		Name:       "GetAuthenticationApiSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthnApiSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAuthenticationApiSettings - Set the Authentication API settings.
//RequestType: PUT
//Input: input *UpdateAuthenticationApiSettingsInput
func (s *AuthenticationApiService) UpdateAuthenticationApiSettings(input *UpdateAuthenticationApiSettingsInput) (output *models.AuthnApiSettings, resp *http.Response, err error) {
	path := "/authenticationApi/settings"
	op := &request.Operation{
		Name:       "UpdateAuthenticationApiSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthnApiSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAuthenticationApiApplications - Get the collection of Authentication API Applications.
//RequestType: GET
//Input:
func (s *AuthenticationApiService) GetAuthenticationApiApplications() (output *models.AuthnApiApplications, resp *http.Response, err error) {
	path := "/authenticationApi/applications"
	op := &request.Operation{
		Name:       "GetAuthenticationApiApplications",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthnApiApplications{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateApplication - Create a new Authentication API Application.
//RequestType: POST
//Input: input *CreateApplicationInput
func (s *AuthenticationApiService) CreateApplication(input *CreateApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications"
	op := &request.Operation{
		Name:       "CreateApplication",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AuthnApiApplication{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetApplication - Find Authentication API Application by ID.
//RequestType: GET
//Input: input *GetApplicationInput
func (s *AuthenticationApiService) GetApplication(input *GetApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetApplication",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthnApiApplication{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateApplication - Update an Authentication API Application.
//RequestType: PUT
//Input: input *UpdateApplicationInput
func (s *AuthenticationApiService) UpdateApplication(input *UpdateApplicationInput) (output *models.AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateApplication",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthnApiApplication{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteApplication - Delete an Authentication API Application.
//RequestType: DELETE
//Input: input *DeleteApplicationInput
func (s *AuthenticationApiService) DeleteApplication(input *DeleteApplicationInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteApplication",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateApplicationInput struct {
	Body models.AuthnApiApplication
}

type DeleteApplicationInput struct {
	Id string
}

type GetApplicationInput struct {
	Id string
}

type UpdateApplicationInput struct {
	Body models.AuthnApiApplication
	Id   string
}

type UpdateAuthenticationApiSettingsInput struct {
	Body models.AuthnApiSettings
}
