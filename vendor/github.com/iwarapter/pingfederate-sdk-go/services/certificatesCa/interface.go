package certificatesCa

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type CertificatesCaAPI interface {
	GetTrustedCAs() (result *models.CertViews, resp *http.Response, err error)
	GetTrustedCert(input *GetTrustedCertInput) (result *models.CertView, resp *http.Response, err error)
	DeleteTrustedCA(input *DeleteTrustedCAInput) (result *models.ApiResult, resp *http.Response, err error)
	ImportTrustedCA(input *ImportTrustedCAInput) (result *models.CertView, resp *http.Response, err error)
	ExportCertificateFile(input *ExportCertificateFileInput) (result *string, resp *http.Response, err error)
}
