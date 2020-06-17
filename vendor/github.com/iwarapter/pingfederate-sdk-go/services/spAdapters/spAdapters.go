package spAdapters

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
	ServiceName = "SpAdapters"
)

type SpAdaptersService struct {
	*client.PfClient
}

// New creates a new instance of the SpAdaptersService client.
func New(cfg *config.Config) *SpAdaptersService {

	return &SpAdaptersService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a SpAdapters operation
func (c *SpAdaptersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSpAdapterDescriptors - Get the list of available SP adapter descriptors.
//RequestType: GET
//Input:
func (s *SpAdaptersService) GetSpAdapterDescriptors() (output *models.SpAdapterDescriptors, resp *http.Response, err error) {
	path := "/sp/adapters/descriptors"
	op := &request.Operation{
		Name:       "GetSpAdapterDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SpAdapterDescriptors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSpAdapterDescriptorsById - Get the description of an SP adapter plugin by ID.
//RequestType: GET
//Input: input *GetSpAdapterDescriptorsByIdInput
func (s *SpAdaptersService) GetSpAdapterDescriptorsById(input *GetSpAdapterDescriptorsByIdInput) (output *models.SpAdapterDescriptor, resp *http.Response, err error) {
	path := "/sp/adapters/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetSpAdapterDescriptorsById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SpAdapterDescriptor{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSpAdapters - Get the list of configured SP adapter instances.
//RequestType: GET
//Input: input *GetSpAdaptersInput
func (s *SpAdaptersService) GetSpAdapters(input *GetSpAdaptersInput) (output *models.SpAdapters, resp *http.Response, err error) {
	path := "/sp/adapters"
	op := &request.Operation{
		Name:       "GetSpAdapters",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SpAdapters{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateSpAdapter - Create a new SP adapter instance.
//RequestType: POST
//Input: input *CreateSpAdapterInput
func (s *SpAdaptersService) CreateSpAdapter(input *CreateSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters"
	op := &request.Operation{
		Name:       "CreateSpAdapter",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.SpAdapter{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSpAdapter - Find an SP adapter instance by ID.
//RequestType: GET
//Input: input *GetSpAdapterInput
func (s *SpAdaptersService) GetSpAdapter(input *GetSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetSpAdapter",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SpAdapter{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSpAdapter - Update an SP adapter instance.
//RequestType: PUT
//Input: input *UpdateSpAdapterInput
func (s *SpAdaptersService) UpdateSpAdapter(input *UpdateSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateSpAdapter",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SpAdapter{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteSpAdapter - Delete an SP adapter instance.
//RequestType: DELETE
//Input: input *DeleteSpAdapterInput
func (s *SpAdaptersService) DeleteSpAdapter(input *DeleteSpAdapterInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteSpAdapter",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetActions - List the actions for an SP adapter instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *SpAdaptersService) GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions"
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

//GetAction - Find an SP adapter instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *SpAdaptersService) GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions/{actionId}"
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

//InvokeAction - Invokes an action for an SP adapter instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *SpAdaptersService) InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions/{actionId}/invokeAction"
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

//GetUrlMappings - (Deprecated) List the mappings between URLs and adapter instances.
//RequestType: GET
//Input:
func (s *SpAdaptersService) GetUrlMappings() (output *models.SpAdapterUrlMappings, resp *http.Response, err error) {
	path := "/sp/adapters/urlMappings"
	op := &request.Operation{
		Name:       "GetUrlMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SpAdapterUrlMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateUrlMappings - (Deprecated) Update the mappings between URLs and adapters instances.
//RequestType: PUT
//Input: input *UpdateUrlMappingsInput
func (s *SpAdaptersService) UpdateUrlMappings(input *UpdateUrlMappingsInput) (output *models.SpAdapterUrlMappings, resp *http.Response, err error) {
	path := "/sp/adapters/urlMappings"
	op := &request.Operation{
		Name:       "UpdateUrlMappings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SpAdapterUrlMappings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateSpAdapterInput struct {
	Body models.SpAdapter
}

type DeleteSpAdapterInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetSpAdapterInput struct {
	Id string
}

type GetSpAdapterDescriptorsByIdInput struct {
	Id string
}

type GetSpAdaptersInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateSpAdapterInput struct {
	Body models.SpAdapter
	Id   string
}

type UpdateUrlMappingsInput struct {
	Body models.SpAdapterUrlMappings
}
