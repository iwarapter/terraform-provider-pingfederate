package oauthAccessTokenMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAccessTokenMappingsAPI interface {
	GetMappings() (output *models.AccessTokenMappings, resp *http.Response, err error)
	CreateMapping(input *CreateMappingInput) (output *models.AccessTokenMapping, resp *http.Response, err error)
	GetMapping(input *GetMappingInput) (output *models.AccessTokenMapping, resp *http.Response, err error)
	UpdateMapping(input *UpdateMappingInput) (output *models.AccessTokenMapping, resp *http.Response, err error)
	DeleteMapping(input *DeleteMappingInput) (output *models.ApiResult, resp *http.Response, err error)
}
