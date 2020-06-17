package keyPairsSigning

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
	ServiceName = "KeyPairsSigning"
)

type KeyPairsSigningService struct {
	*client.PfClient
}

// New creates a new instance of the KeyPairsSigningService client.
func New(cfg *config.Config) *KeyPairsSigningService {

	return &KeyPairsSigningService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a KeyPairsSigning operation
func (c *KeyPairsSigningService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetKeyPairs - Get list of key pairs.
//RequestType: GET
//Input:
func (s *KeyPairsSigningService) GetKeyPairs() (output *models.KeyPairViews, resp *http.Response, err error) {
	path := "/keyPairs/signing"
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
func (s *KeyPairsSigningService) ImportKeyPair(input *ImportKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/signing/import"
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
func (s *KeyPairsSigningService) CreateKeyPair(input *CreateKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/signing/generate"
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
func (s *KeyPairsSigningService) GetKeyPair(input *GetKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}"
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
func (s *KeyPairsSigningService) DeleteKeyPair(input *DeleteKeyPairInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}"
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
func (s *KeyPairsSigningService) ExportCsr(input *ExportCsrInput) (output *string, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/csr"
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
func (s *KeyPairsSigningService) ImportCsrResponse(input *ImportCsrResponseInput) (output *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/csr"
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
func (s *KeyPairsSigningService) ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/pkcs12"
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
func (s *KeyPairsSigningService) ExportCertificateFile(input *ExportCertificateFileInput) (output *string, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/certificate"
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

//GetRotationSettings - Retrieve details of rotation settings for a key pair.
//RequestType: GET
//Input: input *GetRotationSettingsInput
func (s *KeyPairsSigningService) GetRotationSettings(input *GetRotationSettingsInput) (output *models.KeyPairRotationSettings, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/rotationSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetRotationSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KeyPairRotationSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateRotationSettings - Add rotation settings to a key pair
//RequestType: PUT
//Input: input *UpdateRotationSettingsInput
func (s *KeyPairsSigningService) UpdateRotationSettings(input *UpdateRotationSettingsInput) (output *models.KeyPairRotationSettings, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/rotationSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateRotationSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.KeyPairRotationSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteKeyPairRotationSettings - Delete rotation settings for a signing key pair.
//RequestType: DELETE
//Input: input *DeleteKeyPairRotationSettingsInput
func (s *KeyPairsSigningService) DeleteKeyPairRotationSettings(input *DeleteKeyPairRotationSettingsInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/rotationSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteKeyPairRotationSettings",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

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

type DeleteKeyPairRotationSettingsInput struct {
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

type GetRotationSettingsInput struct {
	Id string
}

type ImportCsrResponseInput struct {
	Body models.CSRResponse
	Id   string
}

type ImportKeyPairInput struct {
	Body models.PKCS12File
}

type UpdateRotationSettingsInput struct {
	Body models.KeyPairRotationSettings
	Id   string
}
