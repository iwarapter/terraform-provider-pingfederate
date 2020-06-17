package oauthAuthServerSettings

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
	ServiceName = "OauthAuthServerSettings"
)

type OauthAuthServerSettingsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthAuthServerSettingsService client.
func New(cfg *config.Config) *OauthAuthServerSettingsService {

	return &OauthAuthServerSettingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthAuthServerSettings operation
func (c *OauthAuthServerSettingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetAuthorizationServerSettings - Get the Authorization Server Settings.
//RequestType: GET
//Input:
func (s *OauthAuthServerSettingsService) GetAuthorizationServerSettings() (output *models.AuthorizationServerSettings, resp *http.Response, err error) {
	path := "/oauth/authServerSettings"
	op := &request.Operation{
		Name:       "GetAuthorizationServerSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthorizationServerSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAuthorizationServerSettings - Update the Authorization Server Settings.
//RequestType: PUT
//Input: input *UpdateAuthorizationServerSettingsInput
func (s *OauthAuthServerSettingsService) UpdateAuthorizationServerSettings(input *UpdateAuthorizationServerSettingsInput) (output *models.AuthorizationServerSettings, resp *http.Response, err error) {
	path := "/oauth/authServerSettings"
	op := &request.Operation{
		Name:       "UpdateAuthorizationServerSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthorizationServerSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddCommonScope - Add a new common scope.
//RequestType: POST
//Input: input *AddCommonScopeInput
func (s *OauthAuthServerSettingsService) AddCommonScope(input *AddCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes"
	op := &request.Operation{
		Name:       "AddCommonScope",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ScopeEntry{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetCommonScope - Get an existing common scope.
//RequestType: GET
//Input: input *GetCommonScopeInput
func (s *OauthAuthServerSettingsService) GetCommonScope(input *GetCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "GetCommonScope",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ScopeEntry{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateCommonScope - Update an existing common scope.
//RequestType: PUT
//Input: input *UpdateCommonScopeInput
func (s *OauthAuthServerSettingsService) UpdateCommonScope(input *UpdateCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "UpdateCommonScope",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ScopeEntry{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//RemoveCommonScope - Remove an existing common scope.
//RequestType: DELETE
//Input: input *RemoveCommonScopeInput
func (s *OauthAuthServerSettingsService) RemoveCommonScope(input *RemoveCommonScopeInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "RemoveCommonScope",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddCommonScopeGroup - Create a new common scope group.
//RequestType: POST
//Input: input *AddCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) AddCommonScopeGroup(input *AddCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups"
	op := &request.Operation{
		Name:       "AddCommonScopeGroup",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ScopeGroupEntry{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetCommonScopeGroup - Get an existing common scope group.
//RequestType: GET
//Input: input *GetCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) GetCommonScopeGroup(input *GetCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "GetCommonScopeGroup",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ScopeGroupEntry{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateCommonScopeGroup - Update an existing common scope group.
//RequestType: PUT
//Input: input *UpdateCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) UpdateCommonScopeGroup(input *UpdateCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "UpdateCommonScopeGroup",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ScopeGroupEntry{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//RemoveCommonScopeGroup - Remove an existing common scope group.
//RequestType: DELETE
//Input: input *RemoveCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) RemoveCommonScopeGroup(input *RemoveCommonScopeGroupInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "RemoveCommonScopeGroup",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddExclusiveScope - Add a new exclusive scope.
//RequestType: POST
//Input: input *AddExclusiveScopeInput
func (s *OauthAuthServerSettingsService) AddExclusiveScope(input *AddExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes"
	op := &request.Operation{
		Name:       "AddExclusiveScope",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ScopeEntry{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetExclusiveScope - Get an existing exclusive scope.
//RequestType: GET
//Input: input *GetExclusiveScopeInput
func (s *OauthAuthServerSettingsService) GetExclusiveScope(input *GetExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "GetExclusiveScope",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ScopeEntry{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateExclusiveScope - Update an existing exclusive scope.
//RequestType: PUT
//Input: input *UpdateExclusiveScopeInput
func (s *OauthAuthServerSettingsService) UpdateExclusiveScope(input *UpdateExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "UpdateExclusiveScope",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ScopeEntry{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//RemoveExclusiveScope - Remove an existing exclusive scope.
//RequestType: DELETE
//Input: input *RemoveExclusiveScopeInput
func (s *OauthAuthServerSettingsService) RemoveExclusiveScope(input *RemoveExclusiveScopeInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "RemoveExclusiveScope",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddExclusiveScopeGroup - Create a new exclusive scope group.
//RequestType: POST
//Input: input *AddExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) AddExclusiveScopeGroup(input *AddExclusiveScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups"
	op := &request.Operation{
		Name:       "AddExclusiveScopeGroup",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ScopeGroupEntry{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetExclusiveScopeGroup - Get an existing exclusive scope group.
//RequestType: GET
//Input: input *GetExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) GetExclusiveScopeGroup(input *GetExclusiveScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "GetExclusiveScopeGroup",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ScopeGroupEntry{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateExclusiveScopeGroups - Update an existing exclusive scope group.
//RequestType: PUT
//Input: input *UpdateExclusiveScopeGroupsInput
func (s *OauthAuthServerSettingsService) UpdateExclusiveScopeGroups(input *UpdateExclusiveScopeGroupsInput) (output *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "UpdateExclusiveScopeGroups",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ScopeGroupEntry{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//RemoveExclusiveScopeGroup - Remove an existing exclusive scope group.
//RequestType: DELETE
//Input: input *RemoveExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) RemoveExclusiveScopeGroup(input *RemoveExclusiveScopeGroupInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	op := &request.Operation{
		Name:       "RemoveExclusiveScopeGroup",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddCommonScopeInput struct {
	Body models.ScopeEntry
}

type AddCommonScopeGroupInput struct {
	Body models.ScopeGroupEntry
}

type AddExclusiveScopeInput struct {
	Body models.ScopeEntry
}

type AddExclusiveScopeGroupInput struct {
	Body models.ScopeGroupEntry
}

type GetCommonScopeInput struct {
	Name string
}

type GetCommonScopeGroupInput struct {
	Name string
}

type GetExclusiveScopeInput struct {
	Name string
}

type GetExclusiveScopeGroupInput struct {
	Name string
}

type RemoveCommonScopeInput struct {
	Name string
}

type RemoveCommonScopeGroupInput struct {
	Name string
}

type RemoveExclusiveScopeInput struct {
	Name string
}

type RemoveExclusiveScopeGroupInput struct {
	Name string
}

type UpdateAuthorizationServerSettingsInput struct {
	Body models.AuthorizationServerSettings
}

type UpdateCommonScopeInput struct {
	Body models.ScopeEntry
	Name string
}

type UpdateCommonScopeGroupInput struct {
	Body models.ScopeGroupEntry
	Name string
}

type UpdateExclusiveScopeInput struct {
	Body models.ScopeEntry
	Name string
}

type UpdateExclusiveScopeGroupsInput struct {
	Body models.ScopeGroupEntry
	Name string
}
