package oauthAccessTokenManagers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAccessTokenManagersService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthAccessTokenManagersService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthAccessTokenManagersService {

	return &OauthAccessTokenManagersService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSettings - Get general access token management settings.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetSettings() (result *models.AccessTokenManagementSettings, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/settings"
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

//UpdateSettings - Update general access token management settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthAccessTokenManagersService) UpdateSettings(input *UpdateSettingsInput) (result *models.AccessTokenManagementSettings, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/settings"
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

//GetTokenManagerDescriptors - Get the list of available token management plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetTokenManagerDescriptors() (result *models.AccessTokenManagerDescriptors, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/descriptors"
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

//GetTokenManagerDescriptor - Get the description of a token management plugin descriptor.
//RequestType: GET
//Input: input *GetTokenManagerDescriptorInput
func (s *OauthAccessTokenManagersService) GetTokenManagerDescriptor(input *GetTokenManagerDescriptorInput) (result *models.AccessTokenManagerDescriptor, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/descriptors/{id}"
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

//GetTokenManagers - Get a list of all token management plugin instances.
//RequestType: GET
//Input:
func (s *OauthAccessTokenManagersService) GetTokenManagers() (result *models.AccessTokenManagers, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers"
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

//CreateTokenManager - Create a token management plugin instance.
//RequestType: POST
//Input: input *CreateTokenManagerInput
func (s *OauthAccessTokenManagersService) CreateTokenManager(input *CreateTokenManagerInput) (result *models.AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers"
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

//GetTokenManager - Get a specific token management plugin instance.
//RequestType: GET
//Input: input *GetTokenManagerInput
func (s *OauthAccessTokenManagersService) GetTokenManager(input *GetTokenManagerInput) (result *models.AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
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

//UpdateTokenManager - Update a token management plugin instance.
//RequestType: PUT
//Input: input *UpdateTokenManagerInput
func (s *OauthAccessTokenManagersService) UpdateTokenManager(input *UpdateTokenManagerInput) (result *models.AccessTokenManager, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
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

//DeleteTokenManager - Delete a token management plugin instance.
//RequestType: DELETE
//Input: input *DeleteTokenManagerInput
func (s *OauthAccessTokenManagersService) DeleteTokenManager(input *DeleteTokenManagerInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/accessTokenManagers/{id}"
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

type CreateTokenManagerInput struct {
	Body models.AccessTokenManager
}

type DeleteTokenManagerInput struct {
	Id string
}

type GetTokenManagerInput struct {
	Id string
}

type GetTokenManagerDescriptorInput struct {
	Id string
}

type UpdateSettingsInput struct {
	Body models.AccessTokenManagementSettings
}

type UpdateTokenManagerInput struct {
	Body models.AccessTokenManager
	Id   string
}
