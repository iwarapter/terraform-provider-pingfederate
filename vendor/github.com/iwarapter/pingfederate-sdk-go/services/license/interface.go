package license

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type LicenseAPI interface {
	GetLicenseAgreement() (result *models.LicenseAgreementInfo, resp *http.Response, err error)
	UpdateLicenseAgreement(input *UpdateLicenseAgreementInput) (result *models.LicenseAgreementInfo, resp *http.Response, err error)
	GetLicense() (result *models.LicenseView, resp *http.Response, err error)
	UpdateLicense(input *UpdateLicenseInput) (result *models.LicenseView, resp *http.Response, err error)
}
