package oauthTokenExchangeGenerator

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeGeneratorAPI interface {
	GetSettings() (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error)
	GetSettingsWithContext(ctx context.Context) (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error)

	UpdateSettings(input *UpdateSettingsInput) (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error)
	UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.TokenExchangeGeneratorSettings, resp *http.Response, err error)

	GetGroups() (output *models.TokenExchangeGeneratorGroups, resp *http.Response, err error)
	GetGroupsWithContext(ctx context.Context) (output *models.TokenExchangeGeneratorGroups, resp *http.Response, err error)

	CreateGroup(input *CreateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	CreateGroupWithContext(ctx context.Context, input *CreateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)

	GetGroup(input *GetGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	GetGroupWithContext(ctx context.Context, input *GetGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)

	UpdateGroup(input *UpdateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)
	UpdateGroupWithContext(ctx context.Context, input *UpdateGroupInput) (output *models.TokenExchangeGeneratorGroup, resp *http.Response, err error)

	DeleteGroup(input *DeleteGroupInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteGroupWithContext(ctx context.Context, input *DeleteGroupInput) (output *models.ApiResult, resp *http.Response, err error)
}
