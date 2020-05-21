package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthIdpAdapterMappingsService service

//GetIdpAdapterMappings - Get the list of IdP adapter mappings.
//RequestType: GET
//Input:
func (s *OauthIdpAdapterMappingsService) GetIdpAdapterMappings() (result *IdpAdapterMappings, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings"
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

//CreateIdpAdapterMapping - Create a new IdP adapter mapping.
//RequestType: POST
//Input: input *CreateIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) CreateIdpAdapterMapping(input *CreateIdpAdapterMappingInput) (result *IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings"
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

//GetIdpAdapterMapping - Find the IdP adapter mapping by the ID.
//RequestType: GET
//Input: input *GetIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) GetIdpAdapterMapping(input *GetIdpAdapterMappingInput) (result *IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
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

//UpdateIdpAdapterMapping - Update an IdP adapter mapping.
//RequestType: PUT
//Input: input *UpdateIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) UpdateIdpAdapterMapping(input *UpdateIdpAdapterMappingInput) (result *IdpAdapterMapping, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
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

//DeleteIdpAdapterMapping - Delete an IdP adapter mapping.
//RequestType: DELETE
//Input: input *DeleteIdpAdapterMappingInput
func (s *OauthIdpAdapterMappingsService) DeleteIdpAdapterMapping(input *DeleteIdpAdapterMappingInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/idpAdapterMappings/{id}"
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
