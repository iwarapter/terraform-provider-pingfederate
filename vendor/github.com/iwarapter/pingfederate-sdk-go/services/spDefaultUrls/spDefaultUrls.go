package spDefaultUrls

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpDefaultUrlsService struct {
	Client *client.PfClient
}

// New creates a new instance of the SpDefaultUrlsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *SpDefaultUrlsService {

	return &SpDefaultUrlsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetDefaultUrls - Gets the SP Default URLs. These are Values that affect the user's experience when executing SP-initiated SSO operations.
//RequestType: GET
//Input:
func (s *SpDefaultUrlsService) GetDefaultUrls() (result *models.SpDefaultUrls, resp *http.Response, err error) {
	path := "/sp/defaultUrls"
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

//UpdateDefaultUrls - Update the SP Default URLs. Enter values that affect the user's experience when executing SP-initiated SSO operations.
//RequestType: PUT
//Input: input *UpdateDefaultUrlsInput
func (s *SpDefaultUrlsService) UpdateDefaultUrls(input *UpdateDefaultUrlsInput) (result *models.SpDefaultUrls, resp *http.Response, err error) {
	path := "/sp/defaultUrls"
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

type UpdateDefaultUrlsInput struct {
	Body models.SpDefaultUrls
}
