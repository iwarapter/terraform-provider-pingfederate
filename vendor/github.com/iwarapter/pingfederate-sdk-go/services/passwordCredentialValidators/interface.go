package passwordCredentialValidators

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type PasswordCredentialValidatorsAPI interface {
	GetPasswordCredentialValidatorDescriptors() (output *models.PasswordCredentialValidatorDescriptors, resp *http.Response, err error)
	GetPasswordCredentialValidatorDescriptor(input *GetPasswordCredentialValidatorDescriptorInput) (output *models.PasswordCredentialValidatorDescriptor, resp *http.Response, err error)
	GetPasswordCredentialValidators() (output *models.PasswordCredentialValidators, resp *http.Response, err error)
	CreatePasswordCredentialValidator(input *CreatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)
	GetPasswordCredentialValidator(input *GetPasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)
	UpdatePasswordCredentialValidator(input *UpdatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)
	DeletePasswordCredentialValidator(input *DeletePasswordCredentialValidatorInput) (output *models.ApiResult, resp *http.Response, err error)
}
