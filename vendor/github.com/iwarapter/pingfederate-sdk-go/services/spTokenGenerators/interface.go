package spTokenGenerators

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpTokenGeneratorsAPI interface {
	GetTokenGeneratorDescriptors() (output *models.TokenGeneratorDescriptors, resp *http.Response, err error)
	GetTokenGeneratorDescriptorsById(input *GetTokenGeneratorDescriptorsByIdInput) (output *models.TokenGeneratorDescriptor, resp *http.Response, err error)
	GetTokenGenerators() (output *models.TokenGenerators, resp *http.Response, err error)
	CreateTokenGenerator(input *CreateTokenGeneratorInput) (output *models.TokenGenerator, resp *http.Response, err error)
	GetTokenGenerator(input *GetTokenGeneratorInput) (output *models.TokenGenerator, resp *http.Response, err error)
	UpdateTokenGenerator(input *UpdateTokenGeneratorInput) (output *models.TokenGenerator, resp *http.Response, err error)
	DeleteTokenGenerator(input *DeleteTokenGeneratorInput) (output *models.ApiResult, resp *http.Response, err error)
}
