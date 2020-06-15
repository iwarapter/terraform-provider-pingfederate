package keyPairsSslServer

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsSslServerAPI interface {
	GetKeyPairs() (result *models.KeyPairViews, resp *http.Response, err error)
	ImportKeyPair(input *ImportKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error)
	CreateKeyPair(input *CreateKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error)
	GetKeyPair(input *GetKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error)
	DeleteKeyPair(input *DeleteKeyPairInput) (result *models.ApiResult, resp *http.Response, err error)
	ExportCsr(input *ExportCsrInput) (result *string, resp *http.Response, err error)
	ImportCsrResponse(input *ImportCsrResponseInput) (result *models.KeyPairView, resp *http.Response, err error)
	ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error)
	ExportCertificateFile(input *ExportCertificateFileInput) (result *string, resp *http.Response, err error)
	GetSettings() (result *models.SslServerSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.SslServerSettings, resp *http.Response, err error)
}
