package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SessionService service

//GetSessionSettings - Get general session management settings.
//RequestType: GET
//Input:
func (s *SessionService) GetSessionSettings() (result *SessionSettings, resp *http.Response, err error) {
	path := "/session/settings"
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

//UpdateSessionSettings - Update general session management settings.
//RequestType: PUT
//Input: input *UpdateSessionSettingsInput
func (s *SessionService) UpdateSessionSettings(input *UpdateSessionSettingsInput) (result *SessionSettings, resp *http.Response, err error) {
	path := "/session/settings"
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

//GetGlobalPolicy - Get the global authentication session policy.
//RequestType: GET
//Input:
func (s *SessionService) GetGlobalPolicy() (result *GlobalAuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/global"
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

//UpdateGlobalPolicy - Update the global authentication session policy.
//RequestType: PUT
//Input: input *UpdateGlobalPolicyInput
func (s *SessionService) UpdateGlobalPolicy(input *UpdateGlobalPolicyInput) (result *GlobalAuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/global"
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

//GetApplicationPolicy - Get the application session policy.
//RequestType: GET
//Input:
func (s *SessionService) GetApplicationPolicy() (result *ApplicationSessionPolicy, resp *http.Response, err error) {
	path := "/session/applicationSessionPolicy"
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

//UpdateApplicationPolicy - Update the application session policy.
//RequestType: PUT
//Input: input *UpdateApplicationPolicyInput
func (s *SessionService) UpdateApplicationPolicy(input *UpdateApplicationPolicyInput) (result *ApplicationSessionPolicy, resp *http.Response, err error) {
	path := "/session/applicationSessionPolicy"
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

//GetSourcePolicies - Get list of session policies.
//RequestType: GET
//Input:
func (s *SessionService) GetSourcePolicies() (result *AuthenticationSessionPolicies, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies"
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

//CreateSourcePolicy - Create a new session policy.
//RequestType: POST
//Input: input *CreateSourcePolicyInput
func (s *SessionService) CreateSourcePolicy(input *CreateSourcePolicyInput) (result *AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies"
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

//GetSourcePolicy - Find session policy by ID.
//RequestType: GET
//Input: input *GetSourcePolicyInput
func (s *SessionService) GetSourcePolicy(input *GetSourcePolicyInput) (result *AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
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

//UpdateSourcePolicy - Update a session policy.
//RequestType: PUT
//Input: input *UpdateSourcePolicyInput
func (s *SessionService) UpdateSourcePolicy(input *UpdateSourcePolicyInput) (result *AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
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

//DeleteSourcePolicy - Delete a session policy.
//RequestType: DELETE
//Input: input *DeleteSourcePolicyInput
func (s *SessionService) DeleteSourcePolicy(input *DeleteSourcePolicyInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
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
