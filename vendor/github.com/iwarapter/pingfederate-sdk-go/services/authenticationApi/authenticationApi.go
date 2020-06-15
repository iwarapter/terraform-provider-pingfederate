package authenticationApi

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationApiService struct {
	Client *client.PfClient
}

// New creates a new instance of the AuthenticationApiService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *AuthenticationApiService {

	return &AuthenticationApiService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetAuthenticationApiSettings - Get the Authentication API settings.
//RequestType: GET
//Input:
func (s *AuthenticationApiService) GetAuthenticationApiSettings() (result *models.AuthnApiSettings, resp *http.Response, err error) {
	path := "/authenticationApi/settings"
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

//UpdateAuthenticationApiSettings - Set the Authentication API settings.
//RequestType: PUT
//Input: input *UpdateAuthenticationApiSettingsInput
func (s *AuthenticationApiService) UpdateAuthenticationApiSettings(input *UpdateAuthenticationApiSettingsInput) (result *models.AuthnApiSettings, resp *http.Response, err error) {
	path := "/authenticationApi/settings"
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

//GetAuthenticationApiApplications - Get the collection of Authentication API Applications.
//RequestType: GET
//Input:
func (s *AuthenticationApiService) GetAuthenticationApiApplications() (result *models.AuthnApiApplications, resp *http.Response, err error) {
	path := "/authenticationApi/applications"
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

//CreateApplication - Create a new Authentication API Application.
//RequestType: POST
//Input: input *CreateApplicationInput
func (s *AuthenticationApiService) CreateApplication(input *CreateApplicationInput) (result *models.AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications"
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

//GetApplication - Find Authentication API Application by ID.
//RequestType: GET
//Input: input *GetApplicationInput
func (s *AuthenticationApiService) GetApplication(input *GetApplicationInput) (result *models.AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateApplication - Update an Authentication API Application.
//RequestType: PUT
//Input: input *UpdateApplicationInput
func (s *AuthenticationApiService) UpdateApplication(input *UpdateApplicationInput) (result *models.AuthnApiApplication, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//DeleteApplication - Delete an Authentication API Application.
//RequestType: DELETE
//Input: input *DeleteApplicationInput
func (s *AuthenticationApiService) DeleteApplication(input *DeleteApplicationInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/authenticationApi/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type CreateApplicationInput struct {
	Body models.AuthnApiApplication
}

type DeleteApplicationInput struct {
	Id string
}

type GetApplicationInput struct {
	Id string
}

type UpdateApplicationInput struct {
	Body models.AuthnApiApplication
	Id   string
}

type UpdateAuthenticationApiSettingsInput struct {
	Body models.AuthnApiSettings
}
