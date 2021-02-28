package kerberosRealms

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KerberosRealmsAPI interface {
	GetKerberosRealmSettings() (output *models.KerberosRealmsSettings, resp *http.Response, err error)
	GetKerberosRealmSettingsWithContext(ctx context.Context) (output *models.KerberosRealmsSettings, resp *http.Response, err error)

	UpdateSettings(input *UpdateSettingsInput) (output *models.KerberosRealmsSettings, resp *http.Response, err error)
	UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.KerberosRealmsSettings, resp *http.Response, err error)

	GetKerberosRealms() (output *models.KerberosRealms, resp *http.Response, err error)
	GetKerberosRealmsWithContext(ctx context.Context) (output *models.KerberosRealms, resp *http.Response, err error)

	CreateKerberosRealm(input *CreateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)
	CreateKerberosRealmWithContext(ctx context.Context, input *CreateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)

	GetKerberosRealm(input *GetKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)
	GetKerberosRealmWithContext(ctx context.Context, input *GetKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)

	UpdateKerberosRealm(input *UpdateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)
	UpdateKerberosRealmWithContext(ctx context.Context, input *UpdateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)

	DeleteKerberosRealm(input *DeleteKerberosRealmInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteKerberosRealmWithContext(ctx context.Context, input *DeleteKerberosRealmInput) (output *models.ApiResult, resp *http.Response, err error)
}
