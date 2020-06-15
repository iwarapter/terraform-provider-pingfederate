package kerberosRealms

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KerberosRealmsAPI interface {
	GetKerberosRealmSettings() (result *models.KerberosRealmsSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.KerberosRealmsSettings, resp *http.Response, err error)
	GetKerberosRealms() (result *models.KerberosRealms, resp *http.Response, err error)
	CreateKerberosRealm(input *CreateKerberosRealmInput) (result *models.KerberosRealm, resp *http.Response, err error)
	GetKerberosRealm(input *GetKerberosRealmInput) (result *models.KerberosRealm, resp *http.Response, err error)
	UpdateKerberosRealm(input *UpdateKerberosRealmInput) (result *models.KerberosRealm, resp *http.Response, err error)
	DeleteKerberosRealm(input *DeleteKerberosRealmInput) (result *models.ApiResult, resp *http.Response, err error)
}
