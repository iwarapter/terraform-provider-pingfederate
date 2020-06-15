package oauthAuthServerSettings

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAuthServerSettingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthAuthServerSettingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthAuthServerSettingsService {

	return &OauthAuthServerSettingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetAuthorizationServerSettings - Get the Authorization Server Settings.
//RequestType: GET
//Input:
func (s *OauthAuthServerSettingsService) GetAuthorizationServerSettings() (result *models.AuthorizationServerSettings, resp *http.Response, err error) {
	path := "/oauth/authServerSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateAuthorizationServerSettings - Update the Authorization Server Settings.
//RequestType: PUT
//Input: input *UpdateAuthorizationServerSettingsInput
func (s *OauthAuthServerSettingsService) UpdateAuthorizationServerSettings(input *UpdateAuthorizationServerSettingsInput) (result *models.AuthorizationServerSettings, resp *http.Response, err error) {
	path := "/oauth/authServerSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//AddCommonScope - Add a new common scope.
//RequestType: POST
//Input: input *AddCommonScopeInput
func (s *OauthAuthServerSettingsService) AddCommonScope(input *AddCommonScopeInput) (result *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetCommonScope - Get an existing common scope.
//RequestType: GET
//Input: input *GetCommonScopeInput
func (s *OauthAuthServerSettingsService) GetCommonScope(input *GetCommonScopeInput) (result *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateCommonScope - Update an existing common scope.
//RequestType: PUT
//Input: input *UpdateCommonScopeInput
func (s *OauthAuthServerSettingsService) UpdateCommonScope(input *UpdateCommonScopeInput) (result *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RemoveCommonScope - Remove an existing common scope.
//RequestType: DELETE
//Input: input *RemoveCommonScopeInput
func (s *OauthAuthServerSettingsService) RemoveCommonScope(input *RemoveCommonScopeInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//AddCommonScopeGroup - Create a new common scope group.
//RequestType: POST
//Input: input *AddCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) AddCommonScopeGroup(input *AddCommonScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetCommonScopeGroup - Get an existing common scope group.
//RequestType: GET
//Input: input *GetCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) GetCommonScopeGroup(input *GetCommonScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateCommonScopeGroup - Update an existing common scope group.
//RequestType: PUT
//Input: input *UpdateCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) UpdateCommonScopeGroup(input *UpdateCommonScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RemoveCommonScopeGroup - Remove an existing common scope group.
//RequestType: DELETE
//Input: input *RemoveCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) RemoveCommonScopeGroup(input *RemoveCommonScopeGroupInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//AddExclusiveScope - Add a new exclusive scope.
//RequestType: POST
//Input: input *AddExclusiveScopeInput
func (s *OauthAuthServerSettingsService) AddExclusiveScope(input *AddExclusiveScopeInput) (result *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetExclusiveScope - Get an existing exclusive scope.
//RequestType: GET
//Input: input *GetExclusiveScopeInput
func (s *OauthAuthServerSettingsService) GetExclusiveScope(input *GetExclusiveScopeInput) (result *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateExclusiveScope - Update an existing exclusive scope.
//RequestType: PUT
//Input: input *UpdateExclusiveScopeInput
func (s *OauthAuthServerSettingsService) UpdateExclusiveScope(input *UpdateExclusiveScopeInput) (result *models.ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RemoveExclusiveScope - Remove an existing exclusive scope.
//RequestType: DELETE
//Input: input *RemoveExclusiveScopeInput
func (s *OauthAuthServerSettingsService) RemoveExclusiveScope(input *RemoveExclusiveScopeInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//AddExclusiveScopeGroup - Create a new exclusive scope group.
//RequestType: POST
//Input: input *AddExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) AddExclusiveScopeGroup(input *AddExclusiveScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetExclusiveScopeGroup - Get an existing exclusive scope group.
//RequestType: GET
//Input: input *GetExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) GetExclusiveScopeGroup(input *GetExclusiveScopeGroupInput) (result *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateExclusiveScopeGroups - Update an existing exclusive scope group.
//RequestType: PUT
//Input: input *UpdateExclusiveScopeGroupsInput
func (s *OauthAuthServerSettingsService) UpdateExclusiveScopeGroups(input *UpdateExclusiveScopeGroupsInput) (result *models.ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RemoveExclusiveScopeGroup - Remove an existing exclusive scope group.
//RequestType: DELETE
//Input: input *RemoveExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) RemoveExclusiveScopeGroup(input *RemoveExclusiveScopeGroupInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type AddCommonScopeInput struct {
	Body models.ScopeEntry
}

type AddCommonScopeGroupInput struct {
	Body models.ScopeGroupEntry
}

type AddExclusiveScopeInput struct {
	Body models.ScopeEntry
}

type AddExclusiveScopeGroupInput struct {
	Body models.ScopeGroupEntry
}

type GetCommonScopeInput struct {
	Name string
}

type GetCommonScopeGroupInput struct {
	Name string
}

type GetExclusiveScopeInput struct {
	Name string
}

type GetExclusiveScopeGroupInput struct {
	Name string
}

type RemoveCommonScopeInput struct {
	Name string
}

type RemoveCommonScopeGroupInput struct {
	Name string
}

type RemoveExclusiveScopeInput struct {
	Name string
}

type RemoveExclusiveScopeGroupInput struct {
	Name string
}

type UpdateAuthorizationServerSettingsInput struct {
	Body models.AuthorizationServerSettings
}

type UpdateCommonScopeInput struct {
	Body models.ScopeEntry
	Name string
}

type UpdateCommonScopeGroupInput struct {
	Body models.ScopeGroupEntry
	Name string
}

type UpdateExclusiveScopeInput struct {
	Body models.ScopeEntry
	Name string
}

type UpdateExclusiveScopeGroupsInput struct {
	Body models.ScopeGroupEntry
	Name string
}
