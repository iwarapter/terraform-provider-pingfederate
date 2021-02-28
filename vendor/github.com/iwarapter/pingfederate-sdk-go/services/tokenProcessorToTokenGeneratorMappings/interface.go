package tokenProcessorToTokenGeneratorMappings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type TokenProcessorToTokenGeneratorMappingsAPI interface {
	GetTokenToTokenMappings() (output *models.TokenToTokenMappings, resp *http.Response, err error)
	GetTokenToTokenMappingsWithContext(ctx context.Context) (output *models.TokenToTokenMappings, resp *http.Response, err error)

	CreateTokenToTokenMapping(input *CreateTokenToTokenMappingInput) (output *models.TokenToTokenMapping, resp *http.Response, err error)
	CreateTokenToTokenMappingWithContext(ctx context.Context, input *CreateTokenToTokenMappingInput) (output *models.TokenToTokenMapping, resp *http.Response, err error)

	GetTokenToTokenMappingById(input *GetTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error)
	GetTokenToTokenMappingByIdWithContext(ctx context.Context, input *GetTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error)

	UpdateTokenToTokenMappingById(input *UpdateTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error)
	UpdateTokenToTokenMappingByIdWithContext(ctx context.Context, input *UpdateTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error)

	DeleteTokenToTokenMappingById(input *DeleteTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error)
	DeleteTokenToTokenMappingByIdWithContext(ctx context.Context, input *DeleteTokenToTokenMappingByIdInput) (output *models.TokenToTokenMapping, resp *http.Response, err error)
}
