package license

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type LicenseAPI interface {
	GetLicenseAgreement() (output *models.LicenseAgreementInfo, resp *http.Response, err error)
	GetLicenseAgreementWithContext(ctx context.Context) (output *models.LicenseAgreementInfo, resp *http.Response, err error)

	UpdateLicenseAgreement(input *UpdateLicenseAgreementInput) (output *models.LicenseAgreementInfo, resp *http.Response, err error)
	UpdateLicenseAgreementWithContext(ctx context.Context, input *UpdateLicenseAgreementInput) (output *models.LicenseAgreementInfo, resp *http.Response, err error)

	GetLicense() (output *models.LicenseView, resp *http.Response, err error)
	GetLicenseWithContext(ctx context.Context) (output *models.LicenseView, resp *http.Response, err error)

	UpdateLicense(input *UpdateLicenseInput) (output *models.LicenseView, resp *http.Response, err error)
	UpdateLicenseWithContext(ctx context.Context, input *UpdateLicenseInput) (output *models.LicenseView, resp *http.Response, err error)
}
