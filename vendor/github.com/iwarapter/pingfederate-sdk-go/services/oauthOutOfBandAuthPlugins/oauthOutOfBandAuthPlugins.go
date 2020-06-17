package oauthOutOfBandAuthPlugins

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
	ServiceName = "OauthOutOfBandAuthPlugins"
)

type OauthOutOfBandAuthPluginsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthOutOfBandAuthPluginsService client.
func New(cfg *config.Config) *OauthOutOfBandAuthPluginsService {

	return &OauthOutOfBandAuthPluginsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthOutOfBandAuthPlugins operation
func (c *OauthOutOfBandAuthPluginsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetOOBAuthPluginDescriptors - Get the list of available Out of Band authenticator plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthPluginDescriptors() (output *models.OutOfBandAuthPluginDescriptors, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/descriptors"
	op := &request.Operation{
		Name:       "GetOOBAuthPluginDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OutOfBandAuthPluginDescriptors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetOOBAuthPluginDescriptor - Get the descriptor of an Out of Band authenticator plugin.
//RequestType: GET
//Input: input *GetOOBAuthPluginDescriptorInput
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthPluginDescriptor(input *GetOOBAuthPluginDescriptorInput) (output *models.OutOfBandAuthPluginDescriptor, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetOOBAuthPluginDescriptor",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OutOfBandAuthPluginDescriptor{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetOOBAuthenticators - Get a list of Out of Band authenticator plugin instances.
//RequestType: GET
//Input:
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthenticators() (output *models.OutOfBandAuthenticators, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins"
	op := &request.Operation{
		Name:       "GetOOBAuthenticators",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OutOfBandAuthenticators{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateOOBAuthenticator - Create an Out of Band authenticator plugin instance.
//RequestType: POST
//Input: input *CreateOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) CreateOOBAuthenticator(input *CreateOOBAuthenticatorInput) (output *models.OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins"
	op := &request.Operation{
		Name:       "CreateOOBAuthenticator",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.OutOfBandAuthenticator{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetOOBAuthenticator - Get a specific Out of Band authenticator plugin instance.
//RequestType: GET
//Input: input *GetOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthenticator(input *GetOOBAuthenticatorInput) (output *models.OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetOOBAuthenticator",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OutOfBandAuthenticator{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOOBAuthenticator - Update an Out of Band authenticator plugin instance.
//RequestType: PUT
//Input: input *UpdateOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) UpdateOOBAuthenticator(input *UpdateOOBAuthenticatorInput) (output *models.OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateOOBAuthenticator",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.OutOfBandAuthenticator{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteOOBAuthenticator - Delete an Out of Band authenticator plugin instance.
//RequestType: DELETE
//Input: input *DeleteOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) DeleteOOBAuthenticator(input *DeleteOOBAuthenticatorInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteOOBAuthenticator",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetActions - List of actions for an Out of Band authenticator plugin instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *OauthOutOfBandAuthPluginsService) GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetActions",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Actions{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAction - Find an Out of Band authenticator plugin instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *OauthOutOfBandAuthPluginsService) GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions/{actionId}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	op := &request.Operation{
		Name:       "GetAction",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Action{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//InvokeAction - Invokes an action for Out of Band authenticator plugin instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *OauthOutOfBandAuthPluginsService) InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions/{actionId}/invokeAction"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	op := &request.Operation{
		Name:       "InvokeAction",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ActionResult{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateOOBAuthenticatorInput struct {
	Body models.OutOfBandAuthenticator
}

type DeleteOOBAuthenticatorInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetOOBAuthPluginDescriptorInput struct {
	Id string
}

type GetOOBAuthenticatorInput struct {
	Id string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateOOBAuthenticatorInput struct {
	Body models.OutOfBandAuthenticator
	Id   string
}
