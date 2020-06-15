package oauthIdpAdapterMappings

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthIdpAdapterMappingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthIdpAdapterMappingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthIdpAdapterMappingsService {

	return &OauthIdpAdapterMappingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetIdpAdapterMappings - Get the list of IdP adapter mappings.
//RequestType: GET
//Input:
func (s *OauthIdpAdapterMappingsService) GetIdpAdapterMappings() (result *models.IdpAdapterMappings, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings"
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

//CreateIdpAdapterMapping - Create a new IdP adapter mapping.
//RequestType: POST
//Input: input *CreateIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) CreateIdpAdapterMapping(input *CreateIdpAdapterMappingInput) (result *models.IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings"
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

//GetIdpAdapterMapping - Find the IdP adapter mapping by the ID.
//RequestType: GET
//Input: input *GetIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) GetIdpAdapterMapping(input *GetIdpAdapterMappingInput) (result *models.IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
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

//UpdateIdpAdapterMapping - Update an IdP adapter mapping.
//RequestType: PUT
//Input: input *UpdateIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) UpdateIdpAdapterMapping(input *UpdateIdpAdapterMappingInput) (result *models.IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
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

//DeleteIdpAdapterMapping - Delete an IdP adapter mapping.
//RequestType: DELETE
//Input: input *DeleteIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) DeleteIdpAdapterMapping(input *DeleteIdpAdapterMappingInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
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

type CreateIdpAdapterMappingInput struct {
	Body models.IdpAdapterMapping

	BypassExternalValidation *bool
}

type DeleteIdpAdapterMappingInput struct {
	Id string
}

type GetIdpAdapterMappingInput struct {
	Id string
}

type UpdateIdpAdapterMappingInput struct {
	Body models.IdpAdapterMapping
	Id   string

	BypassExternalValidation *bool
}
