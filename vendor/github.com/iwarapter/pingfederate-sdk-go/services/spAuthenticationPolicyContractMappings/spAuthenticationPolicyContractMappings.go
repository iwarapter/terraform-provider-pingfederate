package spAuthenticationPolicyContractMappings

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpAuthenticationPolicyContractMappingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the SpAuthenticationPolicyContractMappingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *SpAuthenticationPolicyContractMappingsService {

	return &SpAuthenticationPolicyContractMappingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetApcToSpAdapterMappings - Get the list of APC-to-SP Adapter Mappings.
//RequestType: GET
//Input:
func (s *SpAuthenticationPolicyContractMappingsService) GetApcToSpAdapterMappings() (result *models.ApcToSpAdapterMappings, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings"
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

//CreateApcToSpAdapterMapping - Create a new APC-to-SP Adapter Mapping.
//RequestType: POST
//Input: input *CreateApcToSpAdapterMappingInput
func (s *SpAuthenticationPolicyContractMappingsService) CreateApcToSpAdapterMapping(input *CreateApcToSpAdapterMappingInput) (result *models.ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings"
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

//GetApcToSpAdapterMappingById - Get an APC-to-SP Adapter Mapping.
//RequestType: GET
//Input: input *GetApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) GetApcToSpAdapterMappingById(input *GetApcToSpAdapterMappingByIdInput) (result *models.ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
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

//UpdateApcToSpAdapterMappingById - Update an APC-to-SP Adapter Mapping.
//RequestType: PUT
//Input: input *UpdateApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) UpdateApcToSpAdapterMappingById(input *UpdateApcToSpAdapterMappingByIdInput) (result *models.ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
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

//DeleteApcToSpAdapterMappingById - Delete an APC-to-SP Adapter Mapping.
//RequestType: DELETE
//Input: input *DeleteApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) DeleteApcToSpAdapterMappingById(input *DeleteApcToSpAdapterMappingByIdInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
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

type CreateApcToSpAdapterMappingInput struct {
	Body models.ApcToSpAdapterMapping

	BypassExternalValidation *bool
}

type DeleteApcToSpAdapterMappingByIdInput struct {
	Id string
}

type GetApcToSpAdapterMappingByIdInput struct {
	Id string
}

type UpdateApcToSpAdapterMappingByIdInput struct {
	Body models.ApcToSpAdapterMapping
	Id   string

	BypassExternalValidation *bool
}
