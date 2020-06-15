package spAuthenticationPolicyContractMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpAuthenticationPolicyContractMappingsAPI interface {
	GetApcToSpAdapterMappings() (result *models.ApcToSpAdapterMappings, resp *http.Response, err error)
	CreateApcToSpAdapterMapping(input *CreateApcToSpAdapterMappingInput) (result *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	GetApcToSpAdapterMappingById(input *GetApcToSpAdapterMappingByIdInput) (result *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	UpdateApcToSpAdapterMappingById(input *UpdateApcToSpAdapterMappingByIdInput) (result *models.ApcToSpAdapterMapping, resp *http.Response, err error)
	DeleteApcToSpAdapterMappingById(input *DeleteApcToSpAdapterMappingByIdInput) (result *models.ApiResult, resp *http.Response, err error)
}
