package spDefaultUrls

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
	ServiceName = "SpDefaultUrls"
)

type SpDefaultUrlsService struct {
	*client.PfClient
}

// New creates a new instance of the SpDefaultUrlsService client.
func New(cfg *config.Config) *SpDefaultUrlsService {

	return &SpDefaultUrlsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a SpDefaultUrls operation
func (c *SpDefaultUrlsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetDefaultUrls - Gets the SP Default URLs. These are Values that affect the user's experience when executing SP-initiated SSO operations.
//RequestType: GET
//Input:
func (s *SpDefaultUrlsService) GetDefaultUrls() (output *models.SpDefaultUrls, resp *http.Response, err error) {
	path := "/sp/defaultUrls"
	op := &request.Operation{
		Name:       "GetDefaultUrls",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SpDefaultUrls{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateDefaultUrls - Update the SP Default URLs. Enter values that affect the user's experience when executing SP-initiated SSO operations.
//RequestType: PUT
//Input: input *UpdateDefaultUrlsInput
func (s *SpDefaultUrlsService) UpdateDefaultUrls(input *UpdateDefaultUrlsInput) (output *models.SpDefaultUrls, resp *http.Response, err error) {
	path := "/sp/defaultUrls"
	op := &request.Operation{
		Name:       "UpdateDefaultUrls",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SpDefaultUrls{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateDefaultUrlsInput struct {
	Body models.SpDefaultUrls
}
