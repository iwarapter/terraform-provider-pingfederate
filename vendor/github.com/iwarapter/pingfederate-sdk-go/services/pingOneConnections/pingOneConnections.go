package pingOneConnections

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
	ServiceName = "PingOneConnections"
)

type PingOneConnectionsService struct {
	*client.PfClient
}

// New creates a new instance of the PingOneConnectionsService client.
func New(cfg *config.Config) *PingOneConnectionsService {

	return &PingOneConnectionsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a PingOneConnections operation
func (c *PingOneConnectionsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetPingOneConnections - Get the list of all PingOne connections.
//RequestType: GET
//Input:
func (s *PingOneConnectionsService) GetPingOneConnections() (output *models.PingOneConnections, resp *http.Response, err error) {
	return s.GetPingOneConnectionsWithContext(context.Background())
}

//GetPingOneConnectionsWithContext - Get the list of all PingOne connections.
//RequestType: GET
//Input: ctx context.Context,
func (s *PingOneConnectionsService) GetPingOneConnectionsWithContext(ctx context.Context) (output *models.PingOneConnections, resp *http.Response, err error) {
	path := "/pingOneConnections"
	op := &request.Operation{
		Name:       "GetPingOneConnections",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingOneConnections{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreatePingOneConnection - Create a new PingOne connection.
//RequestType: POST
//Input: input *CreatePingOneConnectionInput
func (s *PingOneConnectionsService) CreatePingOneConnection(input *CreatePingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error) {
	return s.CreatePingOneConnectionWithContext(context.Background(), input)
}

//CreatePingOneConnectionWithContext - Create a new PingOne connection.
//RequestType: POST
//Input: ctx context.Context, input *CreatePingOneConnectionInput
func (s *PingOneConnectionsService) CreatePingOneConnectionWithContext(ctx context.Context, input *CreatePingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error) {
	path := "/pingOneConnections"
	op := &request.Operation{
		Name:       "CreatePingOneConnection",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.PingOneConnection{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPingOneConnection - Get a PingOne connection with the specified ID.
//RequestType: GET
//Input: input *GetPingOneConnectionInput
func (s *PingOneConnectionsService) GetPingOneConnection(input *GetPingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error) {
	return s.GetPingOneConnectionWithContext(context.Background(), input)
}

//GetPingOneConnectionWithContext - Get a PingOne connection with the specified ID.
//RequestType: GET
//Input: ctx context.Context, input *GetPingOneConnectionInput
func (s *PingOneConnectionsService) GetPingOneConnectionWithContext(ctx context.Context, input *GetPingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error) {
	path := "/pingOneConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPingOneConnection",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingOneConnection{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdatePingOneConnection - Update a PingOne connection.
//RequestType: PUT
//Input: input *UpdatePingOneConnectionInput
func (s *PingOneConnectionsService) UpdatePingOneConnection(input *UpdatePingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error) {
	return s.UpdatePingOneConnectionWithContext(context.Background(), input)
}

//UpdatePingOneConnectionWithContext - Update a PingOne connection.
//RequestType: PUT
//Input: ctx context.Context, input *UpdatePingOneConnectionInput
func (s *PingOneConnectionsService) UpdatePingOneConnectionWithContext(ctx context.Context, input *UpdatePingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error) {
	path := "/pingOneConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdatePingOneConnection",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.PingOneConnection{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeletePingOneConnection - Delete a PingOne connection.
//RequestType: DELETE
//Input: input *DeletePingOneConnectionInput
func (s *PingOneConnectionsService) DeletePingOneConnection(input *DeletePingOneConnectionInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeletePingOneConnectionWithContext(context.Background(), input)
}

//DeletePingOneConnectionWithContext - Delete a PingOne connection.
//RequestType: DELETE
//Input: ctx context.Context, input *DeletePingOneConnectionInput
func (s *PingOneConnectionsService) DeletePingOneConnectionWithContext(ctx context.Context, input *DeletePingOneConnectionInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/pingOneConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeletePingOneConnection",
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

//GetCredentialStatus - Get the status of the credential associated with the PingOne connection
//RequestType: GET
//Input: input *GetCredentialStatusInput
func (s *PingOneConnectionsService) GetCredentialStatus(input *GetCredentialStatusInput) (output *models.PingOneCredentialStatus, resp *http.Response, err error) {
	return s.GetCredentialStatusWithContext(context.Background(), input)
}

//GetCredentialStatusWithContext - Get the status of the credential associated with the PingOne connection
//RequestType: GET
//Input: ctx context.Context, input *GetCredentialStatusInput
func (s *PingOneConnectionsService) GetCredentialStatusWithContext(ctx context.Context, input *GetCredentialStatusInput) (output *models.PingOneCredentialStatus, resp *http.Response, err error) {
	path := "/pingOneConnections/{id}/credentialStatus"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetCredentialStatus",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.PingOneCredentialStatus{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPingOneConnectionEnvironments - Get the list of environments that the PingOne connection has access to.
//RequestType: GET
//Input: input *GetPingOneConnectionEnvironmentsInput
func (s *PingOneConnectionsService) GetPingOneConnectionEnvironments(input *GetPingOneConnectionEnvironmentsInput) (output *models.PingOneEnvironments, resp *http.Response, err error) {
	return s.GetPingOneConnectionEnvironmentsWithContext(context.Background(), input)
}

//GetPingOneConnectionEnvironmentsWithContext - Get the list of environments that the PingOne connection has access to.
//RequestType: GET
//Input: ctx context.Context, input *GetPingOneConnectionEnvironmentsInput
func (s *PingOneConnectionsService) GetPingOneConnectionEnvironmentsWithContext(ctx context.Context, input *GetPingOneConnectionEnvironmentsInput) (output *models.PingOneEnvironments, resp *http.Response, err error) {
	path := "/pingOneConnections/{id}/environments"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPingOneConnectionEnvironments",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
		},
	}
	output = &models.PingOneEnvironments{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPingOneConnectionUsages - Get the list of resources that reference this PingOne connection.
//RequestType: GET
//Input: input *GetPingOneConnectionUsagesInput
func (s *PingOneConnectionsService) GetPingOneConnectionUsages(input *GetPingOneConnectionUsagesInput) (output *models.ResourceUsages, resp *http.Response, err error) {
	return s.GetPingOneConnectionUsagesWithContext(context.Background(), input)
}

//GetPingOneConnectionUsagesWithContext - Get the list of resources that reference this PingOne connection.
//RequestType: GET
//Input: ctx context.Context, input *GetPingOneConnectionUsagesInput
func (s *PingOneConnectionsService) GetPingOneConnectionUsagesWithContext(ctx context.Context, input *GetPingOneConnectionUsagesInput) (output *models.ResourceUsages, resp *http.Response, err error) {
	path := "/pingOneConnections/{id}/usage"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPingOneConnectionUsages",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ResourceUsages{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetPingOneConnectionAssociations - Get information about components using this connection to access PingOne services.
//RequestType: GET
//Input: input *GetPingOneConnectionAssociationsInput
func (s *PingOneConnectionsService) GetPingOneConnectionAssociations(input *GetPingOneConnectionAssociationsInput) (output *models.ServiceAssociations, resp *http.Response, err error) {
	return s.GetPingOneConnectionAssociationsWithContext(context.Background(), input)
}

//GetPingOneConnectionAssociationsWithContext - Get information about components using this connection to access PingOne services.
//RequestType: GET
//Input: ctx context.Context, input *GetPingOneConnectionAssociationsInput
func (s *PingOneConnectionsService) GetPingOneConnectionAssociationsWithContext(ctx context.Context, input *GetPingOneConnectionAssociationsInput) (output *models.ServiceAssociations, resp *http.Response, err error) {
	path := "/pingOneConnections/{id}/serviceAssociations"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetPingOneConnectionAssociations",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ServiceAssociations{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreatePingOneConnectionInput struct {
	Body models.PingOneConnection

	BypassExternalValidation *bool
}

type DeletePingOneConnectionInput struct {
	Id string
}

type GetCredentialStatusInput struct {
	Id string
}

type GetPingOneConnectionInput struct {
	Id string
}

type GetPingOneConnectionAssociationsInput struct {
	Id string
}

type GetPingOneConnectionEnvironmentsInput struct {
	Page          string
	NumberPerPage string
	Filter        string

	Id string
}

type GetPingOneConnectionUsagesInput struct {
	Id string
}

type UpdatePingOneConnectionInput struct {
	Body models.PingOneConnection
	Id   string

	BypassExternalValidation *bool
}
