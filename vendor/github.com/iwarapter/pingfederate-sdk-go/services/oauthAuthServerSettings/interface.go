package oauthAuthServerSettings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAuthServerSettingsAPI interface {
	GetAuthorizationServerSettings() (result *models.AuthorizationServerSettings, resp *http.Response, err error)
	UpdateAuthorizationServerSettings(input *UpdateAuthorizationServerSettingsInput) (result *models.AuthorizationServerSettings, resp *http.Response, err error)
	AddCommonScope(input *AddCommonScopeInput) (result *models.ScopeEntry, resp *http.Response, err error)
	GetCommonScope(input *GetCommonScopeInput) (result *models.ScopeEntry, resp *http.Response, err error)
	UpdateCommonScope(input *UpdateCommonScopeInput) (result *models.ScopeEntry, resp *http.Response, err error)
	RemoveCommonScope(input *RemoveCommonScopeInput) (result *models.ApiResult, resp *http.Response, err error)
	AddCommonScopeGroup(input *AddCommonScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error)
	GetCommonScopeGroup(input *GetCommonScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error)
	UpdateCommonScopeGroup(input *UpdateCommonScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error)
	RemoveCommonScopeGroup(input *RemoveCommonScopeGroupInput) (result *models.ApiResult, resp *http.Response, err error)
	AddExclusiveScope(input *AddExclusiveScopeInput) (result *models.ScopeEntry, resp *http.Response, err error)
	GetExclusiveScope(input *GetExclusiveScopeInput) (result *models.ScopeEntry, resp *http.Response, err error)
	UpdateExclusiveScope(input *UpdateExclusiveScopeInput) (result *models.ScopeEntry, resp *http.Response, err error)
	RemoveExclusiveScope(input *RemoveExclusiveScopeInput) (result *models.ApiResult, resp *http.Response, err error)
	AddExclusiveScopeGroup(input *AddExclusiveScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error)
	GetExclusiveScopeGroup(input *GetExclusiveScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error)
	UpdateExclusiveScopeGroups(input *UpdateExclusiveScopeGroupsInput) (result *models.ScopeGroupEntry, resp *http.Response, err error)
	RemoveExclusiveScopeGroup(input *RemoveExclusiveScopeGroupInput) (result *models.ApiResult, resp *http.Response, err error)
}
