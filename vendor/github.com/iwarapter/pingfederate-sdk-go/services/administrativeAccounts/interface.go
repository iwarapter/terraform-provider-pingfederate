package administrativeAccounts

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AdministrativeAccountsAPI interface {
	GetAccounts() (output *models.AdministrativeAccounts, resp *http.Response, err error)
	AddAccount(input *AddAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error)
	GetAccount(input *GetAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error)
	UpdateAccount(input *UpdateAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error)
	DeleteAccount(input *DeleteAccountInput) (output *models.ApiResult, resp *http.Response, err error)
	ResetPassword(input *ResetPasswordInput) (output *models.UserCredentials, resp *http.Response, err error)
	ChangePassword(input *ChangePasswordInput) (output *models.UserCredentials, resp *http.Response, err error)
}
