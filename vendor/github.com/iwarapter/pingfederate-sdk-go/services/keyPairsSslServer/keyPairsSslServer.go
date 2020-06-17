package keyPairsSslServer

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
	ServiceName = "KeyPairsSslServer"
)

type KeyPairsSslServerService struct {
	*client.PfClient
}

// New creates a new instance of the KeyPairsSslServerService client.
func New(cfg *config.Config) *KeyPairsSslServerService {

	return &KeyPairsSslServerService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a KeyPairsSslServer operation
func (c *KeyPairsSslServerService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetKeyPairs - Get list of key pairs.
//RequestType: GET
//Input:
func (s *KeyPairsSslServerService) GetKeyPairs() (output *models.KeyPairViews, resp *http.Response, err error) {
	path := "/keyPairs/sslServer"
	op := &request.Operation{
		Name:       "GetKeyPairs",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KeyPairViews{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ImportKeyPair - Import a new key pair.
//RequestType: POST
//Input: input *ImportKeyPairInput
func (s *KeyPairsSslServerService) ImportKeyPair(input *ImportKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/import"
	op := &request.Operation{
		Name:       "ImportKeyPair",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateKeyPair - Generate a new key pair.
//RequestType: POST
//Input: input *CreateKeyPairInput
func (s *KeyPairsSslServerService) CreateKeyPair(input *CreateKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/generate"
	op := &request.Operation{
		Name:       "CreateKeyPair",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetKeyPair - Retrieve details of a key pair.
//RequestType: GET
//Input: input *GetKeyPairInput
func (s *KeyPairsSslServerService) GetKeyPair(input *GetKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetKeyPair",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteKeyPair - Delete a key pair.
//RequestType: DELETE
//Input: input *DeleteKeyPairInput
func (s *KeyPairsSslServerService) DeleteKeyPair(input *DeleteKeyPairInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteKeyPair",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ExportCsr - Generate a new certificate signing request (CSR) for this key pair.
//RequestType: GET
//Input: input *ExportCsrInput
func (s *KeyPairsSslServerService) ExportCsr(input *ExportCsrInput) (output *string, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/csr"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "ExportCsr",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = pingfederate.String("")
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ImportCsrResponse - Import a CSR response for this key pair.
//RequestType: POST
//Input: input *ImportCsrResponseInput
func (s *KeyPairsSslServerService) ImportCsrResponse(input *ImportCsrResponseInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/csr"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "ImportCsrResponse",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.KeyPairView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ExportPKCS12File - Download the key pair in PKCS12 format.
//RequestType: POST
//Input: input *ExportPKCS12FileInput
func (s *KeyPairsSslServerService) ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/pkcs12"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "ExportPKCS12File",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}

	req := s.newRequest(op, input.Body, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//ExportCertificateFile - Download the certificate from a given key pair.
//RequestType: GET
//Input: input *ExportCertificateFileInput
func (s *KeyPairsSslServerService) ExportCertificateFile(input *ExportCertificateFileInput) (output *string, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/certificate"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "ExportCertificateFile",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = pingfederate.String("")
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSettings - Get the SSL Server Certificate Settings.
//RequestType: GET
//Input:
func (s *KeyPairsSslServerService) GetSettings() (output *models.SslServerSettings, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SslServerSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Update the SSL Server Certificate Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *KeyPairsSslServerService) UpdateSettings(input *UpdateSettingsInput) (output *models.SslServerSettings, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SslServerSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateKeyPairInput struct {
	Body models.NewKeyPairSettings
}

type DeleteKeyPairInput struct {
	Id string
}

type ExportCertificateFileInput struct {
	Id string
}

type ExportCsrInput struct {
	Id string
}

type ExportPKCS12FileInput struct {
	Body models.PKCS12ExportSettings
	Id   string
}

type GetKeyPairInput struct {
	Id string
}

type ImportCsrResponseInput struct {
	Body models.CSRResponse
	Id   string
}

type ImportKeyPairInput struct {
	Body models.PKCS12File
}

type UpdateSettingsInput struct {
	Body models.SslServerSettings
}
