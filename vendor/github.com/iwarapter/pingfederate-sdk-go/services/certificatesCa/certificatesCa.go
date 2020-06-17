package certificatesCa

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
	ServiceName = "CertificatesCa"
)

type CertificatesCaService struct {
	*client.PfClient
}

// New creates a new instance of the CertificatesCaService client.
func New(cfg *config.Config) *CertificatesCaService {

	return &CertificatesCaService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a CertificatesCa operation
func (c *CertificatesCaService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetTrustedCAs - Get list of trusted certificate authorities.
//RequestType: GET
//Input:
func (s *CertificatesCaService) GetTrustedCAs() (output *models.CertViews, resp *http.Response, err error) {
	path := "/certificates/ca"
	op := &request.Operation{
		Name:       "GetTrustedCAs",
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

//GetTrustedCert - Retrieve details of a trusted certificate authority.
//RequestType: GET
//Input: input *GetTrustedCertInput
func (s *CertificatesCaService) GetTrustedCert(input *GetTrustedCertInput) (output *models.CertView, resp *http.Response, err error) {
	path := "/certificates/ca/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetTrustedCert",
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

//DeleteTrustedCA - Delete a trusted certificate authority.
//RequestType: DELETE
//Input: input *DeleteTrustedCAInput
func (s *CertificatesCaService) DeleteTrustedCA(input *DeleteTrustedCAInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/certificates/ca/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteTrustedCA",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ImportTrustedCA - Import a new trusted certificate authority.
//RequestType: POST
//Input: input *ImportTrustedCAInput
func (s *CertificatesCaService) ImportTrustedCA(input *ImportTrustedCAInput) (output *models.CertView, resp *http.Response, err error) {
	path := "/certificates/ca/import"
	op := &request.Operation{
		Name:       "ImportTrustedCA",
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

//ExportCertificateFile - Download the certificate from a given trusted certificate authority.
//RequestType: GET
//Input: input *ExportCertificateFileInput
func (s *CertificatesCaService) ExportCertificateFile(input *ExportCertificateFileInput) (output *string, resp *http.Response, err error) {
	path := "/certificates/ca/{id}/file"
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

type DeleteTrustedCAInput struct {
	Id string
}

type ExportCertificateFileInput struct {
	Id string
}

type GetTrustedCertInput struct {
	Id string
}

type ImportTrustedCAInput struct {
	Body models.X509File
}
