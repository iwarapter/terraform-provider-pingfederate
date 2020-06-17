package license

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "License"
)

type LicenseService struct {
	*client.PfClient
}

// New creates a new instance of the LicenseService client.
func New(cfg *config.Config) *LicenseService {

	return &LicenseService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a License operation
func (c *LicenseService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetLicenseAgreement - Get license agreement link.
//RequestType: GET
//Input:
func (s *LicenseService) GetLicenseAgreement() (output *models.LicenseAgreementInfo, resp *http.Response, err error) {
	path := "/license/agreement"
	op := &request.Operation{
		Name:       "GetLicenseAgreement",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.LicenseAgreementInfo{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateLicenseAgreement - Accept license agreement.
//RequestType: PUT
//Input: input *UpdateLicenseAgreementInput
func (s *LicenseService) UpdateLicenseAgreement(input *UpdateLicenseAgreementInput) (output *models.LicenseAgreementInfo, resp *http.Response, err error) {
	path := "/license/agreement"
	op := &request.Operation{
		Name:       "UpdateLicenseAgreement",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.LicenseAgreementInfo{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetLicense - Get a license summary.
//RequestType: GET
//Input:
func (s *LicenseService) GetLicense() (output *models.LicenseView, resp *http.Response, err error) {
	path := "/license"
	op := &request.Operation{
		Name:       "GetLicense",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.LicenseView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateLicense - Import a license.
//RequestType: PUT
//Input: input *UpdateLicenseInput
func (s *LicenseService) UpdateLicense(input *UpdateLicenseInput) (output *models.LicenseView, resp *http.Response, err error) {
	path := "/license"
	op := &request.Operation{
		Name:       "UpdateLicense",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.LicenseView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateLicenseInput struct {
	Body models.LicenseFile
}

type UpdateLicenseAgreementInput struct {
	Body models.LicenseAgreementInfo
}
