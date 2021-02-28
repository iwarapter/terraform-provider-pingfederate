package oauthClientSettings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientSettingsAPI interface {
	GetClientSettings() (output *models.ClientSettings, resp *http.Response, err error)
	GetClientSettingsWithContext(ctx context.Context) (output *models.ClientSettings, resp *http.Response, err error)

	UpdateClientSettings(input *UpdateClientSettingsInput) (output *models.ClientSettings, resp *http.Response, err error)
	UpdateClientSettingsWithContext(ctx context.Context, input *UpdateClientSettingsInput) (output *models.ClientSettings, resp *http.Response, err error)
}
