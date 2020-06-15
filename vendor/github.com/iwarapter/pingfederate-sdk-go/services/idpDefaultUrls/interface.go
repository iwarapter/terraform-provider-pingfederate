package idpDefaultUrls

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpDefaultUrlsAPI interface {
	GetDefaultUrl() (result *models.IdpDefaultUrl, resp *http.Response, err error)
	UpdateDefaultUrlSettings(input *UpdateDefaultUrlSettingsInput) (result *models.IdpDefaultUrl, resp *http.Response, err error)
}
