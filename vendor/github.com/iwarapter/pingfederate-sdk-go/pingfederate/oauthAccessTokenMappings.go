package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthAccessTokenMappingsService service

//GetMappings - Get the list of Access Token Mappings.
//RequestType: GET
//Input:
func (s *OauthAccessTokenMappingsService) GetMappings() (result *AccessTokenMappings, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings"
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

//CreateMapping - Create a new Access Token Mapping.
//RequestType: POST
//Input: input *CreateMappingInput
func (s *OauthAccessTokenMappingsService) CreateMapping(input *CreateMappingInput) (result *AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings"
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

//GetMapping - Find the Access Token Mapping by its ID.
//RequestType: GET
//Input: input *GetMappingInput
func (s *OauthAccessTokenMappingsService) GetMapping(input *GetMappingInput) (result *AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
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

//UpdateMapping - Update an Access Token Mapping.
//RequestType: PUT
//Input: input *UpdateMappingInput
func (s *OauthAccessTokenMappingsService) UpdateMapping(input *UpdateMappingInput) (result *AccessTokenMapping, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
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

//DeleteMapping - Delete an Access Token Mapping.
//RequestType: DELETE
//Input: input *DeleteMappingInput
func (s *OauthAccessTokenMappingsService) DeleteMapping(input *DeleteMappingInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/accessTokenMappings/{id}"
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
