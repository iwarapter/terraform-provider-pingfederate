package passwordCredentialValidators

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type PasswordCredentialValidatorsAPI interface {
	GetPasswordCredentialValidatorDescriptors() (output *models.PasswordCredentialValidatorDescriptors, resp *http.Response, err error)
	GetPasswordCredentialValidatorDescriptorsWithContext(ctx context.Context) (output *models.PasswordCredentialValidatorDescriptors, resp *http.Response, err error)

	GetPasswordCredentialValidatorDescriptor(input *GetPasswordCredentialValidatorDescriptorInput) (output *models.PasswordCredentialValidatorDescriptor, resp *http.Response, err error)
	GetPasswordCredentialValidatorDescriptorWithContext(ctx context.Context, input *GetPasswordCredentialValidatorDescriptorInput) (output *models.PasswordCredentialValidatorDescriptor, resp *http.Response, err error)

	GetPasswordCredentialValidators() (output *models.PasswordCredentialValidators, resp *http.Response, err error)
	GetPasswordCredentialValidatorsWithContext(ctx context.Context) (output *models.PasswordCredentialValidators, resp *http.Response, err error)

	CreatePasswordCredentialValidator(input *CreatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)
	CreatePasswordCredentialValidatorWithContext(ctx context.Context, input *CreatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)

	GetPasswordCredentialValidator(input *GetPasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)
	GetPasswordCredentialValidatorWithContext(ctx context.Context, input *GetPasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)

	UpdatePasswordCredentialValidator(input *UpdatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)
	UpdatePasswordCredentialValidatorWithContext(ctx context.Context, input *UpdatePasswordCredentialValidatorInput) (output *models.PasswordCredentialValidator, resp *http.Response, err error)

	DeletePasswordCredentialValidator(input *DeletePasswordCredentialValidatorInput) (output *models.ApiResult, resp *http.Response, err error)
	DeletePasswordCredentialValidatorWithContext(ctx context.Context, input *DeletePasswordCredentialValidatorInput) (output *models.ApiResult, resp *http.Response, err error)
}
