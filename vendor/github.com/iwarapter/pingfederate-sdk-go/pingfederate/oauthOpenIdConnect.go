package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthOpenIdConnectService service

//GetSettings - Get the OpenID Connect Settings.
//RequestType: GET
//Input:
func (s *OauthOpenIdConnectService) GetSettings() (result *OpenIdConnectSettings, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/settings"
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

//UpdateSettings - Set the OpenID Connect Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthOpenIdConnectService) UpdateSettings(input *UpdateSettingsInput) (result *OpenIdConnectSettings, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/settings"
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

//GetPolicies - Get list of OpenID Connect Policies.
//RequestType: GET
//Input:
func (s *OauthOpenIdConnectService) GetPolicies() (result *OpenIdConnectPolicies, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies"
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

//CreatePolicy - Create a new OpenID Connect Policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthOpenIdConnectService) CreatePolicy(input *CreatePolicyInput) (result *OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
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

//GetPolicy - Find OpenID Connect Policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthOpenIdConnectService) GetPolicy(input *GetPolicyInput) (result *OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
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

//UpdatePolicy - Update an OpenID Connect Policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthOpenIdConnectService) UpdatePolicy(input *UpdatePolicyInput) (result *OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//DeletePolicy - Delete an OpenID Connect Policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthOpenIdConnectService) DeletePolicy(input *DeletePolicyInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
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
