package tokenProcessorToTokenGeneratorMappings

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type TokenProcessorToTokenGeneratorMappingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the TokenProcessorToTokenGeneratorMappingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *TokenProcessorToTokenGeneratorMappingsService {

	return &TokenProcessorToTokenGeneratorMappingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetTokenToTokenMappings - Get the list of Token Processor to Token Generator Mappings.
//RequestType: GET
//Input:
func (s *TokenProcessorToTokenGeneratorMappingsService) GetTokenToTokenMappings() (result *models.TokenToTokenMappings, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings"
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

//CreateTokenToTokenMapping - Create a new Token Processor to Token Generator Mapping.
//RequestType: POST
//Input: input *CreateTokenToTokenMappingInput
func (s *TokenProcessorToTokenGeneratorMappingsService) CreateTokenToTokenMapping(input *CreateTokenToTokenMappingInput) (result *models.TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings"
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

//GetTokenToTokenMappingById - Get a Token Processor to Token Generator Mapping.
//RequestType: GET
//Input: input *GetTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) GetTokenToTokenMappingById(input *GetTokenToTokenMappingByIdInput) (result *models.TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
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

//UpdateTokenToTokenMappingById - Update a Token Processor to Token Generator Mapping.
//RequestType: PUT
//Input: input *UpdateTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) UpdateTokenToTokenMappingById(input *UpdateTokenToTokenMappingByIdInput) (result *models.TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
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

//DeleteTokenToTokenMappingById - Delete a Token Processor to Token Generator Mapping.
//RequestType: DELETE
//Input: input *DeleteTokenToTokenMappingByIdInput
func (s *TokenProcessorToTokenGeneratorMappingsService) DeleteTokenToTokenMappingById(input *DeleteTokenToTokenMappingByIdInput) (result *models.TokenToTokenMapping, resp *http.Response, err error) {
	path := "/tokenProcessorToTokenGeneratorMappings/{id}"
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

type CreateTokenToTokenMappingInput struct {
	Body models.TokenToTokenMapping

	BypassExternalValidation *bool
}

type DeleteTokenToTokenMappingByIdInput struct {
	Id string
}

type GetTokenToTokenMappingByIdInput struct {
	Id string
}

type UpdateTokenToTokenMappingByIdInput struct {
	Body models.TokenToTokenMapping
	Id   string

	BypassExternalValidation *bool
}
