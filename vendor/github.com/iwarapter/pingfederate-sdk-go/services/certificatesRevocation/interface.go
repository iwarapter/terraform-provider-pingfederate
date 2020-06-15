package certificatesRevocation

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type CertificatesRevocationAPI interface {
	GetRevocationSettings() (result *models.CertificateRevocationSettings, resp *http.Response, err error)
	UpdateRevocationSettings(input *UpdateRevocationSettingsInput) (result *models.CertificateRevocationSettings, resp *http.Response, err error)
	GetOcspCertificates() (result *models.CertViews, resp *http.Response, err error)
	ImportOcspCertificate(input *ImportOcspCertificateInput) (result *models.CertView, resp *http.Response, err error)
	GetOcspCertificateById(input *GetOcspCertificateByIdInput) (result *models.CertView, resp *http.Response, err error)
	DeleteOcspCertificateById(input *DeleteOcspCertificateByIdInput) (result *models.ApiResult, resp *http.Response, err error)
}
