package idpDefaultUrls

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpDefaultUrlsAPI interface {
	GetDefaultUrl() (output *models.IdpDefaultUrl, resp *http.Response, err error)
	UpdateDefaultUrlSettings(input *UpdateDefaultUrlSettingsInput) (output *models.IdpDefaultUrl, resp *http.Response, err error)
}
