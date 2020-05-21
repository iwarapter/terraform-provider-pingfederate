package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type RedirectValidationService service

//GetRedirectValidationSettings - Retrieve redirect validation settings.
//RequestType: GET
//Input:
func (s *RedirectValidationService) GetRedirectValidationSettings() (result *RedirectValidationSettings, resp *http.Response, err error) {
	path := "/redirectValidation"
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

//UpdateRedirectValidationSettings - Update redirect validation settings.
//RequestType: PUT
//Input: input *UpdateRedirectValidationSettingsInput
func (s *RedirectValidationService) UpdateRedirectValidationSettings(input *UpdateRedirectValidationSettingsInput) (result *RedirectValidationSettings, resp *http.Response, err error) {
	path := "/redirectValidation"
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
