package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type IdpConnectorsService service

//GetIdpConnectorDescriptors - Get the list of available IdP connector descriptors.
//RequestType: GET
//Input:
func (s *IdpConnectorsService) GetIdpConnectorDescriptors() (result *SaasPluginDescriptors, resp *http.Response, err error) {
	path := "/idp/connectors/descriptors"
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

//GetIdpConnectorDescriptorById - Get the list of available connector descriptors.
//RequestType: GET
//Input: input *GetIdpConnectorDescriptorByIdInput
func (s *IdpConnectorsService) GetIdpConnectorDescriptorById(input *GetIdpConnectorDescriptorByIdInput) (result *SaasPluginDescriptor, resp *http.Response, err error) {
	path := "/idp/connectors/descriptors/{id}"
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
