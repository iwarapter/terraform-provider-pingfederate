package idpToSpAdapterMapping

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpToSpAdapterMappingService struct {
	Client *client.PfClient
}

// New creates a new instance of the IdpToSpAdapterMappingService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *IdpToSpAdapterMappingService {

	return &IdpToSpAdapterMappingService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetIdpToSpAdapterMappings - Get list of IdP-to-SP Adapter Mappings.
//RequestType: GET
//Input:
func (s *IdpToSpAdapterMappingService) GetIdpToSpAdapterMappings() (result *models.IdpToSpAdapterMappings, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping"
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

//CreateIdpToSpAdapterMapping - Create a new IdP-to-SP Adapter mapping.
//RequestType: POST
//Input: input *CreateIdpToSpAdapterMappingInput
func (s *IdpToSpAdapterMappingService) CreateIdpToSpAdapterMapping(input *CreateIdpToSpAdapterMappingInput) (result *models.IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping"
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

//GetIdpToSpAdapterMappingsById - Get an IdP-to-SP Adapter Mapping.
//RequestType: GET
//Input: input *GetIdpToSpAdapterMappingsByIdInput
func (s *IdpToSpAdapterMappingService) GetIdpToSpAdapterMappingsById(input *GetIdpToSpAdapterMappingsByIdInput) (result *models.IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
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

//UpdateIdpToSpAdapterMapping - Update the specified IdP-to-SP Adapter mapping.
//RequestType: PUT
//Input: input *UpdateIdpToSpAdapterMappingInput
func (s *IdpToSpAdapterMappingService) UpdateIdpToSpAdapterMapping(input *UpdateIdpToSpAdapterMappingInput) (result *models.IdpToSpAdapterMapping, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
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

//DeleteIdpToSpAdapterMappingsById - Delete an Adapter to Adapter Mapping.
//RequestType: DELETE
//Input: input *DeleteIdpToSpAdapterMappingsByIdInput
func (s *IdpToSpAdapterMappingService) DeleteIdpToSpAdapterMappingsById(input *DeleteIdpToSpAdapterMappingsByIdInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/idpToSpAdapterMapping/{id}"
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

type CreateIdpToSpAdapterMappingInput struct {
	Body models.IdpToSpAdapterMapping

	BypassExternalValidation *bool
}

type DeleteIdpToSpAdapterMappingsByIdInput struct {
	Id string
}

type GetIdpToSpAdapterMappingsByIdInput struct {
	Id string
}

type UpdateIdpToSpAdapterMappingInput struct {
	Body models.IdpToSpAdapterMapping
	Id   string

	BypassExternalValidation *bool
}
