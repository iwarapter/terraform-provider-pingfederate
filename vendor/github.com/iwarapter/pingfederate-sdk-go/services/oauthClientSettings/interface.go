package oauthClientSettings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientSettingsAPI interface {
	GetClientSettings() (result *models.ClientSettings, resp *http.Response, err error)
	UpdateClientSettings(input *UpdateClientSettingsInput) (result *models.ClientSettings, resp *http.Response, err error)
}
