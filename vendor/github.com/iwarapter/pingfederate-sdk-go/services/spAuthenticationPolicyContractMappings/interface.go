package spAuthenticationPolicyContractMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpAuthenticationPolicyContractMappingsAPI interface {
	GetApcToSpAdapterMappings() (output *models.ApcToSpAdapterMappings, resp *http.Response, err error)
	CreateApcToSpAdapterMapping(input *CreateApcToSpAdapterMappingInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	GetApcToSpAdapterMappingById(input *GetApcToSpAdapterMappingByIdInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	UpdateApcToSpAdapterMappingById(input *UpdateApcToSpAdapterMappingByIdInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	DeleteApcToSpAdapterMappingById(input *DeleteApcToSpAdapterMappingByIdInput) (output *models.ApiResult, resp *http.Response, err error)
}
