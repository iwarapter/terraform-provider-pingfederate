package oauthTokenExchangeGenerator

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeGeneratorAPI interface {
	GetSettings() (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error)
	GetGroups() (output *models.TokenExchangeGeneratorGroups, resp *http.Response, err error)
	CreateGroup(input *CreateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	GetGroup(input *GetGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	UpdateGroup(input *UpdateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	DeleteGroup(input *DeleteGroupInput) (output *models.ApiResult, resp *http.Response, err error)
}
