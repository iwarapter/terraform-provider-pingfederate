package serverSettings

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
	ServiceName = "ServerSettings"
)

type ServerSettingsService struct {
	*client.PfClient
}

// New creates a new instance of the ServerSettingsService client.
func New(cfg *config.Config) *ServerSettingsService {

	return &ServerSettingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a ServerSettings operation
func (c *ServerSettingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetServerSettings - Gets the server settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetServerSettings() (output *models.ServerSettings, resp *http.Response, err error) {
	path := "/serverSettings"
	op := &request.Operation{
		Name:       "GetServerSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ServerSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateServerSettings - Update the server settings.
//RequestType: PUT
//Input: input *UpdateServerSettingsInput
func (s *ServerSettingsService) UpdateServerSettings(input *UpdateServerSettingsInput) (output *models.ServerSettings, resp *http.Response, err error) {
	path := "/serverSettings"
	op := &request.Operation{
		Name:       "UpdateServerSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ServerSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetNotificationSettings - Gets the notification settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetNotificationSettings() (output *models.NotificationSettings, resp *http.Response, err error) {
	path := "/serverSettings/notifications"
	op := &request.Operation{
		Name:       "GetNotificationSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.NotificationSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateNotificationSettings - Update the notification settings.
//RequestType: PUT
//Input: input *UpdateNotificationSettingsInput
func (s *ServerSettingsService) UpdateNotificationSettings(input *UpdateNotificationSettingsInput) (output *models.NotificationSettings, resp *http.Response, err error) {
	path := "/serverSettings/notifications"
	op := &request.Operation{
		Name:       "UpdateNotificationSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.NotificationSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetEmailServerSettings - (Deprecated) Gets the email server settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetEmailServerSettings() (output *models.EmailServerSettings, resp *http.Response, err error) {
	path := "/serverSettings/emailServer"
	op := &request.Operation{
		Name:       "GetEmailServerSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.EmailServerSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateEmailServerSettings - (Deprecated) Update the email server settings
//RequestType: PUT
//Input: input *UpdateEmailServerSettingsInput
func (s *ServerSettingsService) UpdateEmailServerSettings(input *UpdateEmailServerSettingsInput) (output *models.EmailServerSettings, resp *http.Response, err error) {
	path := "/serverSettings/emailServer"
	op := &request.Operation{
		Name:       "UpdateEmailServerSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.EmailServerSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetCaptchaSettings - Gets the CAPTCHA settings.
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetCaptchaSettings() (output *models.CaptchaSettings, resp *http.Response, err error) {
	path := "/serverSettings/captchaSettings"
	op := &request.Operation{
		Name:       "GetCaptchaSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.CaptchaSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateCaptchaSettings - Update the CAPTCHA settings.
//RequestType: PUT
//Input: input *UpdateCaptchaSettingsInput
func (s *ServerSettingsService) UpdateCaptchaSettings(input *UpdateCaptchaSettingsInput) (output *models.CaptchaSettings, resp *http.Response, err error) {
	path := "/serverSettings/captchaSettings"
	op := &request.Operation{
		Name:       "UpdateCaptchaSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.CaptchaSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSystemKeys - Get the system keys.
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetSystemKeys() (output *models.SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys"
	op := &request.Operation{
		Name:       "GetSystemKeys",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SystemKeys{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSystemKeys - Update the system keys.
//RequestType: PUT
//Input: input *UpdateSystemKeysInput
func (s *ServerSettingsService) UpdateSystemKeys(input *UpdateSystemKeysInput) (output *models.SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys"
	op := &request.Operation{
		Name:       "UpdateSystemKeys",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SystemKeys{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//RotateSystemKeys - Rotate the system keys.
//RequestType: POST
//Input:
func (s *ServerSettingsService) RotateSystemKeys() (output *models.SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys/rotate"
	op := &request.Operation{
		Name:       "RotateSystemKeys",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.SystemKeys{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetOutBoundProvisioningSettings - Get database used for outbound provisioning
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetOutBoundProvisioningSettings() (output *models.OutboundProvisionDatabase, resp *http.Response, err error) {
	path := "/serverSettings/outboundProvisioning"
	op := &request.Operation{
		Name:       "GetOutBoundProvisioningSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OutboundProvisionDatabase{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOutBoundProvisioningSettings - Update database used for outbound provisioning
//RequestType: PUT
//Input: input *UpdateOutBoundProvisioningSettingsInput
func (s *ServerSettingsService) UpdateOutBoundProvisioningSettings(input *UpdateOutBoundProvisioningSettingsInput) (output *models.OutboundProvisionDatabase, resp *http.Response, err error) {
	path := "/serverSettings/outboundProvisioning"
	op := &request.Operation{
		Name:       "UpdateOutBoundProvisioningSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.OutboundProvisionDatabase{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateCaptchaSettingsInput struct {
	Body models.CaptchaSettings
}

type UpdateEmailServerSettingsInput struct {
	ValidationEmail string
	ValidateOnly    string

	Body models.EmailServerSettings
}

type UpdateNotificationSettingsInput struct {
	Body models.NotificationSettings
}

type UpdateOutBoundProvisioningSettingsInput struct {
	Body models.OutboundProvisionDatabase
}

type UpdateServerSettingsInput struct {
	Body models.ServerSettings
}

type UpdateSystemKeysInput struct {
	Body models.SystemKeys
}
