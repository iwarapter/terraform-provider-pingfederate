package oauthAuthServerSettings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAuthServerSettingsAPI interface {
	GetAuthorizationServerSettings() (output *models.AuthorizationServerSettings, resp *http.Response, err error)
	UpdateAuthorizationServerSettings(input *UpdateAuthorizationServerSettingsInput) (output *models.AuthorizationServerSettings, resp *http.Response, err error)
	AddCommonScope(input *AddCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	GetCommonScope(input *GetCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	UpdateCommonScope(input *UpdateCommonScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	RemoveCommonScope(input *RemoveCommonScopeInput) (output *models.ApiResult, resp *http.Response, err error)
	AddCommonScopeGroup(input *AddCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	GetCommonScopeGroup(input *GetCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	UpdateCommonScopeGroup(input *UpdateCommonScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	RemoveCommonScopeGroup(input *RemoveCommonScopeGroupInput) (output *models.ApiResult, resp *http.Response, err error)
	AddExclusiveScope(input *AddExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	GetExclusiveScope(input *GetExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	UpdateExclusiveScope(input *UpdateExclusiveScopeInput) (output *models.ScopeEntry, resp *http.Response, err error)
	RemoveExclusiveScope(input *RemoveExclusiveScopeInput) (output *models.ApiResult, resp *http.Response, err error)
	AddExclusiveScopeGroup(input *AddExclusiveScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	GetExclusiveScopeGroup(input *GetExclusiveScopeGroupInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	UpdateExclusiveScopeGroups(input *UpdateExclusiveScopeGroupsInput) (output *models.ScopeGroupEntry, resp *http.Response, err error)
	RemoveExclusiveScopeGroup(input *RemoveExclusiveScopeGroupInput) (output *models.ApiResult, resp *http.Response, err error)
}
