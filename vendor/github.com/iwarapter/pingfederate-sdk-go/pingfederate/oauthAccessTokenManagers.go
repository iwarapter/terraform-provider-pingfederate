package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthAccessTokenManagersService service

//GetSettings - Get general access token management settings.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetSettings() (result *AccessTokenManagementSettings, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/settings"
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

//UpdateSettings - Update general access token management settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthAccessTokenManagersService) UpdateSettings(input *UpdateSettingsInput) (result *AccessTokenManagementSettings, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/settings"
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

//GetTokenManagerDescriptors - Get the list of available token management plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetTokenManagerDescriptors() (result *AccessTokenManagerDescriptors, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/descriptors"
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

//GetTokenManagerDescriptor - Get the description of a token management plugin descriptor.
//RequestType: GET
//Input: input *GetTokenManagerDescriptorInput
func (s *OauthAccessTokenManagersService) GetTokenManagerDescriptor(input *GetTokenManagerDescriptorInput) (result *AccessTokenManagerDescriptor, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/descriptors/{id}"
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

//GetTokenManagers - Get a list of all token management plugin instances.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetTokenManagers() (result *AccessTokenManagers, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers"
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

//CreateTokenManager - Create a token management plugin instance.
//RequestType: POST
//Input: input *CreateTokenManagerInput
func (s *OauthAccessTokenManagersService) CreateTokenManager(input *CreateTokenManagerInput) (result *AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers"
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

//GetTokenManager - Get a specific token management plugin instance.
//RequestType: GET
//Input: input *GetTokenManagerInput
func (s *OauthAccessTokenManagersService) GetTokenManager(input *GetTokenManagerInput) (result *AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
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

//UpdateTokenManager - Update a token management plugin instance.
//RequestType: PUT
//Input: input *UpdateTokenManagerInput
func (s *OauthAccessTokenManagersService) UpdateTokenManager(input *UpdateTokenManagerInput) (result *AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//DeleteTokenManager - Delete a token management plugin instance.
//RequestType: DELETE
//Input: input *DeleteTokenManagerInput
func (s *OauthAccessTokenManagersService) DeleteTokenManager(input *DeleteTokenManagerInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
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
