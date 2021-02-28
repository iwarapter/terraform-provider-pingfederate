package oauthAccessTokenManagers

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
	ServiceName = "OauthAccessTokenManagers"
)

type OauthAccessTokenManagersService struct {
	*client.PfClient
}

// New creates a new instance of the OauthAccessTokenManagersService client.
func New(cfg *config.Config) *OauthAccessTokenManagersService {

	return &OauthAccessTokenManagersService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthAccessTokenManagers operation
func (c *OauthAccessTokenManagersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSettings - Get general access token management settings.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetSettings() (output *models.AccessTokenManagementSettings, resp *http.Response, err error) {
	return s.GetSettingsWithContext(context.Background())
}

//GetSettingsWithContext - Get general access token management settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthAccessTokenManagersService) GetSettingsWithContext(ctx context.Context) (output *models.AccessTokenManagementSettings, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AccessTokenManagementSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Update general access token management settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthAccessTokenManagersService) UpdateSettings(input *UpdateSettingsInput) (output *models.AccessTokenManagementSettings, resp *http.Response, err error) {
	return s.UpdateSettingsWithContext(context.Background(), input)
}

//UpdateSettingsWithContext - Update general access token management settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSettingsInput
func (s *OauthAccessTokenManagersService) UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.AccessTokenManagementSettings, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AccessTokenManagementSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenManagerDescriptors - Get the list of available token management plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetTokenManagerDescriptors() (output *models.AccessTokenManagerDescriptors, resp *http.Response, err error) {
	return s.GetTokenManagerDescriptorsWithContext(context.Background())
}

//GetTokenManagerDescriptorsWithContext - Get the list of available token management plugin descriptors.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthAccessTokenManagersService) GetTokenManagerDescriptorsWithContext(ctx context.Context) (output *models.AccessTokenManagerDescriptors, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/descriptors"
	op := &request.Operation{
		Name:       "GetTokenManagerDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AccessTokenManagerDescriptors{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenManagerDescriptor - Get the description of a token management plugin descriptor.
//RequestType: GET
//Input: input *GetTokenManagerDescriptorInput
func (s *OauthAccessTokenManagersService) GetTokenManagerDescriptor(input *GetTokenManagerDescriptorInput) (output *models.AccessTokenManagerDescriptor, resp *http.Response, err error) {
	return s.GetTokenManagerDescriptorWithContext(context.Background(), input)
}

//GetTokenManagerDescriptorWithContext - Get the description of a token management plugin descriptor.
//RequestType: GET
//Input: ctx context.Context, input *GetTokenManagerDescriptorInput
func (s *OauthAccessTokenManagersService) GetTokenManagerDescriptorWithContext(ctx context.Context, input *GetTokenManagerDescriptorInput) (output *models.AccessTokenManagerDescriptor, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTokenManagerDescriptor",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AccessTokenManagerDescriptor{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenManagers - Get a list of all token management plugin instances.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetTokenManagers() (output *models.AccessTokenManagers, resp *http.Response, err error) {
	return s.GetTokenManagersWithContext(context.Background())
}

//GetTokenManagersWithContext - Get a list of all token management plugin instances.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthAccessTokenManagersService) GetTokenManagersWithContext(ctx context.Context) (output *models.AccessTokenManagers, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers"
	op := &request.Operation{
		Name:       "GetTokenManagers",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AccessTokenManagers{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateTokenManager - Create a token management plugin instance.
//RequestType: POST
//Input: input *CreateTokenManagerInput
func (s *OauthAccessTokenManagersService) CreateTokenManager(input *CreateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error) {
	return s.CreateTokenManagerWithContext(context.Background(), input)
}

//CreateTokenManagerWithContext - Create a token management plugin instance.
//RequestType: POST
//Input: ctx context.Context, input *CreateTokenManagerInput
func (s *OauthAccessTokenManagersService) CreateTokenManagerWithContext(ctx context.Context, input *CreateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers"
	op := &request.Operation{
		Name:       "CreateTokenManager",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AccessTokenManager{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetTokenManager - Get a specific token management plugin instance.
//RequestType: GET
//Input: input *GetTokenManagerInput
func (s *OauthAccessTokenManagersService) GetTokenManager(input *GetTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error) {
	return s.GetTokenManagerWithContext(context.Background(), input)
}

//GetTokenManagerWithContext - Get a specific token management plugin instance.
//RequestType: GET
//Input: ctx context.Context, input *GetTokenManagerInput
func (s *OauthAccessTokenManagersService) GetTokenManagerWithContext(ctx context.Context, input *GetTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTokenManager",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AccessTokenManager{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateTokenManager - Update a token management plugin instance.
//RequestType: PUT
//Input: input *UpdateTokenManagerInput
func (s *OauthAccessTokenManagersService) UpdateTokenManager(input *UpdateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error) {
	return s.UpdateTokenManagerWithContext(context.Background(), input)
}

//UpdateTokenManagerWithContext - Update a token management plugin instance.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateTokenManagerInput
func (s *OauthAccessTokenManagersService) UpdateTokenManagerWithContext(ctx context.Context, input *UpdateTokenManagerInput) (output *models.AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateTokenManager",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AccessTokenManager{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteTokenManager - Delete a token management plugin instance.
//RequestType: DELETE
//Input: input *DeleteTokenManagerInput
func (s *OauthAccessTokenManagersService) DeleteTokenManager(input *DeleteTokenManagerInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteTokenManagerWithContext(context.Background(), input)
}

//DeleteTokenManagerWithContext - Delete a token management plugin instance.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteTokenManagerInput
func (s *OauthAccessTokenManagersService) DeleteTokenManagerWithContext(ctx context.Context, input *DeleteTokenManagerInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteTokenManager",
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

type CreateTokenManagerInput struct {
	Body models.AccessTokenManager
}

type DeleteTokenManagerInput struct {
	Id string
}

type GetTokenManagerInput struct {
	Id string
}

type GetTokenManagerDescriptorInput struct {
	Id string
}

type UpdateSettingsInput struct {
	Body models.AccessTokenManagementSettings
}

type UpdateTokenManagerInput struct {
	Body models.AccessTokenManager
	Id   string
}
