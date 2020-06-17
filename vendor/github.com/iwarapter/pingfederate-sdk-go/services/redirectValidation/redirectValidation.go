package redirectValidation

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
	ServiceName = "RedirectValidation"
)

type RedirectValidationService struct {
	*client.PfClient
}

// New creates a new instance of the RedirectValidationService client.
func New(cfg *config.Config) *RedirectValidationService {

	return &RedirectValidationService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a RedirectValidation operation
func (c *RedirectValidationService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetRedirectValidationSettings - Retrieve redirect validation settings.
//RequestType: GET
//Input:
func (s *RedirectValidationService) GetRedirectValidationSettings() (output *models.RedirectValidationSettings, resp *http.Response, err error) {
	path := "/redirectValidation"
	op := &request.Operation{
		Name:       "GetRedirectValidationSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.RedirectValidationSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateRedirectValidationSettings - Update redirect validation settings.
//RequestType: PUT
//Input: input *UpdateRedirectValidationSettingsInput
func (s *RedirectValidationService) UpdateRedirectValidationSettings(input *UpdateRedirectValidationSettingsInput) (output *models.RedirectValidationSettings, resp *http.Response, err error) {
	path := "/redirectValidation"
	op := &request.Operation{
		Name:       "UpdateRedirectValidationSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.RedirectValidationSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateRedirectValidationSettingsInput struct {
	Body models.RedirectValidationSettings
}
