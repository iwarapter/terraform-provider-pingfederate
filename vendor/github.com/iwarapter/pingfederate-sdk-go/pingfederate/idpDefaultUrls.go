package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type IdpDefaultUrlsService service

//GetDefaultUrl - Gets the IDP Default URL settings.
//RequestType: GET
//Input:
func (s *IdpDefaultUrlsService) GetDefaultUrl() (result *IdpDefaultUrl, resp *http.Response, err error) {
	path := "/idp/defaultUrls"
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

//UpdateDefaultUrlSettings - Update the IDP Default URL settings.
//RequestType: PUT
//Input: input *UpdateDefaultUrlSettingsInput
func (s *IdpDefaultUrlsService) UpdateDefaultUrlSettings(input *UpdateDefaultUrlSettingsInput) (result *IdpDefaultUrl, resp *http.Response, err error) {
	path := "/idp/defaultUrls"
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
