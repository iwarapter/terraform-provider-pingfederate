package keyPairsSslClient

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsSslClientAPI interface {
	GetKeyPairs() (output *models.KeyPairViews, resp *http.Response, err error)
	ImportKeyPair(input *ImportKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	CreateKeyPair(input *CreateKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	GetKeyPair(input *GetKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	DeleteKeyPair(input *DeleteKeyPairInput) (output *models.ApiResult, resp *http.Response, err error)
	ExportCsr(input *ExportCsrInput) (output *string, resp *http.Response, err error)
	ImportCsrResponse(input *ImportCsrResponseInput) (output *models.KeyPairView, resp *http.Response, err error)
	ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error)
	ExportCertificateFile(input *ExportCertificateFileInput) (output *string, resp *http.Response, err error)
}
