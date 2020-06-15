package serverSettings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ServerSettingsAPI interface {
	GetServerSettings() (result *models.ServerSettings, resp *http.Response, err error)
	UpdateServerSettings(input *UpdateServerSettingsInput) (result *models.ServerSettings, resp *http.Response, err error)
	GetNotificationSettings() (result *models.NotificationSettings, resp *http.Response, err error)
	UpdateNotificationSettings(input *UpdateNotificationSettingsInput) (result *models.NotificationSettings, resp *http.Response, err error)
	GetEmailServerSettings() (result *models.EmailServerSettings, resp *http.Response, err error)
	UpdateEmailServerSettings(input *UpdateEmailServerSettingsInput) (result *models.EmailServerSettings, resp *http.Response, err error)
	GetCaptchaSettings() (result *models.CaptchaSettings, resp *http.Response, err error)
	UpdateCaptchaSettings(input *UpdateCaptchaSettingsInput) (result *models.CaptchaSettings, resp *http.Response, err error)
	GetSystemKeys() (result *models.SystemKeys, resp *http.Response, err error)
	UpdateSystemKeys(input *UpdateSystemKeysInput) (result *models.SystemKeys, resp *http.Response, err error)
	RotateSystemKeys() (result *models.SystemKeys, resp *http.Response, err error)
	GetOutBoundProvisioningSettings() (result *models.OutboundProvisionDatabase, resp *http.Response, err error)
	UpdateOutBoundProvisioningSettings(input *UpdateOutBoundProvisioningSettingsInput) (result *models.OutboundProvisionDatabase, resp *http.Response, err error)
}
