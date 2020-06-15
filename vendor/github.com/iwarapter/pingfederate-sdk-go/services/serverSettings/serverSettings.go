package serverSettings

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ServerSettingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the ServerSettingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *ServerSettingsService {

	return &ServerSettingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetServerSettings - Gets the server settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetServerSettings() (result *models.ServerSettings, resp *http.Response, err error) {
	path := "/serverSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateServerSettings - Update the server settings.
//RequestType: PUT
//Input: input *UpdateServerSettingsInput
func (s *ServerSettingsService) UpdateServerSettings(input *UpdateServerSettingsInput) (result *models.ServerSettings, resp *http.Response, err error) {
	path := "/serverSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetNotificationSettings - Gets the notification settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetNotificationSettings() (result *models.NotificationSettings, resp *http.Response, err error) {
	path := "/serverSettings/notifications"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateNotificationSettings - Update the notification settings.
//RequestType: PUT
//Input: input *UpdateNotificationSettingsInput
func (s *ServerSettingsService) UpdateNotificationSettings(input *UpdateNotificationSettingsInput) (result *models.NotificationSettings, resp *http.Response, err error) {
	path := "/serverSettings/notifications"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetEmailServerSettings - (Deprecated) Gets the email server settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetEmailServerSettings() (result *models.EmailServerSettings, resp *http.Response, err error) {
	path := "/serverSettings/emailServer"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateEmailServerSettings - (Deprecated) Update the email server settings
//RequestType: PUT
//Input: input *UpdateEmailServerSettingsInput
func (s *ServerSettingsService) UpdateEmailServerSettings(input *UpdateEmailServerSettingsInput) (result *models.EmailServerSettings, resp *http.Response, err error) {
	path := "/serverSettings/emailServer"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	q := rel.Query()
	if input.ValidationEmail != "" {
		q.Set("validationEmail", input.ValidationEmail)
	}
	if input.ValidateOnly != "" {
		q.Set("validateOnly", input.ValidateOnly)
	}
	rel.RawQuery = q.Encode()
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetCaptchaSettings - Gets the CAPTCHA settings.
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetCaptchaSettings() (result *models.CaptchaSettings, resp *http.Response, err error) {
	path := "/serverSettings/captchaSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateCaptchaSettings - Update the CAPTCHA settings.
//RequestType: PUT
//Input: input *UpdateCaptchaSettingsInput
func (s *ServerSettingsService) UpdateCaptchaSettings(input *UpdateCaptchaSettingsInput) (result *models.CaptchaSettings, resp *http.Response, err error) {
	path := "/serverSettings/captchaSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetSystemKeys - Get the system keys.
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetSystemKeys() (result *models.SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateSystemKeys - Update the system keys.
//RequestType: PUT
//Input: input *UpdateSystemKeysInput
func (s *ServerSettingsService) UpdateSystemKeys(input *UpdateSystemKeysInput) (result *models.SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RotateSystemKeys - Rotate the system keys.
//RequestType: POST
//Input:
func (s *ServerSettingsService) RotateSystemKeys() (result *models.SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys/rotate"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetOutBoundProvisioningSettings - Get database used for outbound provisioning
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetOutBoundProvisioningSettings() (result *models.OutboundProvisionDatabase, resp *http.Response, err error) {
	path := "/serverSettings/outboundProvisioning"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateOutBoundProvisioningSettings - Update database used for outbound provisioning
//RequestType: PUT
//Input: input *UpdateOutBoundProvisioningSettingsInput
func (s *ServerSettingsService) UpdateOutBoundProvisioningSettings(input *UpdateOutBoundProvisioningSettingsInput) (result *models.OutboundProvisionDatabase, resp *http.Response, err error) {
	path := "/serverSettings/outboundProvisioning"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

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
