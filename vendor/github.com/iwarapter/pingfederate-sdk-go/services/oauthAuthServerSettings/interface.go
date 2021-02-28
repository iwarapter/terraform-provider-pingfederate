package oauthAuthServerSettings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAuthServerSettingsAPI interface {
	GetAuthorizationServerSettings() (output *models.AuthorizationServerSettings, resp *http.Response, err error)
	GetAuthorizationServerSettingsWithContext(ctx context.Context) (output *models.AuthorizationServerSettings, resp *http.Response, err error)

	UpdateAuthorizationServerSettings(input *UpdateAuthorizationServerSettingsInput) (output *models.AuthorizationServerSettings, resp *http.Response, err error)
	UpdateAuthorizationServerSettingsWithContext(ctx context.Context, input *UpdateAuthorizationServerSettingsInput) (output *models.AuthorizationServerSettings, resp *http.Response, err error)

	AddCommonScope(input *AddCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	AddCommonScopeWithContext(ctx context.Context, input *AddCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)

	GetCommonScope(input *GetCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	GetCommonScopeWithContext(ctx context.Context, input *GetCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)

	UpdateCommonScope(input *UpdateCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	UpdateCommonScopeWithContext(ctx context.Context, input *UpdateCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)

	RemoveCommonScope(input *RemoveCommonScopeInput) (output *models.ApiResult, resp *http.Response, err error)
	RemoveCommonScopeWithContext(ctx context.Context, input *RemoveCommonScopeInput) (output *models.ApiResult, resp *http.Response, err error)

	AddCommonScopeGroup(input *AddCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	AddCommonScopeGroupWithContext(ctx context.Context, input *AddCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)

	GetCommonScopeGroup(input *GetCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	GetCommonScopeGroupWithContext(ctx context.Context, input *GetCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)

	UpdateCommonScopeGroup(input *UpdateCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	UpdateCommonScopeGroupWithContext(ctx context.Context, input *UpdateCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)

	RemoveCommonScopeGroup(input *RemoveCommonScopeGroupInput) (output *models.ApiResult, resp *http.Response, err error)
	RemoveCommonScopeGroupWithContext(ctx context.Context, input *RemoveCommonScopeGroupInput) (output *models.ApiResult, resp *http.Response, err error)

	AddExclusiveScope(input *AddExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	AddExclusiveScopeWithContext(ctx context.Context, input *AddExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)

	GetExclusiveScope(input *GetExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	GetExclusiveScopeWithContext(ctx context.Context, input *GetExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)

	UpdateExclusiveScope(input *UpdateExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	UpdateExclusiveScopeWithContext(ctx context.Context, input *UpdateExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)

	RemoveExclusiveScope(input *RemoveExclusiveScopeInput) (output *models.ApiResult, resp *http.Response, err error)
	RemoveExclusiveScopeWithContext(ctx context.Context, input *RemoveExclusiveScopeInput) (output *models.ApiResult, resp *http.Response, err error)

	AddExclusiveScopeGroup(input *AddExclusiveScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	AddExclusiveScopeGroupWithContext(ctx context.Context, input *AddExclusiveScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)

	GetExclusiveScopeGroup(input *GetExclusiveScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	GetExclusiveScopeGroupWithContext(ctx context.Context, input *GetExclusiveScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)

	UpdateExclusiveScopeGroups(input *UpdateExclusiveScopeGroupsInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	UpdateExclusiveScopeGroupsWithContext(ctx context.Context, input *UpdateExclusiveScopeGroupsInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)

	RemoveExclusiveScopeGroup(input *RemoveExclusiveScopeGroupInput) (output *models.ApiResult, resp *http.Response, err error)
	RemoveExclusiveScopeGroupWithContext(ctx context.Context, input *RemoveExclusiveScopeGroupInput) (output *models.ApiResult, resp *http.Response, err error)
}
