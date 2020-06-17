package oauthClients

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
	ServiceName = "OauthClients"
)

type OauthClientsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthClientsService client.
func New(cfg *config.Config) *OauthClientsService {

	return &OauthClientsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthClients operation
func (c *OauthClientsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetClients - Get the list of OAuth clients.
//RequestType: GET
//Input: input *GetClientsInput
func (s *OauthClientsService) GetClients(input *GetClientsInput) (output *models.Clients, resp *http.Response, err error) {
	path := "/oauth/clients"
	op := &request.Operation{
		Name:       "GetClients",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Clients{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateClient - Create a new OAuth client.
//RequestType: POST
//Input: input *CreateClientInput
func (s *OauthClientsService) CreateClient(input *CreateClientInput) (output *models.Client, resp *http.Response, err error) {
	path := "/oauth/clients"
	op := &request.Operation{
		Name:       "CreateClient",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.Client{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetClient - Find the OAuth client by ID.
//RequestType: GET
//Input: input *GetClientInput
func (s *OauthClientsService) GetClient(input *GetClientInput) (output *models.Client, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetClient",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Client{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateClient - Updates the OAuth client.
//RequestType: PUT
//Input: input *UpdateClientInput
func (s *OauthClientsService) UpdateClient(input *UpdateClientInput) (output *models.Client, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateClient",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.Client{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteClient - Delete an OAuth client.
//RequestType: DELETE
//Input: input *DeleteClientInput
func (s *OauthClientsService) DeleteClient(input *DeleteClientInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteClient",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetClientSecret - Get the client secret of an existing OAuth client.
//RequestType: GET
//Input: input *GetClientSecretInput
func (s *OauthClientsService) GetClientSecret(input *GetClientSecretInput) (output *models.ClientSecret, resp *http.Response, err error) {
	path := "/oauth/clients/{id}/clientAuth/clientSecret"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetClientSecret",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ClientSecret{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateClientSecret - Update the client secret of an existing OAuth client.
//RequestType: PUT
//Input: input *UpdateClientSecretInput
func (s *OauthClientsService) UpdateClientSecret(input *UpdateClientSecretInput) (output *models.ClientSecret, resp *http.Response, err error) {
	path := "/oauth/clients/{id}/clientAuth/clientSecret"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateClientSecret",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ClientSecret{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateClientInput struct {
	Body models.Client
}

type DeleteClientInput struct {
	Id string
}

type GetClientInput struct {
	Id string
}

type GetClientSecretInput struct {
	Id string
}

type GetClientsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type UpdateClientInput struct {
	Body models.Client
	Id   string
}

type UpdateClientSecretInput struct {
	Body models.ClientSecret
	Id   string
}
