package keyPairsSslClient

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsSslClientAPI interface {
	GetKeyPairs() (output *models.KeyPairViews, resp *http.Response, err error)
	GetKeyPairsWithContext(ctx context.Context) (output *models.KeyPairViews, resp *http.Response, err error)

	ImportKeyPair(input *ImportKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	ImportKeyPairWithContext(ctx context.Context, input *ImportKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)

	CreateKeyPair(input *CreateKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	CreateKeyPairWithContext(ctx context.Context, input *CreateKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)

	GetKeyPair(input *GetKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	GetKeyPairWithContext(ctx context.Context, input *GetKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)

	DeleteKeyPair(input *DeleteKeyPairInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteKeyPairWithContext(ctx context.Context, input *DeleteKeyPairInput) (output *models.ApiResult, resp *http.Response, err error)

	ExportCsr(input *ExportCsrInput) (output *string, resp *http.Response, err error)
	ExportCsrWithContext(ctx context.Context, input *ExportCsrInput) (output *string, resp *http.Response, err error)

	ImportCsrResponse(input *ImportCsrResponseInput) (output *models.KeyPairView, resp *http.Response, err error)
	ImportCsrResponseWithContext(ctx context.Context, input *ImportCsrResponseInput) (output *models.KeyPairView, resp *http.Response, err error)

	ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error)
	ExportPKCS12FileWithContext(ctx context.Context, input *ExportPKCS12FileInput) (resp *http.Response, err error)

	ExportCertificateFile(input *ExportCertificateFileInput) (output *string, resp *http.Response, err error)
	ExportCertificateFileWithContext(ctx context.Context, input *ExportCertificateFileInput) (output *string, resp *http.Response, err error)
}
