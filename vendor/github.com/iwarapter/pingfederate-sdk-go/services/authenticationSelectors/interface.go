package authenticationSelectors

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationSelectorsAPI interface {
	GetAuthenticationSelectorDescriptors() (output *models.AuthenticationSelectorDescriptors, resp *http.Response, err error)
	GetAuthenticationSelectorDescriptorsById(input *GetAuthenticationSelectorDescriptorsByIdInput) (output *models.AuthenticationSelectorDescriptor, resp *http.Response, err error)
	GetAuthenticationSelectors(input *GetAuthenticationSelectorsInput) (output *models.AuthenticationSelectors, resp *http.Response, err error)
	CreateAuthenticationSelector(input *CreateAuthenticationSelectorInput) (output *models.AuthenticationSelector, resp *http.Response, err error)
	GetAuthenticationSelector(input *GetAuthenticationSelectorInput) (output *models.AuthenticationSelector, resp *http.Response, err error)
	UpdateAuthenticationSelector(input *UpdateAuthenticationSelectorInput) (output *models.AuthenticationSelector, resp *http.Response, err error)
	DeleteAuthenticationSelector(input *DeleteAuthenticationSelectorInput) (output *models.ApiResult, resp *http.Response, err error)
}
