package oauthAuthenticationPolicyContractMappings

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAuthenticationPolicyContractMappingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthAuthenticationPolicyContractMappingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthAuthenticationPolicyContractMappingsService {

	return &OauthAuthenticationPolicyContractMappingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetApcMappings - Get the list of authentication policy contract to persistent grant mappings.
//RequestType: GET
//Input:
func (s *OauthAuthenticationPolicyContractMappingsService) GetApcMappings() (result *models.ApcToPersistentGrantMappings, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings"
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

//CreateApcMapping - Create a new authentication policy contract to persistent grant mapping.
//RequestType: POST
//Input: input *CreateApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) CreateApcMapping(input *CreateApcMappingInput) (result *models.ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings"
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

//GetApcMapping - Find the authentication policy contract to persistent grant mapping by ID.
//RequestType: GET
//Input: input *GetApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) GetApcMapping(input *GetApcMappingInput) (result *models.ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
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

//UpdateApcMapping - Update an authentication policy contract to persistent grant mapping.
//RequestType: PUT
//Input: input *UpdateApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) UpdateApcMapping(input *UpdateApcMappingInput) (result *models.ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
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

//DeleteApcMapping - Delete an authentication policy contract to persistent grant mapping.
//RequestType: DELETE
//Input: input *DeleteApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) DeleteApcMapping(input *DeleteApcMappingInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
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

type CreateApcMappingInput struct {
	Body models.ApcToPersistentGrantMapping

	BypassExternalValidation *bool
}

type DeleteApcMappingInput struct {
	Id string
}

type GetApcMappingInput struct {
	Id string
}

type UpdateApcMappingInput struct {
	Body models.ApcToPersistentGrantMapping
	Id   string

	BypassExternalValidation *bool
}
