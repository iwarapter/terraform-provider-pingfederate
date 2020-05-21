package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthTokenExchangeGeneratorService service

//GetSettings - Get general OAuth 2.0 Token Exchange Generator settings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeGeneratorService) GetSettings() (result *TokenExchangeGeneratorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/settings"
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

//UpdateSettings - Update general OAuth 2.0 Token Exchange Generator settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthTokenExchangeGeneratorService) UpdateSettings(input *UpdateSettingsInput) (result *TokenExchangeGeneratorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetGroups - Get list of OAuth 2.0 Token Exchange Generator groups.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeGeneratorService) GetGroups() (result *TokenExchangeGeneratorGroups, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups"
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

//CreateGroup - Create a new OAuth 2.0 Token Exchange Generator group.
//RequestType: POST
//Input: input *CreateGroupInput
func (s *OauthTokenExchangeGeneratorService) CreateGroup(input *CreateGroupInput) (result *TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetGroup - Find an OAuth 2.0 Token Exchange Generator group by ID.
//RequestType: GET
//Input: input *GetGroupInput
func (s *OauthTokenExchangeGeneratorService) GetGroup(input *GetGroupInput) (result *TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateGroup - Update an OAuth 2.0 Token Exchange Generator group.
//RequestType: PUT
//Input: input *UpdateGroupInput
func (s *OauthTokenExchangeGeneratorService) UpdateGroup(input *UpdateGroupInput) (result *TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteGroup - Delete an OAuth 2.0 Token Exchange Generator group.
//RequestType: DELETE
//Input: input *DeleteGroupInput
func (s *OauthTokenExchangeGeneratorService) DeleteGroup(input *DeleteGroupInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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
