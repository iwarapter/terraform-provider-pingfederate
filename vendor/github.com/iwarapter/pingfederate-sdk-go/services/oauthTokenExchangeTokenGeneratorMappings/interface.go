package oauthTokenExchangeTokenGeneratorMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeTokenGeneratorMappingsAPI interface {
	GetTokenGeneratorMappings() (result *models.ProcessorPolicyToGeneratorMappings, resp *http.Response, err error)
	CreateTokenGeneratorMapping(input *CreateTokenGeneratorMappingInput) (result *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
	GetTokenGeneratorMappingById(input *GetTokenGeneratorMappingByIdInput) (result *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
	UpdateTokenGeneratorMappingById(input *UpdateTokenGeneratorMappingByIdInput) (result *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
	DeleteTokenGeneratorMappingById(input *DeleteTokenGeneratorMappingByIdInput) (result *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
}
