package oauthAuthenticationPolicyContractMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAuthenticationPolicyContractMappingsAPI interface {
	GetApcMappings() (result *models.ApcToPersistentGrantMappings, resp *http.Response, err error)
	CreateApcMapping(input *CreateApcMappingInput) (result *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	GetApcMapping(input *GetApcMappingInput) (result *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	UpdateApcMapping(input *UpdateApcMappingInput) (result *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	DeleteApcMapping(input *DeleteApcMappingInput) (result *models.ApiResult, resp *http.Response, err error)
}
