package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type AuthenticationPoliciesService service

//GetSettings - Get the authentication policies settings.
//RequestType: GET
//Input:
func (s *AuthenticationPoliciesService) GetSettings() (result *AuthenticationPoliciesSettings, resp *http.Response, err error) {
	path := "/authenticationPolicies/settings"
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

//UpdateSettings - Set the authentication policies settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *AuthenticationPoliciesService) UpdateSettings(input *UpdateSettingsInput) (result *AuthenticationPoliciesSettings, resp *http.Response, err error) {
	path := "/authenticationPolicies/settings"
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

//GetDefaultAuthenticationPolicy - Get the default configured authentication policy.
//RequestType: GET
//Input:
func (s *AuthenticationPoliciesService) GetDefaultAuthenticationPolicy() (result *AuthenticationPolicy, resp *http.Response, err error) {
	path := "/authenticationPolicies/default"
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

//UpdateDefaultAuthenticationPolicy - Set the default authentication policy.
//RequestType: PUT
//Input: input *UpdateDefaultAuthenticationPolicyInput
func (s *AuthenticationPoliciesService) UpdateDefaultAuthenticationPolicy(input *UpdateDefaultAuthenticationPolicyInput) (result *AuthenticationPolicy, resp *http.Response, err error) {
	path := "/authenticationPolicies/default"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
