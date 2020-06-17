package serverSettings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ServerSettingsAPI interface {
	GetServerSettings() (output *models.ServerSettings, resp *http.Response, err error)
	UpdateServerSettings(input *UpdateServerSettingsInput) (output *models.ServerSettings, resp *http.Response, err error)
	GetNotificationSettings() (output *models.NotificationSettings, resp *http.Response, err error)
	UpdateNotificationSettings(input *UpdateNotificationSettingsInput) (output *models.NotificationSettings, resp *http.Response, err error)
	GetEmailServerSettings() (output *models.EmailServerSettings, resp *http.Response, err error)
	UpdateEmailServerSettings(input *UpdateEmailServerSettingsInput) (output *models.EmailServerSettings, resp *http.Response, err error)
	GetCaptchaSettings() (output *models.CaptchaSettings, resp *http.Response, err error)
	UpdateCaptchaSettings(input *UpdateCaptchaSettingsInput) (output *models.CaptchaSettings, resp *http.Response, err error)
	GetSystemKeys() (output *models.SystemKeys, resp *http.Response, err error)
	UpdateSystemKeys(input *UpdateSystemKeysInput) (output *models.SystemKeys, resp *http.Response, err error)
	RotateSystemKeys() (output *models.SystemKeys, resp *http.Response, err error)
	GetOutBoundProvisioningSettings() (output *models.OutboundProvisionDatabase, resp *http.Response, err error)
	UpdateOutBoundProvisioningSettings(input *UpdateOutBoundProvisioningSettingsInput) (output *models.OutboundProvisionDatabase, resp *http.Response, err error)
}
