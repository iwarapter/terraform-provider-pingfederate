package oauthClientSettings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientSettingsAPI interface {
	GetClientSettings() (output *models.ClientSettings, resp *http.Response, err error)
	UpdateClientSettings(input *UpdateClientSettingsInput) (output *models.ClientSettings, resp *http.Response, err error)
}
