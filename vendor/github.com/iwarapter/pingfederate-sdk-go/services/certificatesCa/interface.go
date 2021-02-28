package certificatesCa

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type CertificatesCaAPI interface {
	GetTrustedCAs() (output *models.CertViews, resp *http.Response, err error)
	GetTrustedCAsWithContext(ctx context.Context) (output *models.CertViews, resp *http.Response, err error)

	GetTrustedCert(input *GetTrustedCertInput) (output *models.CertView, resp *http.Response, err error)
	GetTrustedCertWithContext(ctx context.Context, input *GetTrustedCertInput) (output *models.CertView, resp *http.Response, err error)

	DeleteTrustedCA(input *DeleteTrustedCAInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteTrustedCAWithContext(ctx context.Context, input *DeleteTrustedCAInput) (output *models.ApiResult, resp *http.Response, err error)

	ImportTrustedCA(input *ImportTrustedCAInput) (output *models.CertView, resp *http.Response, err error)
	ImportTrustedCAWithContext(ctx context.Context, input *ImportTrustedCAInput) (output *models.CertView, resp *http.Response, err error)

	ExportCertificateFile(input *ExportCertificateFileInput) (output *string, resp *http.Response, err error)
	ExportCertificateFileWithContext(ctx context.Context, input *ExportCertificateFileInput) (output *string, resp *http.Response, err error)
}
