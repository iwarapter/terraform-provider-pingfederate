package oauthClientSettings

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
	ServiceName = "OauthClientSettings"
)

type OauthClientSettingsService struct {
	*client.PfClient
}

// New creates a new instance of the OauthClientSettingsService client.
func New(cfg *config.Config) *OauthClientSettingsService {

	return &OauthClientSettingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a OauthClientSettings operation
func (c *OauthClientSettingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetClientSettings - Configure the client settings.
//RequestType: GET
//Input:
func (s *OauthClientSettingsService) GetClientSettings() (output *models.ClientSettings, resp *http.Response, err error) {
	return s.GetClientSettingsWithContext(context.Background())
}

//GetClientSettingsWithContext - Configure the client settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *OauthClientSettingsService) GetClientSettingsWithContext(ctx context.Context) (output *models.ClientSettings, resp *http.Response, err error) {
	path := "/oauth/clientSettings"
	op := &request.Operation{
		Name:       "GetClientSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ClientSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateClientSettings - Update the client settings.
//RequestType: PUT
//Input: input *UpdateClientSettingsInput
func (s *OauthClientSettingsService) UpdateClientSettings(input *UpdateClientSettingsInput) (output *models.ClientSettings, resp *http.Response, err error) {
	return s.UpdateClientSettingsWithContext(context.Background(), input)
}

//UpdateClientSettingsWithContext - Update the client settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateClientSettingsInput
func (s *OauthClientSettingsService) UpdateClientSettingsWithContext(ctx context.Context, input *UpdateClientSettingsInput) (output *models.ClientSettings, resp *http.Response, err error) {
	path := "/oauth/clientSettings"
	op := &request.Operation{
		Name:       "UpdateClientSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ClientSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateClientSettingsInput struct {
	Body models.ClientSettings
}
