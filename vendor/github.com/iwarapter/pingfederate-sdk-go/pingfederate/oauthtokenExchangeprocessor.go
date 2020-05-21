package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthTokenExchangeProcessorService service

//GetSettings - Get general OAuth 2.0 Token Exchange Processor settings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeProcessorService) GetSettings() (result *TokenExchangeProcessorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/settings"
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

//UpdateSettings - Update general OAuth 2.0 Token Exchange Processor settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthTokenExchangeProcessorService) UpdateSettings(input *UpdateSettingsInput) (result *TokenExchangeProcessorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPolicies - Get list of OAuth 2.0 Token Exchange Processor policies.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeProcessorService) GetPolicies() (result *TokenExchangeProcessorPolicies, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies"
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

//CreatePolicy - Create a new OAuth 2.0 Token Exchange Processor policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthTokenExchangeProcessorService) CreatePolicy(input *CreatePolicyInput) (result *TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPolicy - Find an OAuth 2.0 Token Exchange Processor policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthTokenExchangeProcessorService) GetPolicy(input *GetPolicyInput) (result *TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
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

//UpdatePolicy - Update an OAuth 2.0 Token Exchange Processor policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthTokenExchangeProcessorService) UpdatePolicy(input *UpdatePolicyInput) (result *TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeletePolicy - Delete an OAuth 2.0 Token Exchange Processor policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthTokenExchangeProcessorService) DeletePolicy(input *DeletePolicyInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
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
