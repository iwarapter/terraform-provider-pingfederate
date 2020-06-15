package spIdpConnections

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpIdpConnectionsAPI interface {
	GetConnections(input *GetConnectionsInput) (result *models.IdpConnections, resp *http.Response, err error)
	CreateConnection(input *CreateConnectionInput) (result *models.IdpConnection, resp *http.Response, err error)
	GetConnection(input *GetConnectionInput) (result *models.IdpConnection, resp *http.Response, err error)
	UpdateConnection(input *UpdateConnectionInput) (result *models.IdpConnection, resp *http.Response, err error)
	DeleteConnection(input *DeleteConnectionInput) (result *models.ApiResult, resp *http.Response, err error)
	GetSigningSettings(input *GetSigningSettingsInput) (result *models.SigningSettings, resp *http.Response, err error)
	UpdateSigningSettings(input *UpdateSigningSettingsInput) (result *models.SigningSettings, resp *http.Response, err error)
	AddConnectionCert(input *AddConnectionCertInput) (result *models.ConnectionCert, resp *http.Response, err error)
	GetConnectionCerts(input *GetConnectionCertsInput) (result *models.ConnectionCerts, resp *http.Response, err error)
	UpdateConnectionCerts(input *UpdateConnectionCertsInput) (result *models.ConnectionCerts, resp *http.Response, err error)
	GetDecryptionKeys(input *GetDecryptionKeysInput) (result *models.DecryptionKeys, resp *http.Response, err error)
	UpdateDecryptionKeys(input *UpdateDecryptionKeysInput) (result *models.DecryptionKeys, resp *http.Response, err error)
}
