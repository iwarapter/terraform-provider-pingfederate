package redirectValidation

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type RedirectValidationAPI interface {
	GetRedirectValidationSettings() (output *models.RedirectValidationSettings, resp *http.Response, err error)
	UpdateRedirectValidationSettings(input *UpdateRedirectValidationSettingsInput) (output *models.RedirectValidationSettings, resp *http.Response, err error)
}
