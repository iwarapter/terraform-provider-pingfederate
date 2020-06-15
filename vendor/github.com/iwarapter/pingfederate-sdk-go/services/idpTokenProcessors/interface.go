package idpTokenProcessors

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpTokenProcessorsAPI interface {
	GetTokenProcessorDescriptors() (result *models.TokenProcessorDescriptors, resp *http.Response, err error)
	GetTokenProcessorDescriptorsById(input *GetTokenProcessorDescriptorsByIdInput) (result *models.TokenProcessorDescriptor, resp *http.Response, err error)
	GetTokenProcessors() (result *models.TokenProcessors, resp *http.Response, err error)
	CreateTokenProcessor(input *CreateTokenProcessorInput) (result *models.TokenProcessor, resp *http.Response, err error)
	GetTokenProcessor(input *GetTokenProcessorInput) (result *models.TokenProcessor, resp *http.Response, err error)
	UpdateTokenProcessor(input *UpdateTokenProcessorInput) (result *models.TokenProcessor, resp *http.Response, err error)
	DeleteTokenProcessor(input *DeleteTokenProcessorInput) (result *models.ApiResult, resp *http.Response, err error)
}
