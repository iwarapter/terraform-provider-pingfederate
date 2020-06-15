package oauthResourceOwnerCredentialsMappings

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthResourceOwnerCredentialsMappingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthResourceOwnerCredentialsMappingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthResourceOwnerCredentialsMappingsService {

	return &OauthResourceOwnerCredentialsMappingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetResourceOwnerCredentialsMappings - Get the list of Resource Owner Credentials mappings.
//RequestType: GET
//Input:
func (s *OauthResourceOwnerCredentialsMappingsService) GetResourceOwnerCredentialsMappings() (result *models.ResourceOwnerCredentialsMappings, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings"
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

//CreateResourceOwnerCredentialsMapping - Create a new Resource Owner Credentials mapping.
//RequestType: POST
//Input: input *CreateResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) CreateResourceOwnerCredentialsMapping(input *CreateResourceOwnerCredentialsMappingInput) (result *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings"
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

//GetResourceOwnerCredentialsMapping - Find the Resource Owner Credentials mapping by the ID.
//RequestType: GET
//Input: input *GetResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) GetResourceOwnerCredentialsMapping(input *GetResourceOwnerCredentialsMappingInput) (result *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
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

//UpdateResourceOwnerCredentialsMapping - Update a Resource Owner Credentials mapping.
//RequestType: PUT
//Input: input *UpdateResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) UpdateResourceOwnerCredentialsMapping(input *UpdateResourceOwnerCredentialsMappingInput) (result *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
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

//DeleteResourceOwnerCredentialsMapping - Delete a Resource Owner Credentials mapping.
//RequestType: DELETE
//Input: input *DeleteResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) DeleteResourceOwnerCredentialsMapping(input *DeleteResourceOwnerCredentialsMappingInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
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

type CreateResourceOwnerCredentialsMappingInput struct {
	Body models.ResourceOwnerCredentialsMapping

	BypassExternalValidation *bool
}

type DeleteResourceOwnerCredentialsMappingInput struct {
	Id string
}

type GetResourceOwnerCredentialsMappingInput struct {
	Id string
}

type UpdateResourceOwnerCredentialsMappingInput struct {
	Body models.ResourceOwnerCredentialsMapping
	Id   string

	BypassExternalValidation *bool
}
