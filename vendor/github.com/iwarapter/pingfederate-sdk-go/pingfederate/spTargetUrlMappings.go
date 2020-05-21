package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type SpTargetUrlMappingsService service

//GetUrlMappings - List the mappings between URLs and adapter or connection instances.
//RequestType: GET
//Input:
func (s *SpTargetUrlMappingsService) GetUrlMappings() (result *SpUrlMappings, resp *http.Response, err error) {
	path := "/sp/targetUrlMappings"
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

//UpdateUrlMappings - Update the mappings between URLs and adapters or connections instances.
//RequestType: PUT
//Input: input *UpdateUrlMappingsInput
func (s *SpTargetUrlMappingsService) UpdateUrlMappings(input *UpdateUrlMappingsInput) (result *SpUrlMappings, resp *http.Response, err error) {
	path := "/sp/targetUrlMappings"
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
