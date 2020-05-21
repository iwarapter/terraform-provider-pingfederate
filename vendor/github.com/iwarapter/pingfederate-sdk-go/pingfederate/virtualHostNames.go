package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type VirtualHostNamesService service

//GetVirtualHostNamesSettings - Retrieve virtual host names settings.
//RequestType: GET
//Input:
func (s *VirtualHostNamesService) GetVirtualHostNamesSettings() (result *VirtualHostNameSettings, resp *http.Response, err error) {
	path := "/virtualHostNames"
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

//UpdateVirtualHostNamesSettings - Update virtual host names settings.
//RequestType: PUT
//Input: input *UpdateVirtualHostNamesSettingsInput
func (s *VirtualHostNamesService) UpdateVirtualHostNamesSettings(input *UpdateVirtualHostNamesSettingsInput) (result *VirtualHostNameSettings, resp *http.Response, err error) {
	path := "/virtualHostNames"
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
