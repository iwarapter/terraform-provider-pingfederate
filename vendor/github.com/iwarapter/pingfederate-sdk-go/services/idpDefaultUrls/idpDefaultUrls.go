package idpDefaultUrls

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpDefaultUrlsService struct {
	Client *client.PfClient
}

// New creates a new instance of the IdpDefaultUrlsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *IdpDefaultUrlsService {

	return &IdpDefaultUrlsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetDefaultUrl - Gets the IDP Default URL settings.
//RequestType: GET
//Input:
func (s *IdpDefaultUrlsService) GetDefaultUrl() (result *models.IdpDefaultUrl, resp *http.Response, err error) {
	path := "/idp/defaultUrls"
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

//UpdateDefaultUrlSettings - Update the IDP Default URL settings.
//RequestType: PUT
//Input: input *UpdateDefaultUrlSettingsInput
func (s *IdpDefaultUrlsService) UpdateDefaultUrlSettings(input *UpdateDefaultUrlSettingsInput) (result *models.IdpDefaultUrl, resp *http.Response, err error) {
	path := "/idp/defaultUrls"
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

type UpdateDefaultUrlSettingsInput struct {
	Body models.IdpDefaultUrl
}
