package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthTokenExchangeService service

//GetTokenGeneratorMappings - Get the list of Token Exchange Processor policy to Token Generator Mappings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeService) GetTokenGeneratorMappings() (result *ProcessorPolicyToGeneratorMappings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings"
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

//CreateTokenGeneratorMapping - Create a new Token Exchange Processor policy to Token Generator Mapping.
//RequestType: POST
//Input: input *CreateTokenGeneratorMappingInput
func (s *OauthTokenExchangeService) CreateTokenGeneratorMapping(input *CreateTokenGeneratorMappingInput) (result *ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings"
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

//GetTokenGeneratorMappingById - Get a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: GET
//Input: input *GetTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeService) GetTokenGeneratorMappingById(input *GetTokenGeneratorMappingByIdInput) (result *ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
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

//UpdateTokenGeneratorMappingById - Update a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: PUT
//Input: input *UpdateTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeService) UpdateTokenGeneratorMappingById(input *UpdateTokenGeneratorMappingByIdInput) (result *ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
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

//DeleteTokenGeneratorMappingById - Delete a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: DELETE
//Input: input *DeleteTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeService) DeleteTokenGeneratorMappingById(input *DeleteTokenGeneratorMappingByIdInput) (result *ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
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
