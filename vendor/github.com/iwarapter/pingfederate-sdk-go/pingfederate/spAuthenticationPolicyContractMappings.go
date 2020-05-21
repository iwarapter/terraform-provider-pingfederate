package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SpAuthenticationPolicyContractMappingsService service

//GetApcToSpAdapterMappings - Get the list of APC-to-SP Adapter Mappings.
//RequestType: GET
//Input:
func (s *SpAuthenticationPolicyContractMappingsService) GetApcToSpAdapterMappings() (result *ApcToSpAdapterMappings, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings"
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

//CreateApcToSpAdapterMapping - Create a new APC-to-SP Adapter Mapping.
//RequestType: POST
//Input: input *CreateApcToSpAdapterMappingInput
func (s *SpAuthenticationPolicyContractMappingsService) CreateApcToSpAdapterMapping(input *CreateApcToSpAdapterMappingInput) (result *ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings"
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

//GetApcToSpAdapterMappingById - Get an APC-to-SP Adapter Mapping.
//RequestType: GET
//Input: input *GetApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) GetApcToSpAdapterMappingById(input *GetApcToSpAdapterMappingByIdInput) (result *ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
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

//UpdateApcToSpAdapterMappingById - Update an APC-to-SP Adapter Mapping.
//RequestType: PUT
//Input: input *UpdateApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) UpdateApcToSpAdapterMappingById(input *UpdateApcToSpAdapterMappingByIdInput) (result *ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
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

//DeleteApcToSpAdapterMappingById - Delete an APC-to-SP Adapter Mapping.
//RequestType: DELETE
//Input: input *DeleteApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) DeleteApcToSpAdapterMappingById(input *DeleteApcToSpAdapterMappingByIdInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
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
