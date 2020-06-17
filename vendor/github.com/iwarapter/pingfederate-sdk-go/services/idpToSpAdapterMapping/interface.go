package idpToSpAdapterMapping

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpToSpAdapterMappingAPI interface {
	GetIdpToSpAdapterMappings() (output *models.IdpToSpAdapterMappings, resp *http.Response, err error)
	CreateIdpToSpAdapterMapping(input *CreateIdpToSpAdapterMappingInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	GetIdpToSpAdapterMappingsById(input *GetIdpToSpAdapterMappingsByIdInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	UpdateIdpToSpAdapterMapping(input *UpdateIdpToSpAdapterMappingInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	DeleteIdpToSpAdapterMappingsById(input *DeleteIdpToSpAdapterMappingsByIdInput) (output *models.ApiResult, resp *http.Response, err error)
}
