package license

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type LicenseAPI interface {
	GetLicenseAgreement() (output *models.LicenseAgreementInfo, resp *http.Response, err error)
	UpdateLicenseAgreement(input *UpdateLicenseAgreementInput) (output *models.LicenseAgreementInfo, resp *http.Response, err error)
	GetLicense() (output *models.LicenseView, resp *http.Response, err error)
	UpdateLicense(input *UpdateLicenseInput) (output *models.LicenseView, resp *http.Response, err error)
}
