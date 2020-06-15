package oauthResourceOwnerCredentialsMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthResourceOwnerCredentialsMappingsAPI interface {
	GetResourceOwnerCredentialsMappings() (result *models.ResourceOwnerCredentialsMappings, resp *http.Response, err error)
	CreateResourceOwnerCredentialsMapping(input *CreateResourceOwnerCredentialsMappingInput) (result *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	GetResourceOwnerCredentialsMapping(input *GetResourceOwnerCredentialsMappingInput) (result *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	UpdateResourceOwnerCredentialsMapping(input *UpdateResourceOwnerCredentialsMappingInput) (result *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	DeleteResourceOwnerCredentialsMapping(input *DeleteResourceOwnerCredentialsMappingInput) (result *models.ApiResult, resp *http.Response, err error)
}
