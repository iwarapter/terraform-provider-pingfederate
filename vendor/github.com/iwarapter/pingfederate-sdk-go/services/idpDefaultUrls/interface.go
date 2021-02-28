package idpDefaultUrls

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpDefaultUrlsAPI interface {
	GetDefaultUrl() (output *models.IdpDefaultUrl, resp *http.Response, err error)
	GetDefaultUrlWithContext(ctx context.Context) (output *models.IdpDefaultUrl, resp *http.Response, err error)

	UpdateDefaultUrlSettings(input *UpdateDefaultUrlSettingsInput) (output *models.IdpDefaultUrl, resp *http.Response, err error)
	UpdateDefaultUrlSettingsWithContext(ctx context.Context, input *UpdateDefaultUrlSettingsInput) (output *models.IdpDefaultUrl, resp *http.Response, err error)
}
