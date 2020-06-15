package spTokenGenerators

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpTokenGeneratorsService struct {
	Client *client.PfClient
}

// New creates a new instance of the SpTokenGeneratorsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *SpTokenGeneratorsService {

	return &SpTokenGeneratorsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetTokenGeneratorDescriptors - Get the list of available token generators.
//RequestType: GET
//Input:
func (s *SpTokenGeneratorsService) GetTokenGeneratorDescriptors() (result *models.TokenGeneratorDescriptors, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/descriptors"
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

//GetTokenGeneratorDescriptorsById - Get the description of a token generator plugin by ID.
//RequestType: GET
//Input: input *GetTokenGeneratorDescriptorsByIdInput
func (s *SpTokenGeneratorsService) GetTokenGeneratorDescriptorsById(input *GetTokenGeneratorDescriptorsByIdInput) (result *models.TokenGeneratorDescriptor, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/descriptors/{id}"
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

//GetTokenGenerators - Get the list of token generator instances.
//RequestType: GET
//Input:
func (s *SpTokenGeneratorsService) GetTokenGenerators() (result *models.TokenGenerators, resp *http.Response, err error) {
	path := "/sp/tokenGenerators"
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

//CreateTokenGenerator - Create a new token generator instance.
//RequestType: POST
//Input: input *CreateTokenGeneratorInput
func (s *SpTokenGeneratorsService) CreateTokenGenerator(input *CreateTokenGeneratorInput) (result *models.TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators"
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

//GetTokenGenerator - Find a token generator instance by ID.
//RequestType: GET
//Input: input *GetTokenGeneratorInput
func (s *SpTokenGeneratorsService) GetTokenGenerator(input *GetTokenGeneratorInput) (result *models.TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
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

//UpdateTokenGenerator - Update a token generator instance.
//RequestType: PUT
//Input: input *UpdateTokenGeneratorInput
func (s *SpTokenGeneratorsService) UpdateTokenGenerator(input *UpdateTokenGeneratorInput) (result *models.TokenGenerator, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
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

//DeleteTokenGenerator - Delete a token generator instance.
//RequestType: DELETE
//Input: input *DeleteTokenGeneratorInput
func (s *SpTokenGeneratorsService) DeleteTokenGenerator(input *DeleteTokenGeneratorInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/sp/tokenGenerators/{id}"
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

type CreateTokenGeneratorInput struct {
	Body models.TokenGenerator
}

type DeleteTokenGeneratorInput struct {
	Id string
}

type GetTokenGeneratorInput struct {
	Id string
}

type GetTokenGeneratorDescriptorsByIdInput struct {
	Id string
}

type UpdateTokenGeneratorInput struct {
	Body models.TokenGenerator
	Id   string
}
