package spTokenGenerators

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpTokenGeneratorsAPI interface {
	GetTokenGeneratorDescriptors() (result *models.TokenGeneratorDescriptors, resp *http.Response, err error)
	GetTokenGeneratorDescriptorsById(input *GetTokenGeneratorDescriptorsByIdInput) (result *models.TokenGeneratorDescriptor, resp *http.Response, err error)
	GetTokenGenerators() (result *models.TokenGenerators, resp *http.Response, err error)
	CreateTokenGenerator(input *CreateTokenGeneratorInput) (result *models.TokenGenerator, resp *http.Response, err error)
	GetTokenGenerator(input *GetTokenGeneratorInput) (result *models.TokenGenerator, resp *http.Response, err error)
	UpdateTokenGenerator(input *UpdateTokenGeneratorInput) (result *models.TokenGenerator, resp *http.Response, err error)
	DeleteTokenGenerator(input *DeleteTokenGeneratorInput) (result *models.ApiResult, resp *http.Response, err error)
}
