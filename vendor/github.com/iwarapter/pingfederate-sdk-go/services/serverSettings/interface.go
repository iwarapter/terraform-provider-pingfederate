package serverSettings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ServerSettingsAPI interface {
	GetServerSettings() (output *models.ServerSettings, resp *http.Response, err error)
	GetServerSettingsWithContext(ctx context.Context) (output *models.ServerSettings, resp *http.Response, err error)

	UpdateServerSettings(input *UpdateServerSettingsInput) (output *models.ServerSettings, resp *http.Response, err error)
	UpdateServerSettingsWithContext(ctx context.Context, input *UpdateServerSettingsInput) (output *models.ServerSettings, resp *http.Response, err error)

	GetNotificationSettings() (output *models.NotificationSettings, resp *http.Response, err error)
	GetNotificationSettingsWithContext(ctx context.Context) (output *models.NotificationSettings, resp *http.Response, err error)

	UpdateNotificationSettings(input *UpdateNotificationSettingsInput) (output *models.NotificationSettings, resp *http.Response, err error)
	UpdateNotificationSettingsWithContext(ctx context.Context, input *UpdateNotificationSettingsInput) (output *models.NotificationSettings, resp *http.Response, err error)

	GetEmailServerSettings() (output *models.EmailServerSettings, resp *http.Response, err error)
	GetEmailServerSettingsWithContext(ctx context.Context) (output *models.EmailServerSettings, resp *http.Response, err error)

	UpdateEmailServerSettings(input *UpdateEmailServerSettingsInput) (output *models.EmailServerSettings, resp *http.Response, err error)
	UpdateEmailServerSettingsWithContext(ctx context.Context, input *UpdateEmailServerSettingsInput) (output *models.EmailServerSettings, resp *http.Response, err error)

	GetCaptchaSettings() (output *models.CaptchaSettings, resp *http.Response, err error)
	GetCaptchaSettingsWithContext(ctx context.Context) (output *models.CaptchaSettings, resp *http.Response, err error)

	UpdateCaptchaSettings(input *UpdateCaptchaSettingsInput) (output *models.CaptchaSettings, resp *http.Response, err error)
	UpdateCaptchaSettingsWithContext(ctx context.Context, input *UpdateCaptchaSettingsInput) (output *models.CaptchaSettings, resp *http.Response, err error)

	UpdateSystemKeys(input *UpdateSystemKeysInput) (output *models.SystemKeys, resp *http.Response, err error)
	UpdateSystemKeysWithContext(ctx context.Context, input *UpdateSystemKeysInput) (output *models.SystemKeys, resp *http.Response, err error)

	GetSystemKeys() (output *models.SystemKeys, resp *http.Response, err error)
	GetSystemKeysWithContext(ctx context.Context) (output *models.SystemKeys, resp *http.Response, err error)

	RotateSystemKeys() (output *models.SystemKeys, resp *http.Response, err error)
	RotateSystemKeysWithContext(ctx context.Context) (output *models.SystemKeys, resp *http.Response, err error)

	GetOutBoundProvisioningSettings() (output *models.OutboundProvisionDatabase, resp *http.Response, err error)
	GetOutBoundProvisioningSettingsWithContext(ctx context.Context) (output *models.OutboundProvisionDatabase, resp *http.Response, err error)

	UpdateOutBoundProvisioningSettings(input *UpdateOutBoundProvisioningSettingsInput) (output *models.OutboundProvisionDatabase, resp *http.Response, err error)
	UpdateOutBoundProvisioningSettingsWithContext(ctx context.Context, input *UpdateOutBoundProvisioningSettingsInput) (output *models.OutboundProvisionDatabase, resp *http.Response, err error)
}
