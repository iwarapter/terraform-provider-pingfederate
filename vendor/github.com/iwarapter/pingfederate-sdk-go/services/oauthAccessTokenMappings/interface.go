package oauthAccessTokenMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAccessTokenMappingsAPI interface {
	GetMappings() (result *models.AccessTokenMappings, resp *http.Response, err error)
	CreateMapping(input *CreateMappingInput) (result *models.AccessTokenMapping, resp *http.Response, err error)
	GetMapping(input *GetMappingInput) (result *models.AccessTokenMapping, resp *http.Response, err error)
	UpdateMapping(input *UpdateMappingInput) (result *models.AccessTokenMapping, resp *http.Response, err error)
	DeleteMapping(input *DeleteMappingInput) (result *models.ApiResult, resp *http.Response, err error)
}
