package oauthIdpAdapterMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthIdpAdapterMappingsAPI interface {
	GetIdpAdapterMappings() (output *models.IdpAdapterMappings, resp *http.Response, err error)
	CreateIdpAdapterMapping(input *CreateIdpAdapterMappingInput) (output *models.IdpAdapterMapping, resp *http.Response, err error)
	GetIdpAdapterMapping(input *GetIdpAdapterMappingInput) (output *models.IdpAdapterMapping, resp *http.Response, err error)
	UpdateIdpAdapterMapping(input *UpdateIdpAdapterMappingInput) (output *models.IdpAdapterMapping, resp *http.Response, err error)
	DeleteIdpAdapterMapping(input *DeleteIdpAdapterMappingInput) (output *models.ApiResult, resp *http.Response, err error)
}
