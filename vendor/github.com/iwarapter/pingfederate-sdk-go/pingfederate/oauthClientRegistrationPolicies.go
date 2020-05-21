package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthClientRegistrationPoliciesService service

//GetDynamicClientRegistrationDescriptors - Get the list of available client registration policy plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationDescriptors() (result *ClientRegistrationPolicyDescriptors, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/descriptors"
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

//GetDynamicClientRegistrationDescriptor - Get the description of a client registration policy plugin descriptor.
//RequestType: GET
//Input: input *GetDynamicClientRegistrationDescriptorInput
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationDescriptor(input *GetDynamicClientRegistrationDescriptorInput) (result *ClientRegistrationPolicyDescriptor, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/descriptors/{id}"
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

//GetDynamicClientRegistrationPolicies - Get a list of client registration policy plugin instances.
//RequestType: GET
//Input:
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationPolicies() (result *ClientRegistrationPolicies, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies"
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

//CreateDynamicClientRegistrationPolicy - Create a client registration policy plugin instance.
//RequestType: POST
//Input: input *CreateDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) CreateDynamicClientRegistrationPolicy(input *CreateDynamicClientRegistrationPolicyInput) (result *ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetDynamicClientRegistrationPolicy - Get a specific client registration policy plugin instance.
//RequestType: GET
//Input: input *GetDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationPolicy(input *GetDynamicClientRegistrationPolicyInput) (result *ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
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

//UpdateDynamicClientRegistrationPolicy - Update a client registration policy plugin instance.
//RequestType: PUT
//Input: input *UpdateDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) UpdateDynamicClientRegistrationPolicy(input *UpdateDynamicClientRegistrationPolicyInput) (result *ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
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

//DeleteDynamicClientRegistrationPolicy - Delete a client registration policy plugin instance.
//RequestType: DELETE
//Input: input *DeleteDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) DeleteDynamicClientRegistrationPolicy(input *DeleteDynamicClientRegistrationPolicyInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
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
