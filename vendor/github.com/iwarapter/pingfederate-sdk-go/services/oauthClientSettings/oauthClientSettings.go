package oauthClientSettings

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientSettingsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthClientSettingsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthClientSettingsService {

	return &OauthClientSettingsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetClientSettings - Configure the client settings.
//RequestType: GET
//Input:
func (s *OauthClientSettingsService) GetClientSettings() (result *models.ClientSettings, resp *http.Response, err error) {
	path := "/oauth/clientSettings"
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

//UpdateClientSettings - Update the client settings.
//RequestType: PUT
//Input: input *UpdateClientSettingsInput
func (s *OauthClientSettingsService) UpdateClientSettings(input *UpdateClientSettingsInput) (result *models.ClientSettings, resp *http.Response, err error) {
	path := "/oauth/clientSettings"
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

type UpdateClientSettingsInput struct {
	Body models.ClientSettings
}
