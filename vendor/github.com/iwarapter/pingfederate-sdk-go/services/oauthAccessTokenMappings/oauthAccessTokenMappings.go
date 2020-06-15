package oauthAccessTokenMappings

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAccessTokenMappingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthAccessTokenMappingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthAccessTokenMappingsService {

	return &OauthAccessTokenMappingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetMappings - Get the list of Access Token Mappings.
//RequestType: GET
//Input:
func (s *OauthAccessTokenMappingsService) GetMappings() (result *models.AccessTokenMappings, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings"
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

//CreateMapping - Create a new Access Token Mapping.
//RequestType: POST
//Input: input *CreateMappingInput
func (s *OauthAccessTokenMappingsService) CreateMapping(input *CreateMappingInput) (result *models.AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetMapping - Find the Access Token Mapping by its ID.
//RequestType: GET
//Input: input *GetMappingInput
func (s *OauthAccessTokenMappingsService) GetMapping(input *GetMappingInput) (result *models.AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
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

//UpdateMapping - Update an Access Token Mapping.
//RequestType: PUT
//Input: input *UpdateMappingInput
func (s *OauthAccessTokenMappingsService) UpdateMapping(input *UpdateMappingInput) (result *models.AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteMapping - Delete an Access Token Mapping.
//RequestType: DELETE
//Input: input *DeleteMappingInput
func (s *OauthAccessTokenMappingsService) DeleteMapping(input *DeleteMappingInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
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

type CreateMappingInput struct {
	Body models.AccessTokenMapping

	BypassExternalValidation *bool
}

type DeleteMappingInput struct {
	Id string
}

type GetMappingInput struct {
	Id string
}

type UpdateMappingInput struct {
	Body models.AccessTokenMapping
	Id   string

	BypassExternalValidation *bool
}
