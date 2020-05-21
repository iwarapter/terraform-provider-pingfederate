package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type ExtendedPropertiesService service

//GetExtendedProperties - Get the defined Extended Properties.
//RequestType: GET
//Input:
func (s *ExtendedPropertiesService) GetExtendedProperties() (result *ExtendedProperties, resp *http.Response, err error) {
	path := "/extendedProperties"
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

//UpdateExtendedProperties - Update the Extended Properties.
//RequestType: PUT
//Input: input *UpdateExtendedPropertiesInput
func (s *ExtendedPropertiesService) UpdateExtendedProperties(input *UpdateExtendedPropertiesInput) (result *ExtendedProperties, resp *http.Response, err error) {
	path := "/extendedProperties"
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
