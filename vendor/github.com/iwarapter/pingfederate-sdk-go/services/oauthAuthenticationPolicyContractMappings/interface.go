package oauthAuthenticationPolicyContractMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAuthenticationPolicyContractMappingsAPI interface {
	GetApcMappings() (output *models.ApcToPersistentGrantMappings, resp *http.Response, err error)
	CreateApcMapping(input *CreateApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	GetApcMapping(input *GetApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	UpdateApcMapping(input *UpdateApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	DeleteApcMapping(input *DeleteApcMappingInput) (output *models.ApiResult, resp *http.Response, err error)
}
