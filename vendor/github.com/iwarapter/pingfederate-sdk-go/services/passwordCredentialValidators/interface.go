package passwordCredentialValidators

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type PasswordCredentialValidatorsAPI interface {
	GetPasswordCredentialValidatorDescriptors() (result *models.PasswordCredentialValidatorDescriptors, resp *http.Response, err error)
	GetPasswordCredentialValidatorDescriptor(input *GetPasswordCredentialValidatorDescriptorInput) (result *models.PasswordCredentialValidatorDescriptor, resp *http.Response, err error)
	GetPasswordCredentialValidators() (result *models.PasswordCredentialValidators, resp *http.Response, err error)
	CreatePasswordCredentialValidator(input *CreatePasswordCredentialValidatorInput) (result *models.PasswordCredentialValidator, resp *http.Response, err error)
	GetPasswordCredentialValidator(input *GetPasswordCredentialValidatorInput) (result *models.PasswordCredentialValidator, resp *http.Response, err error)
	UpdatePasswordCredentialValidator(input *UpdatePasswordCredentialValidatorInput) (result *models.PasswordCredentialValidator, resp *http.Response, err error)
	DeletePasswordCredentialValidator(input *DeletePasswordCredentialValidatorInput) (result *models.ApiResult, resp *http.Response, err error)
}
