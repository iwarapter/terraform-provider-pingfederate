package idpDefaultUrls

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
	ServiceName = "IdpDefaultUrls"
)

type IdpDefaultUrlsService struct {
	*client.PfClient
}

// New creates a new instance of the IdpDefaultUrlsService client.
func New(cfg *config.Config) *IdpDefaultUrlsService {

	return &IdpDefaultUrlsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IdpDefaultUrls operation
func (c *IdpDefaultUrlsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetDefaultUrl - Gets the IDP Default URL settings.
//RequestType: GET
//Input:
func (s *IdpDefaultUrlsService) GetDefaultUrl() (output *models.IdpDefaultUrl, resp *http.Response, err error) {
	path := "/idp/defaultUrls"
	op := &request.Operation{
		Name:       "GetDefaultUrl",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IdpDefaultUrl{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateDefaultUrlSettings - Update the IDP Default URL settings.
//RequestType: PUT
//Input: input *UpdateDefaultUrlSettingsInput
func (s *IdpDefaultUrlsService) UpdateDefaultUrlSettings(input *UpdateDefaultUrlSettingsInput) (output *models.IdpDefaultUrl, resp *http.Response, err error) {
	path := "/idp/defaultUrls"
	op := &request.Operation{
		Name:       "UpdateDefaultUrlSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.IdpDefaultUrl{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateDefaultUrlSettingsInput struct {
	Body models.IdpDefaultUrl
}
