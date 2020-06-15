package idpTokenProcessors

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpTokenProcessorsService struct {
	Client *client.PfClient
}

// New creates a new instance of the IdpTokenProcessorsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *IdpTokenProcessorsService {

	return &IdpTokenProcessorsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetTokenProcessorDescriptors - Get the list of available token processors.
//RequestType: GET
//Input:
func (s *IdpTokenProcessorsService) GetTokenProcessorDescriptors() (result *models.TokenProcessorDescriptors, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/descriptors"
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

//GetTokenProcessorDescriptorsById - Get the description of a token processor plugin by ID.
//RequestType: GET
//Input: input *GetTokenProcessorDescriptorsByIdInput
func (s *IdpTokenProcessorsService) GetTokenProcessorDescriptorsById(input *GetTokenProcessorDescriptorsByIdInput) (result *models.TokenProcessorDescriptor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/descriptors/{id}"
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

//GetTokenProcessors - Get the list of token processor instances.
//RequestType: GET
//Input:
func (s *IdpTokenProcessorsService) GetTokenProcessors() (result *models.TokenProcessors, resp *http.Response, err error) {
	path := "/idp/tokenProcessors"
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

//CreateTokenProcessor - Create a new token processor instance.
//RequestType: POST
//Input: input *CreateTokenProcessorInput
func (s *IdpTokenProcessorsService) CreateTokenProcessor(input *CreateTokenProcessorInput) (result *models.TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetTokenProcessor - Find a token processor instance by ID.
//RequestType: GET
//Input: input *GetTokenProcessorInput
func (s *IdpTokenProcessorsService) GetTokenProcessor(input *GetTokenProcessorInput) (result *models.TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
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

//UpdateTokenProcessor - Update a token processor instance.
//RequestType: PUT
//Input: input *UpdateTokenProcessorInput
func (s *IdpTokenProcessorsService) UpdateTokenProcessor(input *UpdateTokenProcessorInput) (result *models.TokenProcessor, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteTokenProcessor - Delete a token processor instance.
//RequestType: DELETE
//Input: input *DeleteTokenProcessorInput
func (s *IdpTokenProcessorsService) DeleteTokenProcessor(input *DeleteTokenProcessorInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/idp/tokenProcessors/{id}"
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

type CreateTokenProcessorInput struct {
	Body models.TokenProcessor
}

type DeleteTokenProcessorInput struct {
	Id string
}

type GetTokenProcessorInput struct {
	Id string
}

type GetTokenProcessorDescriptorsByIdInput struct {
	Id string
}

type UpdateTokenProcessorInput struct {
	Body models.TokenProcessor
	Id   string
}
