package certificatesCa

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type CertificatesCaAPI interface {
	GetTrustedCAs() (output *models.CertViews, resp *http.Response, err error)
	GetTrustedCert(input *GetTrustedCertInput) (output *models.CertView, resp *http.Response, err error)
	DeleteTrustedCA(input *DeleteTrustedCAInput) (output *models.ApiResult, resp *http.Response, err error)
	ImportTrustedCA(input *ImportTrustedCAInput) (output *models.CertView, resp *http.Response, err error)
	ExportCertificateFile(input *ExportCertificateFileInput) (output *string, resp *http.Response, err error)
}
