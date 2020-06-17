package certificatesRevocation

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
	ServiceName = "CertificatesRevocation"
)

type CertificatesRevocationService struct {
	*client.PfClient
}

// New creates a new instance of the CertificatesRevocationService client.
func New(cfg *config.Config) *CertificatesRevocationService {

	return &CertificatesRevocationService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a CertificatesRevocation operation
func (c *CertificatesRevocationService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetRevocationSettings - Get certificate revocation settings.
//RequestType: GET
//Input:
func (s *CertificatesRevocationService) GetRevocationSettings() (output *models.CertificateRevocationSettings, resp *http.Response, err error) {
	path := "/certificates/revocation/settings"
	op := &request.Operation{
		Name:       "GetRevocationSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CertificateRevocationSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateRevocationSettings - Update certificate revocation settings.
//RequestType: PUT
//Input: input *UpdateRevocationSettingsInput
func (s *CertificatesRevocationService) UpdateRevocationSettings(input *UpdateRevocationSettingsInput) (output *models.CertificateRevocationSettings, resp *http.Response, err error) {
	path := "/certificates/revocation/settings"
	op := &request.Operation{
		Name:       "UpdateRevocationSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.CertificateRevocationSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetOcspCertificates - Get the list of available OCSP responder signature verification certificates.
//RequestType: GET
//Input:
func (s *CertificatesRevocationService) GetOcspCertificates() (output *models.CertViews, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates"
	op := &request.Operation{
		Name:       "GetOcspCertificates",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CertViews{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ImportOcspCertificate - Import an OCSP responder signature verification certificate.
//RequestType: POST
//Input: input *ImportOcspCertificateInput
func (s *CertificatesRevocationService) ImportOcspCertificate(input *ImportOcspCertificateInput) (output *models.CertView, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates"
	op := &request.Operation{
		Name:       "ImportOcspCertificate",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.CertView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetOcspCertificateById - Get an OCSP responder signature verification certificate by ID.
//RequestType: GET
//Input: input *GetOcspCertificateByIdInput
func (s *CertificatesRevocationService) GetOcspCertificateById(input *GetOcspCertificateByIdInput) (output *models.CertView, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetOcspCertificateById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CertView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteOcspCertificateById - Delete an OCSP responder signature verification certificate by ID.
//RequestType: DELETE
//Input: input *DeleteOcspCertificateByIdInput
func (s *CertificatesRevocationService) DeleteOcspCertificateById(input *DeleteOcspCertificateByIdInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteOcspCertificateById",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type DeleteOcspCertificateByIdInput struct {
	Id string
}

type GetOcspCertificateByIdInput struct {
	Id string
}

type ImportOcspCertificateInput struct {
	Body models.X509File
}

type UpdateRevocationSettingsInput struct {
	Body models.CertificateRevocationSettings
}
