package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type IdpToSpAdapterMappingService service

//GetIdpToSpAdapterMappings - Get list of IdP-to-SP Adapter Mappings.
//RequestType: GET
//Input:
func (s *IdpToSpAdapterMappingService) GetIdpToSpAdapterMappings() (result *IdpToSpAdapterMappings, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping"
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

//CreateIdpToSpAdapterMapping - Create a new IdP-to-SP Adapter mapping.
//RequestType: POST
//Input: input *CreateIdpToSpAdapterMappingInput
func (s *IdpToSpAdapterMappingService) CreateIdpToSpAdapterMapping(input *CreateIdpToSpAdapterMappingInput) (result *IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping"
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

//GetIdpToSpAdapterMappingsById - Get an IdP-to-SP Adapter Mapping.
//RequestType: GET
//Input: input *GetIdpToSpAdapterMappingsByIdInput
func (s *IdpToSpAdapterMappingService) GetIdpToSpAdapterMappingsById(input *GetIdpToSpAdapterMappingsByIdInput) (result *IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
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

//UpdateIdpToSpAdapterMapping - Update the specified IdP-to-SP Adapter mapping.
//RequestType: PUT
//Input: input *UpdateIdpToSpAdapterMappingInput
func (s *IdpToSpAdapterMappingService) UpdateIdpToSpAdapterMapping(input *UpdateIdpToSpAdapterMappingInput) (result *IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
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

//DeleteIdpToSpAdapterMappingsById - Delete an Adapter to Adapter Mapping.
//RequestType: DELETE
//Input: input *DeleteIdpToSpAdapterMappingsByIdInput
func (s *IdpToSpAdapterMappingService) DeleteIdpToSpAdapterMappingsById(input *DeleteIdpToSpAdapterMappingsByIdInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
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
