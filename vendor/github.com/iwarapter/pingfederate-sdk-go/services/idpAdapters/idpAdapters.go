package idpAdapters

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
	ServiceName = "IdpAdapters"
)

type IdpAdaptersService struct {
	*client.PfClient
}

// New creates a new instance of the IdpAdaptersService client.
func New(cfg *config.Config) *IdpAdaptersService {

	return &IdpAdaptersService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IdpAdapters operation
func (c *IdpAdaptersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetIdpAdapterDescriptors - Get the list of available IdP adapter descriptors.
//RequestType: GET
//Input:
func (s *IdpAdaptersService) GetIdpAdapterDescriptors() (output *models.IdpAdapterDescriptors, resp *http.Response, err error) {
	path := "/idp/adapters/descriptors"
	op := &request.Operation{
		Name:       "GetIdpAdapterDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpAdapterDescriptors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetIdpAdapterDescriptorsById - Get the description of an IdP adapter plugin by ID.
//RequestType: GET
//Input: input *GetIdpAdapterDescriptorsByIdInput
func (s *IdpAdaptersService) GetIdpAdapterDescriptorsById(input *GetIdpAdapterDescriptorsByIdInput) (output *models.IdpAdapterDescriptor, resp *http.Response, err error) {
	path := "/idp/adapters/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetIdpAdapterDescriptorsById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpAdapterDescriptor{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetIdpAdapters - Get the list of configured IdP adapter instances.
//RequestType: GET
//Input: input *GetIdpAdaptersInput
func (s *IdpAdaptersService) GetIdpAdapters(input *GetIdpAdaptersInput) (output *models.IdpAdapters, resp *http.Response, err error) {
	path := "/idp/adapters"
	op := &request.Operation{
		Name:       "GetIdpAdapters",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
		},
	}
	output = &models.IdpAdapters{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateIdpAdapter - Create a new IdP adapter instance.
//RequestType: POST
//Input: input *CreateIdpAdapterInput
func (s *IdpAdaptersService) CreateIdpAdapter(input *CreateIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters"
	op := &request.Operation{
		Name:       "CreateIdpAdapter",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.IdpAdapter{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetIdpAdapter - Find an IdP adapter instance by ID.
//RequestType: GET
//Input: input *GetIdpAdapterInput
func (s *IdpAdaptersService) GetIdpAdapter(input *GetIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetIdpAdapter",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpAdapter{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateIdpAdapter - Update an IdP adapter instance.
//RequestType: PUT
//Input: input *UpdateIdpAdapterInput
func (s *IdpAdaptersService) UpdateIdpAdapter(input *UpdateIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateIdpAdapter",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.IdpAdapter{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteIdpAdapter - Delete an IdP adapter instance.
//RequestType: DELETE
//Input: input *DeleteIdpAdapterInput
func (s *IdpAdaptersService) DeleteIdpAdapter(input *DeleteIdpAdapterInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteIdpAdapter",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetActions - List the actions for an IdP adapter instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *IdpAdaptersService) GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions"
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

//GetAction - Find an IdP adapter instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *IdpAdaptersService) GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions/{actionId}"
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

//InvokeAction - Invokes an action for an IdP adapter instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *IdpAdaptersService) InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions/{actionId}/invokeAction"
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

type CreateIdpAdapterInput struct {
	Body models.IdpAdapter

	BypassExternalValidation *bool
}

type DeleteIdpAdapterInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetIdpAdapterInput struct {
	Id string
}

type GetIdpAdapterDescriptorsByIdInput struct {
	Id string
}

type GetIdpAdaptersInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateIdpAdapterInput struct {
	Body models.IdpAdapter
	Id   string

	BypassExternalValidation *bool
}
