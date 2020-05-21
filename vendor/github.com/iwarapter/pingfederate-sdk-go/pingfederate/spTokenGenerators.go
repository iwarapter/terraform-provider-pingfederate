package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SpTokenGeneratorsService service

//GetTokenGeneratorDescriptors - Get the list of available token generators.
//RequestType: GET
//Input:
func (s *SpTokenGeneratorsService) GetTokenGeneratorDescriptors() (result *TokenGeneratorDescriptors, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/descriptors"
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

//GetTokenGeneratorDescriptorsById - Get the description of a token generator plugin by ID.
//RequestType: GET
//Input: input *GetTokenGeneratorDescriptorsByIdInput
func (s *SpTokenGeneratorsService) GetTokenGeneratorDescriptorsById(input *GetTokenGeneratorDescriptorsByIdInput) (result *TokenGeneratorDescriptor, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/descriptors/{id}"
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

//GetTokenGenerators - Get the list of token generator instances.
//RequestType: GET
//Input:
func (s *SpTokenGeneratorsService) GetTokenGenerators() (result *TokenGenerators, resp *http.Response, err error) {
	path := "/sp/tokenGenerators"
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

//CreateTokenGenerator - Create a new token generator instance.
//RequestType: POST
//Input: input *CreateTokenGeneratorInput
func (s *SpTokenGeneratorsService) CreateTokenGenerator(input *CreateTokenGeneratorInput) (result *TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators"
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

//GetTokenGenerator - Find a token generator instance by ID.
//RequestType: GET
//Input: input *GetTokenGeneratorInput
func (s *SpTokenGeneratorsService) GetTokenGenerator(input *GetTokenGeneratorInput) (result *TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
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

//UpdateTokenGenerator - Update a token generator instance.
//RequestType: PUT
//Input: input *UpdateTokenGeneratorInput
func (s *SpTokenGeneratorsService) UpdateTokenGenerator(input *UpdateTokenGeneratorInput) (result *TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
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

//DeleteTokenGenerator - Delete a token generator instance.
//RequestType: DELETE
//Input: input *DeleteTokenGeneratorInput
func (s *SpTokenGeneratorsService) DeleteTokenGenerator(input *DeleteTokenGeneratorInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
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
