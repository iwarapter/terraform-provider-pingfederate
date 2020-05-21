package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type TokenProcessorToTokenGeneratorMappingsService service

//GetTokenToTokenMappings - Get the list of Token Processor to Token Generator Mappings.
//RequestType: GET
//Input:
func (s *TokenProcessorToTokenGeneratorMappingsService) GetTokenToTokenMappings() (result *TokenToTokenMappings, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings"
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

//CreateTokenToTokenMapping - Create a new Token Processor to Token Generator Mapping.
//RequestType: POST
//Input: input *CreateTokenToTokenMappingInput
func (s *TokenProcessorToTokenGeneratorMappingsService) CreateTokenToTokenMapping(input *CreateTokenToTokenMappingInput) (result *TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings"
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

//GetTokenToTokenMappingById - Get a Token Processor to Token Generator Mapping.
//RequestType: GET
//Input: input *GetTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) GetTokenToTokenMappingById(input *GetTokenToTokenMappingByIdInput) (result *TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
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

//UpdateTokenToTokenMappingById - Update a Token Processor to Token Generator Mapping.
//RequestType: PUT
//Input: input *UpdateTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) UpdateTokenToTokenMappingById(input *UpdateTokenToTokenMappingByIdInput) (result *TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
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

//DeleteTokenToTokenMappingById - Delete a Token Processor to Token Generator Mapping.
//RequestType: DELETE
//Input: input *DeleteTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) DeleteTokenToTokenMappingById(input *DeleteTokenToTokenMappingByIdInput) (result *TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
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
