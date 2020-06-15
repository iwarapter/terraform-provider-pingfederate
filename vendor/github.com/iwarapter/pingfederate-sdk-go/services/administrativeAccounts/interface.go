package administrativeAccounts

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AdministrativeAccountsAPI interface {
	GetAccounts() (result *models.AdministrativeAccounts, resp *http.Response, err error)
	AddAccount(input *AddAccountInput) (result *models.AdministrativeAccount, resp *http.Response, err error)
	GetAccount(input *GetAccountInput) (result *models.AdministrativeAccount, resp *http.Response, err error)
	UpdateAccount(input *UpdateAccountInput) (result *models.AdministrativeAccount, resp *http.Response, err error)
	DeleteAccount(input *DeleteAccountInput) (result *models.ApiResult, resp *http.Response, err error)
	ResetPassword(input *ResetPasswordInput) (result *models.UserCredentials, resp *http.Response, err error)
	ChangePassword(input *ChangePasswordInput) (result *models.UserCredentials, resp *http.Response, err error)
}
