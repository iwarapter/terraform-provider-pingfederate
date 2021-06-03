package incomingProxySettings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IncomingProxySettingsAPI interface {
	GetIncomingProxySettings() (output *models.IncomingProxySettings, resp *http.Response, err error)
	GetIncomingProxySettingsWithContext(ctx context.Context) (output *models.IncomingProxySettings, resp *http.Response, err error)

	UpdateIncomingProxySettings(input *UpdateIncomingProxySettingsInput) (output *models.IncomingProxySettings, resp *http.Response, err error)
	UpdateIncomingProxySettingsWithContext(ctx context.Context, input *UpdateIncomingProxySettingsInput) (output *models.IncomingProxySettings, resp *http.Response, err error)
}
