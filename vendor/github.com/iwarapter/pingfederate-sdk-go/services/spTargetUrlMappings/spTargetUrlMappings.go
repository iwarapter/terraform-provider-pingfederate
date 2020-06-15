package spTargetUrlMappings

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpTargetUrlMappingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the SpTargetUrlMappingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *SpTargetUrlMappingsService {

	return &SpTargetUrlMappingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetUrlMappings - List the mappings between URLs and adapter or connection instances.
//RequestType: GET
//Input:
func (s *SpTargetUrlMappingsService) GetUrlMappings() (result *models.SpUrlMappings, resp *http.Response, err error) {
	path := "/sp/targetUrlMappings"
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

//UpdateUrlMappings - Update the mappings between URLs and adapters or connections instances.
//RequestType: PUT
//Input: input *UpdateUrlMappingsInput
func (s *SpTargetUrlMappingsService) UpdateUrlMappings(input *UpdateUrlMappingsInput) (result *models.SpUrlMappings, resp *http.Response, err error) {
	path := "/sp/targetUrlMappings"
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

type UpdateUrlMappingsInput struct {
	Body models.SpUrlMappings
}
