package configStore

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConfigStoreService struct {
	Client *client.PfClient
}

// New creates a new instance of the ConfigStoreService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *ConfigStoreService {

	return &ConfigStoreService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSetting - Get a single setting from a bundle.
//RequestType: GET
//Input: input *GetSettingInput
func (s *ConfigStoreService) GetSetting(input *GetSettingInput) (result *models.ConfigStoreSetting, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

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

//UpdateSetting - Create or update a setting/bundle.
//RequestType: PUT
//Input: input *UpdateSettingInput
func (s *ConfigStoreService) UpdateSetting(input *UpdateSettingInput) (result *models.ConfigStoreSetting, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

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

//DeleteSetting - Delete a setting.
//RequestType: DELETE
//Input: input *DeleteSettingInput
func (s *ConfigStoreService) DeleteSetting(input *DeleteSettingInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

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

//GetSettings - Get all settings from a bundle.
//RequestType: GET
//Input: input *GetSettingsInput
func (s *ConfigStoreService) GetSettings(input *GetSettingsInput) (result *models.ConfigStoreBundle, resp *http.Response, err error) {
	path := "/configStore/{bundle}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

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

type DeleteSettingInput struct {
	Bundle string
	Id     string
}

type GetSettingInput struct {
	Bundle string
	Id     string
}

type GetSettingsInput struct {
	Bundle string
}

type UpdateSettingInput struct {
	Body   models.ConfigStoreSetting
	Bundle string
	Id     string
}
