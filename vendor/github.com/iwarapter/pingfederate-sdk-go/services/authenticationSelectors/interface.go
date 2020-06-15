package authenticationSelectors

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationSelectorsAPI interface {
	GetAuthenticationSelectorDescriptors() (result *models.AuthenticationSelectorDescriptors, resp *http.Response, err error)
	GetAuthenticationSelectorDescriptorsById(input *GetAuthenticationSelectorDescriptorsByIdInput) (result *models.AuthenticationSelectorDescriptor, resp *http.Response, err error)
	GetAuthenticationSelectors(input *GetAuthenticationSelectorsInput) (result *models.AuthenticationSelectors, resp *http.Response, err error)
	CreateAuthenticationSelector(input *CreateAuthenticationSelectorInput) (result *models.AuthenticationSelector, resp *http.Response, err error)
	GetAuthenticationSelector(input *GetAuthenticationSelectorInput) (result *models.AuthenticationSelector, resp *http.Response, err error)
	UpdateAuthenticationSelector(input *UpdateAuthenticationSelectorInput) (result *models.AuthenticationSelector, resp *http.Response, err error)
	DeleteAuthenticationSelector(input *DeleteAuthenticationSelectorInput) (result *models.ApiResult, resp *http.Response, err error)
}
