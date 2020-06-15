package redirectValidation

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type RedirectValidationService struct {
	Client *client.PfClient
}

// New creates a new instance of the RedirectValidationService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *RedirectValidationService {

	return &RedirectValidationService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetRedirectValidationSettings - Retrieve redirect validation settings.
//RequestType: GET
//Input:
func (s *RedirectValidationService) GetRedirectValidationSettings() (result *models.RedirectValidationSettings, resp *http.Response, err error) {
	path := "/redirectValidation"
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

//UpdateRedirectValidationSettings - Update redirect validation settings.
//RequestType: PUT
//Input: input *UpdateRedirectValidationSettingsInput
func (s *RedirectValidationService) UpdateRedirectValidationSettings(input *UpdateRedirectValidationSettingsInput) (result *models.RedirectValidationSettings, resp *http.Response, err error) {
	path := "/redirectValidation"
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

type UpdateRedirectValidationSettingsInput struct {
	Body models.RedirectValidationSettings
}
