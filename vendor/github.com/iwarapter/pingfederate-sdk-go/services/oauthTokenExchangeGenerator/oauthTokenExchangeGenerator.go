package oauthTokenExchangeGenerator

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeGeneratorService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthTokenExchangeGeneratorService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthTokenExchangeGeneratorService {

	return &OauthTokenExchangeGeneratorService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSettings - Get general OAuth 2.0 Token Exchange Generator settings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeGeneratorService) GetSettings() (result *models.TokenExchangeGeneratorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/settings"
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

//UpdateSettings - Update general OAuth 2.0 Token Exchange Generator settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthTokenExchangeGeneratorService) UpdateSettings(input *UpdateSettingsInput) (result *models.TokenExchangeGeneratorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetGroups - Get list of OAuth 2.0 Token Exchange Generator groups.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeGeneratorService) GetGroups() (result *models.TokenExchangeGeneratorGroups, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups"
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

//CreateGroup - Create a new OAuth 2.0 Token Exchange Generator group.
//RequestType: POST
//Input: input *CreateGroupInput
func (s *OauthTokenExchangeGeneratorService) CreateGroup(input *CreateGroupInput) (result *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetGroup - Find an OAuth 2.0 Token Exchange Generator group by ID.
//RequestType: GET
//Input: input *GetGroupInput
func (s *OauthTokenExchangeGeneratorService) GetGroup(input *GetGroupInput) (result *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
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

//UpdateGroup - Update an OAuth 2.0 Token Exchange Generator group.
//RequestType: PUT
//Input: input *UpdateGroupInput
func (s *OauthTokenExchangeGeneratorService) UpdateGroup(input *UpdateGroupInput) (result *models.TokenExchangeGeneratorGroup, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteGroup - Delete an OAuth 2.0 Token Exchange Generator group.
//RequestType: DELETE
//Input: input *DeleteGroupInput
func (s *OauthTokenExchangeGeneratorService) DeleteGroup(input *DeleteGroupInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/generator/groups/{id}"
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

type CreateGroupInput struct {
	Body models.TokenExchangeGeneratorGroup

	BypassExternalValidation *bool
}

type DeleteGroupInput struct {
	Id string
}

type GetGroupInput struct {
	Id string
}

type UpdateGroupInput struct {
	Body models.TokenExchangeGeneratorGroup
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.TokenExchangeGeneratorSettings

	BypassExternalValidation *bool
}
