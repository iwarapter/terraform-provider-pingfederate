package spAuthenticationPolicyContractMappings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpAuthenticationPolicyContractMappingsAPI interface {
	GetApcToSpAdapterMappings() (output *models.ApcToSpAdapterMappings, resp *http.Response, err error)
	GetApcToSpAdapterMappingsWithContext(ctx context.Context) (output *models.ApcToSpAdapterMappings, resp *http.Response, err error)

	CreateApcToSpAdapterMapping(input *CreateApcToSpAdapterMappingInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	CreateApcToSpAdapterMappingWithContext(ctx context.Context, input *CreateApcToSpAdapterMappingInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)

	GetApcToSpAdapterMappingById(input *GetApcToSpAdapterMappingByIdInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	GetApcToSpAdapterMappingByIdWithContext(ctx context.Context, input *GetApcToSpAdapterMappingByIdInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)

	UpdateApcToSpAdapterMappingById(input *UpdateApcToSpAdapterMappingByIdInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	UpdateApcToSpAdapterMappingByIdWithContext(ctx context.Context, input *UpdateApcToSpAdapterMappingByIdInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)

	DeleteApcToSpAdapterMappingById(input *DeleteApcToSpAdapterMappingByIdInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteApcToSpAdapterMappingByIdWithContext(ctx context.Context, input *DeleteApcToSpAdapterMappingByIdInput) (output *models.ApiResult, resp *http.Response, err error)
}
