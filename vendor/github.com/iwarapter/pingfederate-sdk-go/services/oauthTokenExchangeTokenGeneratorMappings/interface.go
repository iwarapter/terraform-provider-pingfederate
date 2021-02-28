package oauthTokenExchangeTokenGeneratorMappings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeTokenGeneratorMappingsAPI interface {
	GetTokenGeneratorMappings() (output *models.ProcessorPolicyToGeneratorMappings, resp *http.Response, err error)
	GetTokenGeneratorMappingsWithContext(ctx context.Context) (output *models.ProcessorPolicyToGeneratorMappings, resp *http.Response, err error)

	CreateTokenGeneratorMapping(input *CreateTokenGeneratorMappingInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
	CreateTokenGeneratorMappingWithContext(ctx context.Context, input *CreateTokenGeneratorMappingInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)

	GetTokenGeneratorMappingById(input *GetTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
	GetTokenGeneratorMappingByIdWithContext(ctx context.Context, input *GetTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)

	UpdateTokenGeneratorMappingById(input *UpdateTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
	UpdateTokenGeneratorMappingByIdWithContext(ctx context.Context, input *UpdateTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)

	DeleteTokenGeneratorMappingById(input *DeleteTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
	DeleteTokenGeneratorMappingByIdWithContext(ctx context.Context, input *DeleteTokenGeneratorMappingByIdInput) (output *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error)
}
