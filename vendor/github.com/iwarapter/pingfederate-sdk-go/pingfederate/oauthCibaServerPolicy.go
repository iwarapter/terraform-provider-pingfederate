package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthCibaServerPolicyService service

//GetSettings - Get general ciba server request policy settings.
//RequestType: GET
//Input:
func (s *OauthCibaServerPolicyService) GetSettings() (result *CibaServerPolicySettings, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/settings"
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

//UpdateSettings - Update general ciba server request policy settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthCibaServerPolicyService) UpdateSettings(input *UpdateSettingsInput) (result *CibaServerPolicySettings, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/settings"
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

//GetPolicies - Get list of request policies.
//RequestType: GET
//Input:
func (s *OauthCibaServerPolicyService) GetPolicies() (result *RequestPolicies, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies"
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

//CreatePolicy - Create a new request policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthCibaServerPolicyService) CreatePolicy(input *CreatePolicyInput) (result *RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies"
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

//GetPolicy - Find request policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthCibaServerPolicyService) GetPolicy(input *GetPolicyInput) (result *RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
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

//UpdatePolicy - Update a request policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthCibaServerPolicyService) UpdatePolicy(input *UpdatePolicyInput) (result *RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
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

//DeletePolicy - Delete a request policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthCibaServerPolicyService) DeletePolicy(input *DeletePolicyInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
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
