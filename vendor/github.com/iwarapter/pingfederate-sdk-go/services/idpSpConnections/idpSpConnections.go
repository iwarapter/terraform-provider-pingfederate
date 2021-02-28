package idpSpConnections

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
	ServiceName = "IdpSpConnections"
)

type IdpSpConnectionsService struct {
	*client.PfClient
}

// New creates a new instance of the IdpSpConnectionsService client.
func New(cfg *config.Config) *IdpSpConnectionsService {

	return &IdpSpConnectionsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IdpSpConnections operation
func (c *IdpSpConnectionsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetConnections - Get list of SP connections.
//RequestType: GET
//Input: input *GetConnectionsInput
func (s *IdpSpConnectionsService) GetConnections(input *GetConnectionsInput) (output *models.SpConnections, resp *http.Response, err error) {
	return s.GetConnectionsWithContext(context.Background(), input)
}

//GetConnectionsWithContext - Get list of SP connections.
//RequestType: GET
//Input: ctx context.Context, input *GetConnectionsInput
func (s *IdpSpConnectionsService) GetConnectionsWithContext(ctx context.Context, input *GetConnectionsInput) (output *models.SpConnections, resp *http.Response, err error) {
	path := "/idp/spConnections"
	op := &request.Operation{
		Name:       "GetConnections",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"entityId":      input.EntityId,
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
		},
	}
	output = &models.SpConnections{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateConnection - Create a new SP connection.
//RequestType: POST
//Input: input *CreateConnectionInput
func (s *IdpSpConnectionsService) CreateConnection(input *CreateConnectionInput) (output *models.SpConnection, resp *http.Response, err error) {
	return s.CreateConnectionWithContext(context.Background(), input)
}

//CreateConnectionWithContext - Create a new SP connection.
//RequestType: POST
//Input: ctx context.Context, input *CreateConnectionInput
func (s *IdpSpConnectionsService) CreateConnectionWithContext(ctx context.Context, input *CreateConnectionInput) (output *models.SpConnection, resp *http.Response, err error) {
	path := "/idp/spConnections"
	op := &request.Operation{
		Name:       "CreateConnection",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.SpConnection{}
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

//GetConnection - Find SP connection by ID.
//RequestType: GET
//Input: input *GetConnectionInput
func (s *IdpSpConnectionsService) GetConnection(input *GetConnectionInput) (output *models.SpConnection, resp *http.Response, err error) {
	return s.GetConnectionWithContext(context.Background(), input)
}

//GetConnectionWithContext - Find SP connection by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetConnectionInput
func (s *IdpSpConnectionsService) GetConnectionWithContext(ctx context.Context, input *GetConnectionInput) (output *models.SpConnection, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetConnection",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SpConnection{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateConnection - Update an SP connection.
//RequestType: PUT
//Input: input *UpdateConnectionInput
func (s *IdpSpConnectionsService) UpdateConnection(input *UpdateConnectionInput) (output *models.SpConnection, resp *http.Response, err error) {
	return s.UpdateConnectionWithContext(context.Background(), input)
}

//UpdateConnectionWithContext - Update an SP connection.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateConnectionInput
func (s *IdpSpConnectionsService) UpdateConnectionWithContext(ctx context.Context, input *UpdateConnectionInput) (output *models.SpConnection, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateConnection",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SpConnection{}
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

//DeleteConnection - Delete an SP connection.
//RequestType: DELETE
//Input: input *DeleteConnectionInput
func (s *IdpSpConnectionsService) DeleteConnection(input *DeleteConnectionInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteConnectionWithContext(context.Background(), input)
}

//DeleteConnectionWithContext - Delete an SP connection.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteConnectionInput
func (s *IdpSpConnectionsService) DeleteConnectionWithContext(ctx context.Context, input *DeleteConnectionInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteConnection",
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

//GetSigningSettings - Get the SP connection's signature settings.
//RequestType: GET
//Input: input *GetSigningSettingsInput
func (s *IdpSpConnectionsService) GetSigningSettings(input *GetSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error) {
	return s.GetSigningSettingsWithContext(context.Background(), input)
}

//GetSigningSettingsWithContext - Get the SP connection's signature settings.
//RequestType: GET
//Input: ctx context.Context, input *GetSigningSettingsInput
func (s *IdpSpConnectionsService) GetSigningSettingsWithContext(ctx context.Context, input *GetSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/signingSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetSigningSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SigningSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSigningSettings - Update the SP connection's signature settings.
//RequestType: PUT
//Input: input *UpdateSigningSettingsInput
func (s *IdpSpConnectionsService) UpdateSigningSettings(input *UpdateSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error) {
	return s.UpdateSigningSettingsWithContext(context.Background(), input)
}

//UpdateSigningSettingsWithContext - Update the SP connection's signature settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSigningSettingsInput
func (s *IdpSpConnectionsService) UpdateSigningSettingsWithContext(ctx context.Context, input *UpdateSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/signingSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateSigningSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SigningSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddConnectionCert - Add a new SP connection certificate.
//RequestType: POST
//Input: input *AddConnectionCertInput
func (s *IdpSpConnectionsService) AddConnectionCert(input *AddConnectionCertInput) (output *models.ConnectionCert, resp *http.Response, err error) {
	return s.AddConnectionCertWithContext(context.Background(), input)
}

//AddConnectionCertWithContext - Add a new SP connection certificate.
//RequestType: POST
//Input: ctx context.Context, input *AddConnectionCertInput
func (s *IdpSpConnectionsService) AddConnectionCertWithContext(ctx context.Context, input *AddConnectionCertInput) (output *models.ConnectionCert, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "AddConnectionCert",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ConnectionCert{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetConnectionCerts - Get the SP connection's certificates.
//RequestType: GET
//Input: input *GetConnectionCertsInput
func (s *IdpSpConnectionsService) GetConnectionCerts(input *GetConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error) {
	return s.GetConnectionCertsWithContext(context.Background(), input)
}

//GetConnectionCertsWithContext - Get the SP connection's certificates.
//RequestType: GET
//Input: ctx context.Context, input *GetConnectionCertsInput
func (s *IdpSpConnectionsService) GetConnectionCertsWithContext(ctx context.Context, input *GetConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetConnectionCerts",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ConnectionCerts{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateConnectionCerts - Update the SP connection's certificates.
//RequestType: PUT
//Input: input *UpdateConnectionCertsInput
func (s *IdpSpConnectionsService) UpdateConnectionCerts(input *UpdateConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error) {
	return s.UpdateConnectionCertsWithContext(context.Background(), input)
}

//UpdateConnectionCertsWithContext - Update the SP connection's certificates.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateConnectionCertsInput
func (s *IdpSpConnectionsService) UpdateConnectionCertsWithContext(ctx context.Context, input *UpdateConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateConnectionCerts",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ConnectionCerts{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDecryptionKeys - Get the decryption keys of an SP connection.
//RequestType: GET
//Input: input *GetDecryptionKeysInput
func (s *IdpSpConnectionsService) GetDecryptionKeys(input *GetDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error) {
	return s.GetDecryptionKeysWithContext(context.Background(), input)
}

//GetDecryptionKeysWithContext - Get the decryption keys of an SP connection.
//RequestType: GET
//Input: ctx context.Context, input *GetDecryptionKeysInput
func (s *IdpSpConnectionsService) GetDecryptionKeysWithContext(ctx context.Context, input *GetDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/decryptionKeys"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetDecryptionKeys",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.DecryptionKeys{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateDecryptionKeys - Updating the SP connection's decryption keys.
//RequestType: PUT
//Input: input *UpdateDecryptionKeysInput
func (s *IdpSpConnectionsService) UpdateDecryptionKeys(input *UpdateDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error) {
	return s.UpdateDecryptionKeysWithContext(context.Background(), input)
}

//UpdateDecryptionKeysWithContext - Updating the SP connection's decryption keys.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateDecryptionKeysInput
func (s *IdpSpConnectionsService) UpdateDecryptionKeysWithContext(ctx context.Context, input *UpdateDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/decryptionKeys"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateDecryptionKeys",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.DecryptionKeys{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddConnectionCertInput struct {
	Body models.ConnectionCert
	Id   string
}

type CreateConnectionInput struct {
	Body models.SpConnection

	BypassExternalValidation *bool
}

type DeleteConnectionInput struct {
	Id string
}

type GetConnectionInput struct {
	Id string
}

type GetConnectionCertsInput struct {
	Id string
}

type GetConnectionsInput struct {
	EntityId      string
	Page          string
	NumberPerPage string
	Filter        string
}

type GetDecryptionKeysInput struct {
	Id string
}

type GetSigningSettingsInput struct {
	Id string
}

type UpdateConnectionInput struct {
	Body models.SpConnection
	Id   string

	BypassExternalValidation *bool
}

type UpdateConnectionCertsInput struct {
	Body models.ConnectionCerts
	Id   string
}

type UpdateDecryptionKeysInput struct {
	Body models.DecryptionKeys
	Id   string
}

type UpdateSigningSettingsInput struct {
	Body models.SigningSettings
	Id   string
}
