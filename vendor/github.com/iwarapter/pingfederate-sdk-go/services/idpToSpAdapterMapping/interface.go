package idpToSpAdapterMapping

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpToSpAdapterMappingAPI interface {
	GetIdpToSpAdapterMappings() (output *models.IdpToSpAdapterMappings, resp *http.Response, err error)
	GetIdpToSpAdapterMappingsWithContext(ctx context.Context) (output *models.IdpToSpAdapterMappings, resp *http.Response, err error)

	CreateIdpToSpAdapterMapping(input *CreateIdpToSpAdapterMappingInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	CreateIdpToSpAdapterMappingWithContext(ctx context.Context, input *CreateIdpToSpAdapterMappingInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)

	GetIdpToSpAdapterMappingsById(input *GetIdpToSpAdapterMappingsByIdInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	GetIdpToSpAdapterMappingsByIdWithContext(ctx context.Context, input *GetIdpToSpAdapterMappingsByIdInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)

	UpdateIdpToSpAdapterMapping(input *UpdateIdpToSpAdapterMappingInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)
	UpdateIdpToSpAdapterMappingWithContext(ctx context.Context, input *UpdateIdpToSpAdapterMappingInput) (output *models.IdpToSpAdapterMapping, resp *http.Response, err error)

	DeleteIdpToSpAdapterMappingsById(input *DeleteIdpToSpAdapterMappingsByIdInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteIdpToSpAdapterMappingsByIdWithContext(ctx context.Context, input *DeleteIdpToSpAdapterMappingsByIdInput) (output *models.ApiResult, resp *http.Response, err error)
}
