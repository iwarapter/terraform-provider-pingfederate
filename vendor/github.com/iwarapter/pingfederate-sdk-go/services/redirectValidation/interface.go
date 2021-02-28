package redirectValidation

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type RedirectValidationAPI interface {
	GetRedirectValidationSettings() (output *models.RedirectValidationSettings, resp *http.Response, err error)
	GetRedirectValidationSettingsWithContext(ctx context.Context) (output *models.RedirectValidationSettings, resp *http.Response, err error)

	UpdateRedirectValidationSettings(input *UpdateRedirectValidationSettingsInput) (output *models.RedirectValidationSettings, resp *http.Response, err error)
	UpdateRedirectValidationSettingsWithContext(ctx context.Context, input *UpdateRedirectValidationSettingsInput) (output *models.RedirectValidationSettings, resp *http.Response, err error)
}
