package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type IdpTokenProcessorsService service

//GetTokenProcessorDescriptors - Get the list of available token processors.
//RequestType: GET
//Input:
func (s *IdpTokenProcessorsService) GetTokenProcessorDescriptors() (result *TokenProcessorDescriptors, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/descriptors"
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

//GetTokenProcessorDescriptorsById - Get the description of a token processor plugin by ID.
//RequestType: GET
//Input: input *GetTokenProcessorDescriptorsByIdInput
func (s *IdpTokenProcessorsService) GetTokenProcessorDescriptorsById(input *GetTokenProcessorDescriptorsByIdInput) (result *TokenProcessorDescriptor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/descriptors/{id}"
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

//GetTokenProcessors - Get the list of token processor instances.
//RequestType: GET
//Input:
func (s *IdpTokenProcessorsService) GetTokenProcessors() (result *TokenProcessors, resp *http.Response, err error) {
	path := "/idp/tokenProcessors"
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

//CreateTokenProcessor - Create a new token processor instance.
//RequestType: POST
//Input: input *CreateTokenProcessorInput
func (s *IdpTokenProcessorsService) CreateTokenProcessor(input *CreateTokenProcessorInput) (result *TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetTokenProcessor - Find a token processor instance by ID.
//RequestType: GET
//Input: input *GetTokenProcessorInput
func (s *IdpTokenProcessorsService) GetTokenProcessor(input *GetTokenProcessorInput) (result *TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
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

//UpdateTokenProcessor - Update a token processor instance.
//RequestType: PUT
//Input: input *UpdateTokenProcessorInput
func (s *IdpTokenProcessorsService) UpdateTokenProcessor(input *UpdateTokenProcessorInput) (result *TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteTokenProcessor - Delete a token processor instance.
//RequestType: DELETE
//Input: input *DeleteTokenProcessorInput
func (s *IdpTokenProcessorsService) DeleteTokenProcessor(input *DeleteTokenProcessorInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
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
