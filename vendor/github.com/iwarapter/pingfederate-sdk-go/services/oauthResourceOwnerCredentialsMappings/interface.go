package oauthResourceOwnerCredentialsMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthResourceOwnerCredentialsMappingsAPI interface {
	GetResourceOwnerCredentialsMappings() (output *models.ResourceOwnerCredentialsMappings, resp *http.Response, err error)
	CreateResourceOwnerCredentialsMapping(input *CreateResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	GetResourceOwnerCredentialsMapping(input *GetResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	UpdateResourceOwnerCredentialsMapping(input *UpdateResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	DeleteResourceOwnerCredentialsMapping(input *DeleteResourceOwnerCredentialsMappingInput) (output *models.ApiResult, resp *http.Response, err error)
}
