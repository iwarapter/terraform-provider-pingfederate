package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ConfigStoreService service

//GetSetting - Get a single setting from a bundle.
//RequestType: GET
//Input: input *GetSettingInput
func (s *ConfigStoreService) GetSetting(input *GetSettingInput) (result *ConfigStoreSetting, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateSetting - Create or update a setting/bundle.
//RequestType: PUT
//Input: input *UpdateSettingInput
func (s *ConfigStoreService) UpdateSetting(input *UpdateSettingInput) (result *ConfigStoreSetting, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

	path = strings.Replace(path, "{id}", input.Id, -1)

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

//DeleteSetting - Delete a setting.
//RequestType: DELETE
//Input: input *DeleteSettingInput
func (s *ConfigStoreService) DeleteSetting(input *DeleteSettingInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetSettings - Get all settings from a bundle.
//RequestType: GET
//Input: input *GetSettingsInput
func (s *ConfigStoreService) GetSettings(input *GetSettingsInput) (result *ConfigStoreBundle, resp *http.Response, err error) {
	path := "/configStore/{bundle}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

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
