package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthAuthServerSettingsService service

//GetAuthorizationServerSettings - Get the Authorization Server Settings.
//RequestType: GET
//Input:
func (s *OauthAuthServerSettingsService) GetAuthorizationServerSettings() (result *AuthorizationServerSettings, resp *http.Response, err error) {
	path := "/oauth/authServerSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateAuthorizationServerSettings - Update the Authorization Server Settings.
//RequestType: PUT
//Input: input *UpdateAuthorizationServerSettingsInput
func (s *OauthAuthServerSettingsService) UpdateAuthorizationServerSettings(input *UpdateAuthorizationServerSettingsInput) (result *AuthorizationServerSettings, resp *http.Response, err error) {
	path := "/oauth/authServerSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//AddCommonScope - Add a new common scope.
//RequestType: POST
//Input: input *AddCommonScopeInput
func (s *OauthAuthServerSettingsService) AddCommonScope(input *AddCommonScopeInput) (result *ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetCommonScope - Get an existing common scope.
//RequestType: GET
//Input: input *GetCommonScopeInput
func (s *OauthAuthServerSettingsService) GetCommonScope(input *GetCommonScopeInput) (result *ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateCommonScope - Update an existing common scope.
//RequestType: PUT
//Input: input *UpdateCommonScopeInput
func (s *OauthAuthServerSettingsService) UpdateCommonScope(input *UpdateCommonScopeInput) (result *ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RemoveCommonScope - Remove an existing common scope.
//RequestType: DELETE
//Input: input *RemoveCommonScopeInput
func (s *OauthAuthServerSettingsService) RemoveCommonScope(input *RemoveCommonScopeInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//AddCommonScopeGroup - Create a new common scope group.
//RequestType: POST
//Input: input *AddCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) AddCommonScopeGroup(input *AddCommonScopeGroupInput) (result *ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetCommonScopeGroup - Get an existing common scope group.
//RequestType: GET
//Input: input *GetCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) GetCommonScopeGroup(input *GetCommonScopeGroupInput) (result *ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateCommonScopeGroup - Update an existing common scope group.
//RequestType: PUT
//Input: input *UpdateCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) UpdateCommonScopeGroup(input *UpdateCommonScopeGroupInput) (result *ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RemoveCommonScopeGroup - Remove an existing common scope group.
//RequestType: DELETE
//Input: input *RemoveCommonScopeGroupInput
func (s *OauthAuthServerSettingsService) RemoveCommonScopeGroup(input *RemoveCommonScopeGroupInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/commonScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//AddExclusiveScope - Add a new exclusive scope.
//RequestType: POST
//Input: input *AddExclusiveScopeInput
func (s *OauthAuthServerSettingsService) AddExclusiveScope(input *AddExclusiveScopeInput) (result *ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetExclusiveScope - Get an existing exclusive scope.
//RequestType: GET
//Input: input *GetExclusiveScopeInput
func (s *OauthAuthServerSettingsService) GetExclusiveScope(input *GetExclusiveScopeInput) (result *ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateExclusiveScope - Update an existing exclusive scope.
//RequestType: PUT
//Input: input *UpdateExclusiveScopeInput
func (s *OauthAuthServerSettingsService) UpdateExclusiveScope(input *UpdateExclusiveScopeInput) (result *ScopeEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RemoveExclusiveScope - Remove an existing exclusive scope.
//RequestType: DELETE
//Input: input *RemoveExclusiveScopeInput
func (s *OauthAuthServerSettingsService) RemoveExclusiveScope(input *RemoveExclusiveScopeInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopes/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//AddExclusiveScopeGroup - Create a new exclusive scope group.
//RequestType: POST
//Input: input *AddExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) AddExclusiveScopeGroup(input *AddExclusiveScopeGroupInput) (result *ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetExclusiveScopeGroup - Get an existing exclusive scope group.
//RequestType: GET
//Input: input *GetExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) GetExclusiveScopeGroup(input *GetExclusiveScopeGroupInput) (result *ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateExclusiveScopeGroups - Update an existing exclusive scope group.
//RequestType: PUT
//Input: input *UpdateExclusiveScopeGroupsInput
func (s *OauthAuthServerSettingsService) UpdateExclusiveScopeGroups(input *UpdateExclusiveScopeGroupsInput) (result *ScopeGroupEntry, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//RemoveExclusiveScopeGroup - Remove an existing exclusive scope group.
//RequestType: DELETE
//Input: input *RemoveExclusiveScopeGroupInput
func (s *OauthAuthServerSettingsService) RemoveExclusiveScopeGroup(input *RemoveExclusiveScopeGroupInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/authServerSettings/scopes/exclusiveScopeGroups/{name}"
	path = strings.Replace(path, "{name}", input.Name, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
