package keyPairsSigning

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsSigningAPI interface {
	GetKeyPairs() (output *models.KeyPairViews, resp *http.Response, err error)
	ImportKeyPair(input *ImportKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	CreateKeyPair(input *CreateKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	GetKeyPair(input *GetKeyPairInput) (output *models.KeyPairView, resp *http.Response, err error)
	DeleteKeyPair(input *DeleteKeyPairInput) (output *models.ApiResult, resp *http.Response, err error)
	ExportCsr(input *ExportCsrInput) (output *string, resp *http.Response, err error)
	ImportCsrResponse(input *ImportCsrResponseInput) (output *models.KeyPairView, resp *http.Response, err error)
	ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error)
	ExportCertificateFile(input *ExportCertificateFileInput) (output *string, resp *http.Response, err error)
	GetRotationSettings(input *GetRotationSettingsInput) (output *models.KeyPairRotationSettings, resp *http.Response, err error)
	UpdateRotationSettings(input *UpdateRotationSettingsInput) (output *models.KeyPairRotationSettings, resp *http.Response, err error)
	DeleteKeyPairRotationSettings(input *DeleteKeyPairRotationSettingsInput) (output *models.ApiResult, resp *http.Response, err error)
}
