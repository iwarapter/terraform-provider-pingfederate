package idpTokenProcessors

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpTokenProcessorsAPI interface {
	GetTokenProcessorDescriptors() (output *models.TokenProcessorDescriptors, resp *http.Response, err error)
	GetTokenProcessorDescriptorsWithContext(ctx context.Context) (output *models.TokenProcessorDescriptors, resp *http.Response, err error)

	GetTokenProcessorDescriptorsById(input *GetTokenProcessorDescriptorsByIdInput) (output *models.TokenProcessorDescriptor, resp *http.Response, err error)
	GetTokenProcessorDescriptorsByIdWithContext(ctx context.Context, input *GetTokenProcessorDescriptorsByIdInput) (output *models.TokenProcessorDescriptor, resp *http.Response, err error)

	GetTokenProcessors() (output *models.TokenProcessors, resp *http.Response, err error)
	GetTokenProcessorsWithContext(ctx context.Context) (output *models.TokenProcessors, resp *http.Response, err error)

	CreateTokenProcessor(input *CreateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)
	CreateTokenProcessorWithContext(ctx context.Context, input *CreateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)

	GetTokenProcessor(input *GetTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)
	GetTokenProcessorWithContext(ctx context.Context, input *GetTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)

	UpdateTokenProcessor(input *UpdateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)
	UpdateTokenProcessorWithContext(ctx context.Context, input *UpdateTokenProcessorInput) (output *models.TokenProcessor, resp *http.Response, err error)

	DeleteTokenProcessor(input *DeleteTokenProcessorInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteTokenProcessorWithContext(ctx context.Context, input *DeleteTokenProcessorInput) (output *models.ApiResult, resp *http.Response, err error)
}
