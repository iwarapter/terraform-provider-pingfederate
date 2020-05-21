package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AuthenticationApiService service

//GetAuthenticationApiSettings - Get the Authentication API settings.
//RequestType: GET
//Input:
func (s *AuthenticationApiService) GetAuthenticationApiSettings() (result *AuthnApiSettings, resp *http.Response, err error) {
	path := "/authenticationApi/settings"
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

//UpdateAuthenticationApiSettings - Set the Authentication API settings.
//RequestType: PUT
//Input: input *UpdateAuthenticationApiSettingsInput
func (s *AuthenticationApiService) UpdateAuthenticationApiSettings(input *UpdateAuthenticationApiSettingsInput) (result *AuthnApiSettings, resp *http.Response, err error) {
	path := "/authenticationApi/settings"
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

//GetAuthenticationApiApplications - Get the collection of Authentication API Applications.
//RequestType: GET
//Input:
func (s *AuthenticationApiService) GetAuthenticationApiApplications() (result *AuthnApiApplications, resp *http.Response, err error) {
	path := "/authenticationApi/applications"
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

//CreateApplication - Create a new Authentication API Application.
//RequestType: POST
//Input: input *CreateApplicationInput
func (s *AuthenticationApiService) CreateApplication(input *CreateApplicationInput) (result *AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications"
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

//GetApplication - Find Authentication API Application by ID.
//RequestType: GET
//Input: input *GetApplicationInput
func (s *AuthenticationApiService) GetApplication(input *GetApplicationInput) (result *AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
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

//UpdateApplication - Update an Authentication API Application.
//RequestType: PUT
//Input: input *UpdateApplicationInput
func (s *AuthenticationApiService) UpdateApplication(input *UpdateApplicationInput) (result *AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
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

//DeleteApplication - Delete an Authentication API Application.
//RequestType: DELETE
//Input: input *DeleteApplicationInput
func (s *AuthenticationApiService) DeleteApplication(input *DeleteApplicationInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
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
