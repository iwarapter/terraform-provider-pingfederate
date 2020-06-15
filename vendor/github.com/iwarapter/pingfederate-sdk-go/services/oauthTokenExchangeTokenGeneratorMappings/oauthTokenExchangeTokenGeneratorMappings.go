package oauthTokenExchangeTokenGeneratorMappings

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeTokenGeneratorMappingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthTokenExchangeTokenGeneratorMappingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthTokenExchangeTokenGeneratorMappingsService {

	return &OauthTokenExchangeTokenGeneratorMappingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetTokenGeneratorMappings - Get the list of Token Exchange Processor policy to Token Generator Mappings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeTokenGeneratorMappingsService) GetTokenGeneratorMappings() (result *models.ProcessorPolicyToGeneratorMappings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//CreateTokenGeneratorMapping - Create a new Token Exchange Processor policy to Token Generator Mapping.
//RequestType: POST
//Input: input *CreateTokenGeneratorMappingInput
func (s *OauthTokenExchangeTokenGeneratorMappingsService) CreateTokenGeneratorMapping(input *CreateTokenGeneratorMappingInput) (result *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetTokenGeneratorMappingById - Get a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: GET
//Input: input *GetTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeTokenGeneratorMappingsService) GetTokenGeneratorMappingById(input *GetTokenGeneratorMappingByIdInput) (result *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateTokenGeneratorMappingById - Update a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: PUT
//Input: input *UpdateTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeTokenGeneratorMappingsService) UpdateTokenGeneratorMappingById(input *UpdateTokenGeneratorMappingByIdInput) (result *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteTokenGeneratorMappingById - Delete a Token Exchange Processor policy to Token Generator Mapping.
//RequestType: DELETE
//Input: input *DeleteTokenGeneratorMappingByIdInput
func (s *OauthTokenExchangeTokenGeneratorMappingsService) DeleteTokenGeneratorMappingById(input *DeleteTokenGeneratorMappingByIdInput) (result *models.ProcessorPolicyToGeneratorMapping, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type CreateTokenGeneratorMappingInput struct {
	Body models.ProcessorPolicyToGeneratorMapping

	BypassExternalValidation *bool
}

type DeleteTokenGeneratorMappingByIdInput struct {
	Id string
}

type GetTokenGeneratorMappingByIdInput struct {
	Id string
}

type UpdateTokenGeneratorMappingByIdInput struct {
	Body models.ProcessorPolicyToGeneratorMapping
	Id   string

	BypassExternalValidation *bool
}
