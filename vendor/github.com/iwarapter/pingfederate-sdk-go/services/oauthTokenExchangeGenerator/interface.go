package oauthTokenExchangeGenerator

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeGeneratorAPI interface {
	GetSettings() (result *models.TokenExchangeGeneratorSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.TokenExchangeGeneratorSettings, resp *http.Response, err error)
	GetGroups() (result *models.TokenExchangeGeneratorGroups, resp *http.Response, err error)
	CreateGroup(input *CreateGroupInput) (result *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	GetGroup(input *GetGroupInput) (result *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	UpdateGroup(input *UpdateGroupInput) (result *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	DeleteGroup(input *DeleteGroupInput) (result *models.ApiResult, resp *http.Response, err error)
}
