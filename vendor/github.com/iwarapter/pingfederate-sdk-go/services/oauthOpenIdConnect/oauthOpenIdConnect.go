package oauthOpenIdConnect

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthOpenIdConnectService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthOpenIdConnectService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthOpenIdConnectService {

	return &OauthOpenIdConnectService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSettings - Get the OpenID Connect Settings.
//RequestType: GET
//Input:
func (s *OauthOpenIdConnectService) GetSettings() (result *models.OpenIdConnectSettings, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateSettings - Set the OpenID Connect Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthOpenIdConnectService) UpdateSettings(input *UpdateSettingsInput) (result *models.OpenIdConnectSettings, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPolicies - Get list of OpenID Connect Policies.
//RequestType: GET
//Input:
func (s *OauthOpenIdConnectService) GetPolicies() (result *models.OpenIdConnectPolicies, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//CreatePolicy - Create a new OpenID Connect Policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthOpenIdConnectService) CreatePolicy(input *CreatePolicyInput) (result *models.OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPolicy - Find OpenID Connect Policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthOpenIdConnectService) GetPolicy(input *GetPolicyInput) (result *models.OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdatePolicy - Update an OpenID Connect Policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthOpenIdConnectService) UpdatePolicy(input *UpdatePolicyInput) (result *models.OpenIdConnectPolicy, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeletePolicy - Delete an OpenID Connect Policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthOpenIdConnectService) DeletePolicy(input *DeletePolicyInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/openIdConnect/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type CreatePolicyInput struct {
	Body models.OpenIdConnectPolicy

	BypassExternalValidation *bool
}

type DeletePolicyInput struct {
	Id string
}

type GetPolicyInput struct {
	Id string
}

type UpdatePolicyInput struct {
	Body models.OpenIdConnectPolicy
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.OpenIdConnectSettings
}
