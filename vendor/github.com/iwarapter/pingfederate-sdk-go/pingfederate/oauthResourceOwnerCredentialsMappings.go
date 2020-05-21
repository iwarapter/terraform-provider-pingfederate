package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthResourceOwnerCredentialsMappingsService service

//GetResourceOwnerCredentialsMappings - Get the list of Resource Owner Credentials mappings.
//RequestType: GET
//Input:
func (s *OauthResourceOwnerCredentialsMappingsService) GetResourceOwnerCredentialsMappings() (result *ResourceOwnerCredentialsMappings, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings"
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

//CreateResourceOwnerCredentialsMapping - Create a new Resource Owner Credentials mapping.
//RequestType: POST
//Input: input *CreateResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) CreateResourceOwnerCredentialsMapping(input *CreateResourceOwnerCredentialsMappingInput) (result *ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetResourceOwnerCredentialsMapping - Find the Resource Owner Credentials mapping by the ID.
//RequestType: GET
//Input: input *GetResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) GetResourceOwnerCredentialsMapping(input *GetResourceOwnerCredentialsMappingInput) (result *ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
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

//UpdateResourceOwnerCredentialsMapping - Update a Resource Owner Credentials mapping.
//RequestType: PUT
//Input: input *UpdateResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) UpdateResourceOwnerCredentialsMapping(input *UpdateResourceOwnerCredentialsMappingInput) (result *ResourceOwnerCredentialsMapping, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteResourceOwnerCredentialsMapping - Delete a Resource Owner Credentials mapping.
//RequestType: DELETE
//Input: input *DeleteResourceOwnerCredentialsMappingInput
func (s *OauthResourceOwnerCredentialsMappingsService) DeleteResourceOwnerCredentialsMapping(input *DeleteResourceOwnerCredentialsMappingInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/resourceOwnerCredentialsMappings/{id}"
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
