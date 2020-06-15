package keyPairsSigning

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsSigningAPI interface {
	GetKeyPairs() (result *models.KeyPairViews, resp *http.Response, err error)
	ImportKeyPair(input *ImportKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error)
	CreateKeyPair(input *CreateKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error)
	GetKeyPair(input *GetKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error)
	DeleteKeyPair(input *DeleteKeyPairInput) (result *models.ApiResult, resp *http.Response, err error)
	ExportCsr(input *ExportCsrInput) (result *string, resp *http.Response, err error)
	ImportCsrResponse(input *ImportCsrResponseInput) (result *models.KeyPairView, resp *http.Response, err error)
	ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error)
	ExportCertificateFile(input *ExportCertificateFileInput) (result *string, resp *http.Response, err error)
	GetRotationSettings(input *GetRotationSettingsInput) (result *models.KeyPairRotationSettings, resp *http.Response, err error)
	UpdateRotationSettings(input *UpdateRotationSettingsInput) (result *models.KeyPairRotationSettings, resp *http.Response, err error)
	DeleteKeyPairRotationSettings(input *DeleteKeyPairRotationSettingsInput) (result *models.ApiResult, resp *http.Response, err error)
}
