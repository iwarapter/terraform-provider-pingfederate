package idpToSpAdapterMapping

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpToSpAdapterMappingAPI interface {
	GetIdpToSpAdapterMappings() (result *models.IdpToSpAdapterMappings, resp *http.Response, err error)
	CreateIdpToSpAdapterMapping(input *CreateIdpToSpAdapterMappingInput) (result *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	GetIdpToSpAdapterMappingsById(input *GetIdpToSpAdapterMappingsByIdInput) (result *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	UpdateIdpToSpAdapterMapping(input *UpdateIdpToSpAdapterMappingInput) (result *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	DeleteIdpToSpAdapterMappingsById(input *DeleteIdpToSpAdapterMappingsByIdInput) (result *models.ApiResult, resp *http.Response, err error)
}
