package incomingProxySettings

import (
	"context"
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
	ServiceName = "IncomingProxySettings"
)

type IncomingProxySettingsService struct {
	*client.PfClient
}

// New creates a new instance of the IncomingProxySettingsService client.
func New(cfg *config.Config) *IncomingProxySettingsService {

	return &IncomingProxySettingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IncomingProxySettings operation
func (c *IncomingProxySettingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetIncomingProxySettings - Get incoming proxy settings.
//RequestType: GET
//Input:
func (s *IncomingProxySettingsService) GetIncomingProxySettings() (output *models.IncomingProxySettings, resp *http.Response, err error) {
	return s.GetIncomingProxySettingsWithContext(context.Background())
}

//GetIncomingProxySettingsWithContext - Get incoming proxy settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *IncomingProxySettingsService) GetIncomingProxySettingsWithContext(ctx context.Context) (output *models.IncomingProxySettings, resp *http.Response, err error) {
	path := "/incomingProxySettings"
	op := &request.Operation{
		Name:       "GetIncomingProxySettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.IncomingProxySettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateIncomingProxySettings - Update incoming proxy settings.
//RequestType: PUT
//Input: input *UpdateIncomingProxySettingsInput
func (s *IncomingProxySettingsService) UpdateIncomingProxySettings(input *UpdateIncomingProxySettingsInput) (output *models.IncomingProxySettings, resp *http.Response, err error) {
	return s.UpdateIncomingProxySettingsWithContext(context.Background(), input)
}

//UpdateIncomingProxySettingsWithContext - Update incoming proxy settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateIncomingProxySettingsInput
func (s *IncomingProxySettingsService) UpdateIncomingProxySettingsWithContext(ctx context.Context, input *UpdateIncomingProxySettingsInput) (output *models.IncomingProxySettings, resp *http.Response, err error) {
	path := "/incomingProxySettings"
	op := &request.Operation{
		Name:       "UpdateIncomingProxySettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.IncomingProxySettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateIncomingProxySettingsInput struct {
	Body models.IncomingProxySettings
}
