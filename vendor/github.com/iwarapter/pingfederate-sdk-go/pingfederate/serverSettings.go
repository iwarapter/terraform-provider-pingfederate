package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type ServerSettingsService service

//GetServerSettings - Gets the server settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetServerSettings() (result *ServerSettings, resp *http.Response, err error) {
	path := "/serverSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateServerSettings - Update the server settings.
//RequestType: PUT
//Input: input *UpdateServerSettingsInput
func (s *ServerSettingsService) UpdateServerSettings(input *UpdateServerSettingsInput) (result *ServerSettings, resp *http.Response, err error) {
	path := "/serverSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetNotificationSettings - Gets the notification settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetNotificationSettings() (result *NotificationSettings, resp *http.Response, err error) {
	path := "/serverSettings/notifications"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateNotificationSettings - Update the notification settings.
//RequestType: PUT
//Input: input *UpdateNotificationSettingsInput
func (s *ServerSettingsService) UpdateNotificationSettings(input *UpdateNotificationSettingsInput) (result *NotificationSettings, resp *http.Response, err error) {
	path := "/serverSettings/notifications"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetEmailServerSettings - (Deprecated) Gets the email server settings
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetEmailServerSettings() (result *EmailServerSettings, resp *http.Response, err error) {
	path := "/serverSettings/emailServer"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateEmailServerSettings - (Deprecated) Update the email server settings
//RequestType: PUT
//Input: input *UpdateEmailServerSettingsInput
func (s *ServerSettingsService) UpdateEmailServerSettings(input *UpdateEmailServerSettingsInput) (result *EmailServerSettings, resp *http.Response, err error) {
	path := "/serverSettings/emailServer"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.ValidationEmail != "" {
		q.Set("validationEmail", input.ValidationEmail)
	}
	if input.ValidateOnly != "" {
		q.Set("validateOnly", input.ValidateOnly)
	}
	rel.RawQuery = q.Encode()
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetCaptchaSettings - Gets the CAPTCHA settings.
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetCaptchaSettings() (result *CaptchaSettings, resp *http.Response, err error) {
	path := "/serverSettings/captchaSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateCaptchaSettings - Update the CAPTCHA settings.
//RequestType: PUT
//Input: input *UpdateCaptchaSettingsInput
func (s *ServerSettingsService) UpdateCaptchaSettings(input *UpdateCaptchaSettingsInput) (result *CaptchaSettings, resp *http.Response, err error) {
	path := "/serverSettings/captchaSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetSystemKeys - Get the system keys.
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetSystemKeys() (result *SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateSystemKeys - Update the system keys.
//RequestType: PUT
//Input: input *UpdateSystemKeysInput
func (s *ServerSettingsService) UpdateSystemKeys(input *UpdateSystemKeysInput) (result *SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RotateSystemKeys - Rotate the system keys.
//RequestType: POST
//Input:
func (s *ServerSettingsService) RotateSystemKeys() (result *SystemKeys, resp *http.Response, err error) {
	path := "/serverSettings/systemKeys/rotate"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetOutBoundProvisioningSettings - Get database used for outbound provisioning
//RequestType: GET
//Input:
func (s *ServerSettingsService) GetOutBoundProvisioningSettings() (result *OutboundProvisionDatabase, resp *http.Response, err error) {
	path := "/serverSettings/outboundProvisioning"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateOutBoundProvisioningSettings - Update database used for outbound provisioning
//RequestType: PUT
//Input: input *UpdateOutBoundProvisioningSettingsInput
func (s *ServerSettingsService) UpdateOutBoundProvisioningSettings(input *UpdateOutBoundProvisioningSettingsInput) (result *OutboundProvisionDatabase, resp *http.Response, err error) {
	path := "/serverSettings/outboundProvisioning"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
