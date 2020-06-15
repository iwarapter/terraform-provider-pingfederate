package extendedProperties

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ExtendedPropertiesService struct {
	Client *client.PfClient
}

// New creates a new instance of the ExtendedPropertiesService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *ExtendedPropertiesService {

	return &ExtendedPropertiesService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetExtendedProperties - Get the defined Extended Properties.
//RequestType: GET
//Input:
func (s *ExtendedPropertiesService) GetExtendedProperties() (result *models.ExtendedProperties, resp *http.Response, err error) {
	path := "/extendedProperties"
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

//UpdateExtendedProperties - Update the Extended Properties.
//RequestType: PUT
//Input: input *UpdateExtendedPropertiesInput
func (s *ExtendedPropertiesService) UpdateExtendedProperties(input *UpdateExtendedPropertiesInput) (result *models.ExtendedProperties, resp *http.Response, err error) {
	path := "/extendedProperties"
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

type UpdateExtendedPropertiesInput struct {
	Body models.ExtendedProperties
}
