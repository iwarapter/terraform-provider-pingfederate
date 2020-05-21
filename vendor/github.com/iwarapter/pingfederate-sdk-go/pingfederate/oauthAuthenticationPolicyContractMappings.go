package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthAuthenticationPolicyContractMappingsService service

//GetApcMappings - Get the list of authentication policy contract to persistent grant mappings.
//RequestType: GET
//Input:
func (s *OauthAuthenticationPolicyContractMappingsService) GetApcMappings() (result *ApcToPersistentGrantMappings, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings"
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

//CreateApcMapping - Create a new authentication policy contract to persistent grant mapping.
//RequestType: POST
//Input: input *CreateApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) CreateApcMapping(input *CreateApcMappingInput) (result *ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings"
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

//GetApcMapping - Find the authentication policy contract to persistent grant mapping by ID.
//RequestType: GET
//Input: input *GetApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) GetApcMapping(input *GetApcMappingInput) (result *ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
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

//UpdateApcMapping - Update an authentication policy contract to persistent grant mapping.
//RequestType: PUT
//Input: input *UpdateApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) UpdateApcMapping(input *UpdateApcMappingInput) (result *ApcToPersistentGrantMapping, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
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

//DeleteApcMapping - Delete an authentication policy contract to persistent grant mapping.
//RequestType: DELETE
//Input: input *DeleteApcMappingInput
func (s *OauthAuthenticationPolicyContractMappingsService) DeleteApcMapping(input *DeleteApcMappingInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/authenticationPolicyContractMappings/{id}"
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
