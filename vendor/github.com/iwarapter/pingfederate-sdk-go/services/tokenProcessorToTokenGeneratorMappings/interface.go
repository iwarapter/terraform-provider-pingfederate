package tokenProcessorToTokenGeneratorMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type TokenProcessorToTokenGeneratorMappingsAPI interface {
	GetTokenToTokenMappings() (result *models.TokenToTokenMappings, resp *http.Response, err error)
	CreateTokenToTokenMapping(input *CreateTokenToTokenMappingInput) (result *models.TokenToTokenMapping, resp *http.Response, err error)
	GetTokenToTokenMappingById(input *GetTokenToTokenMappingByIdInput) (result *models.TokenToTokenMapping, resp *http.Response, err error)
	UpdateTokenToTokenMappingById(input *UpdateTokenToTokenMappingByIdInput) (result *models.TokenToTokenMapping, resp *http.Response, err error)
	DeleteTokenToTokenMappingById(input *DeleteTokenToTokenMappingByIdInput) (result *models.TokenToTokenMapping, resp *http.Response, err error)
}
