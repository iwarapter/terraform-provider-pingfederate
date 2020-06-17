package certificatesRevocation

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type CertificatesRevocationAPI interface {
	GetRevocationSettings() (output *models.CertificateRevocationSettings, resp *http.Response, err error)
	UpdateRevocationSettings(input *UpdateRevocationSettingsInput) (output *models.CertificateRevocationSettings, resp *http.Response, err error)
	GetOcspCertificates() (output *models.CertViews, resp *http.Response, err error)
	ImportOcspCertificate(input *ImportOcspCertificateInput) (output *models.CertView, resp *http.Response, err error)
	GetOcspCertificateById(input *GetOcspCertificateByIdInput) (output *models.CertView, resp *http.Response, err error)
	DeleteOcspCertificateById(input *DeleteOcspCertificateByIdInput) (output *models.ApiResult, resp *http.Response, err error)
}
