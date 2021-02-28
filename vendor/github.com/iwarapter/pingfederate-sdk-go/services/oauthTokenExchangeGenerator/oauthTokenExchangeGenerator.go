package oauthTokenExchangeGenerator

import (
	"context"
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
	ServiceName = "OauthTokenExchangeGenerator"
)

type OauthTokenExchangeGeneratorService struct {
	*client.PfClient
}

// New creates a new instance of the OauthTokenExchangeGeneratorService client.
func New(cfg *config.Config) *OauthTokenExchangeGeneratorService {

	return &OauthTokenExchangeGeneratorService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthTokenExchangeGenerator operation
func (c *OauthTokenExchangeGeneratorService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSettings - Get general OAuth 2.0 Token Exchange Generator settings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeGeneratorService) GetSettings() (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error) {
	return s.GetSettingsWithContext(context.Background())
}

//GetSettingsWithContext - Get general OAuth 2.0 Token Exchange Generator settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthTokenExchangeGeneratorService) GetSettingsWithContext(ctx context.Context) (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeGeneratorSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Update general OAuth 2.0 Token Exchange Generator settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthTokenExchangeGeneratorService) UpdateSettings(input *UpdateSettingsInput) (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error) {
	return s.UpdateSettingsWithContext(context.Background(), input)
}

//UpdateSettingsWithContext - Update general OAuth 2.0 Token Exchange Generator settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSettingsInput
func (s *OauthTokenExchangeGeneratorService) UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeGeneratorSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetGroups - Get list of OAuth 2.0 Token Exchange Generator groups.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeGeneratorService) GetGroups() (output *models.TokenExchangeGeneratorGroups, resp *http.Response, err error) {
	return s.GetGroupsWithContext(context.Background())
}

//GetGroupsWithContext - Get list of OAuth 2.0 Token Exchange Generator groups.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthTokenExchangeGeneratorService) GetGroupsWithContext(ctx context.Context) (output *models.TokenExchangeGeneratorGroups, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups"
	op := &request.Operation{
		Name:       "GetGroups",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeGeneratorGroups{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateGroup - Create a new OAuth 2.0 Token Exchange Generator group.
//RequestType: POST
//Input: input *CreateGroupInput
func (s *OauthTokenExchangeGeneratorService) CreateGroup(input *CreateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	return s.CreateGroupWithContext(context.Background(), input)
}

//CreateGroupWithContext - Create a new OAuth 2.0 Token Exchange Generator group.
//RequestType: POST
//Input: ctx context.Context, input *CreateGroupInput
func (s *OauthTokenExchangeGeneratorService) CreateGroupWithContext(ctx context.Context, input *CreateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups"
	op := &request.Operation{
		Name:       "CreateGroup",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeGeneratorGroup{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetGroup - Find an OAuth 2.0 Token Exchange Generator group by ID.
//RequestType: GET
//Input: input *GetGroupInput
func (s *OauthTokenExchangeGeneratorService) GetGroup(input *GetGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	return s.GetGroupWithContext(context.Background(), input)
}

//GetGroupWithContext - Find an OAuth 2.0 Token Exchange Generator group by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetGroupInput
func (s *OauthTokenExchangeGeneratorService) GetGroupWithContext(ctx context.Context, input *GetGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetGroup",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeGeneratorGroup{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateGroup - Update an OAuth 2.0 Token Exchange Generator group.
//RequestType: PUT
//Input: input *UpdateGroupInput
func (s *OauthTokenExchangeGeneratorService) UpdateGroup(input *UpdateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	return s.UpdateGroupWithContext(context.Background(), input)
}

//UpdateGroupWithContext - Update an OAuth 2.0 Token Exchange Generator group.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateGroupInput
func (s *OauthTokenExchangeGeneratorService) UpdateGroupWithContext(ctx context.Context, input *UpdateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateGroup",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.TokenExchangeGeneratorGroup{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteGroup - Delete an OAuth 2.0 Token Exchange Generator group.
//RequestType: DELETE
//Input: input *DeleteGroupInput
func (s *OauthTokenExchangeGeneratorService) DeleteGroup(input *DeleteGroupInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteGroupWithContext(context.Background(), input)
}

//DeleteGroupWithContext - Delete an OAuth 2.0 Token Exchange Generator group.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteGroupInput
func (s *OauthTokenExchangeGeneratorService) DeleteGroupWithContext(ctx context.Context, input *DeleteGroupInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteGroup",
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

type CreateGroupInput struct {
	Body models.TokenExchangeGeneratorGroup

	BypassExternalValidation *bool
}

type DeleteGroupInput struct {
	Id string
}

type GetGroupInput struct {
	Id string
}

type UpdateGroupInput struct {
	Body models.TokenExchangeGeneratorGroup
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.TokenExchangeGeneratorSettings

	BypassExternalValidation *bool
}
