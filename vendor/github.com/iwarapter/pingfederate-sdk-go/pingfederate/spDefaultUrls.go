package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type SpDefaultUrlsService service

//GetDefaultUrls - Gets the SP Default URLs. These are Values that affect the user's experience when executing SP-initiated SSO operations.
//RequestType: GET
//Input:
func (s *SpDefaultUrlsService) GetDefaultUrls() (result *SpDefaultUrls, resp *http.Response, err error) {
	path := "/sp/defaultUrls"
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

//UpdateDefaultUrls - Update the SP Default URLs. Enter values that affect the user's experience when executing SP-initiated SSO operations.
//RequestType: PUT
//Input: input *UpdateDefaultUrlsInput
func (s *SpDefaultUrlsService) UpdateDefaultUrls(input *UpdateDefaultUrlsInput) (result *SpDefaultUrls, resp *http.Response, err error) {
	path := "/sp/defaultUrls"
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
