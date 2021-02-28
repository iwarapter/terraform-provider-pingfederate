package spIdpConnections

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpIdpConnectionsAPI interface {
	GetConnections(input *GetConnectionsInput) (output *models.IdpConnections, resp *http.Response, err error)
	GetConnectionsWithContext(ctx context.Context, input *GetConnectionsInput) (output *models.IdpConnections, resp *http.Response, err error)

	CreateConnection(input *CreateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)
	CreateConnectionWithContext(ctx context.Context, input *CreateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)

	GetConnection(input *GetConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)
	GetConnectionWithContext(ctx context.Context, input *GetConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)

	UpdateConnection(input *UpdateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)
	UpdateConnectionWithContext(ctx context.Context, input *UpdateConnectionInput) (output *models.IdpConnection, resp *http.Response, err error)

	DeleteConnection(input *DeleteConnectionInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteConnectionWithContext(ctx context.Context, input *DeleteConnectionInput) (output *models.ApiResult, resp *http.Response, err error)

	GetSigningSettings(input *GetSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error)
	GetSigningSettingsWithContext(ctx context.Context, input *GetSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error)

	UpdateSigningSettings(input *UpdateSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error)
	UpdateSigningSettingsWithContext(ctx context.Context, input *UpdateSigningSettingsInput) (output *models.SigningSettings, resp *http.Response, err error)

	AddConnectionCert(input *AddConnectionCertInput) (output *models.ConnectionCert, resp *http.Response, err error)
	AddConnectionCertWithContext(ctx context.Context, input *AddConnectionCertInput) (output *models.ConnectionCert, resp *http.Response, err error)

	GetConnectionCerts(input *GetConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error)
	GetConnectionCertsWithContext(ctx context.Context, input *GetConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error)

	UpdateConnectionCerts(input *UpdateConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error)
	UpdateConnectionCertsWithContext(ctx context.Context, input *UpdateConnectionCertsInput) (output *models.ConnectionCerts, resp *http.Response, err error)

	GetDecryptionKeys(input *GetDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error)
	GetDecryptionKeysWithContext(ctx context.Context, input *GetDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error)

	UpdateDecryptionKeys(input *UpdateDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error)
	UpdateDecryptionKeysWithContext(ctx context.Context, input *UpdateDecryptionKeysInput) (output *models.DecryptionKeys, resp *http.Response, err error)
}
