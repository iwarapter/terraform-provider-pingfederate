package idpTokenProcessors

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpTokenProcessorsAPI interface {
	GetTokenProcessorDescriptors() (output *models.TokenProcessorDescriptors, resp *http.Response, err error)
	GetTokenProcessorDescriptorsById(input *GetTokenProcessorDescriptorsByIdInput) (output *models.TokenProcessorDescriptor, resp *http.Response, err error)
	GetTokenProcessors() (output *models.TokenProcessors, resp *http.Response, err error)
	CreateTokenProcessor(input *CreateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)
	GetTokenProcessor(input *GetTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)
	UpdateTokenProcessor(input *UpdateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)
	DeleteTokenProcessor(input *DeleteTokenProcessorInput) (output *models.ApiResult, resp *http.Response, err error)
}
