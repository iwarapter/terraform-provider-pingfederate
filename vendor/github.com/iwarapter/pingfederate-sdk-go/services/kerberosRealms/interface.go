package kerberosRealms

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KerberosRealmsAPI interface {
	GetKerberosRealmSettings() (output *models.KerberosRealmsSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (output *models.KerberosRealmsSettings, resp *http.Response, err error)
	GetKerberosRealms() (output *models.KerberosRealms, resp *http.Response, err error)
	CreateKerberosRealm(input *CreateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)
	GetKerberosRealm(input *GetKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)
	UpdateKerberosRealm(input *UpdateKerberosRealmInput) (output *models.KerberosRealm, resp *http.Response, err error)
	DeleteKerberosRealm(input *DeleteKerberosRealmInput) (output *models.ApiResult, resp *http.Response, err error)
}
