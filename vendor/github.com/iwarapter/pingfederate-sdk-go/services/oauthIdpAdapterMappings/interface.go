package oauthIdpAdapterMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthIdpAdapterMappingsAPI interface {
	GetIdpAdapterMappings() (result *models.IdpAdapterMappings, resp *http.Response, err error)
	CreateIdpAdapterMapping(input *CreateIdpAdapterMappingInput) (result *models.IdpAdapterMapping, resp *http.Response, err error)
	GetIdpAdapterMapping(input *GetIdpAdapterMappingInput) (result *models.IdpAdapterMapping, resp *http.Response, err error)
	UpdateIdpAdapterMapping(input *UpdateIdpAdapterMappingInput) (result *models.IdpAdapterMapping, resp *http.Response, err error)
	DeleteIdpAdapterMapping(input *DeleteIdpAdapterMappingInput) (result *models.ApiResult, resp *http.Response, err error)
}
