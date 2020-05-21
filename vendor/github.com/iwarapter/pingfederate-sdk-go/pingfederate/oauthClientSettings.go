package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type OauthClientSettingsService service

//GetClientSettings - Configure the client settings.
//RequestType: GET
//Input:
func (s *OauthClientSettingsService) GetClientSettings() (result *ClientSettings, resp *http.Response, err error) {
	path := "/oauth/clientSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateClientSettings - Update the client settings.
//RequestType: PUT
//Input: input *UpdateClientSettingsInput
func (s *OauthClientSettingsService) UpdateClientSettings(input *UpdateClientSettingsInput) (result *ClientSettings, resp *http.Response, err error) {
	path := "/oauth/clientSettings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
