package authenticationPolicies

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationPoliciesService struct {
	Client *client.PfClient
}

// New creates a new instance of the AuthenticationPoliciesService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *AuthenticationPoliciesService {

	return &AuthenticationPoliciesService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSettings - Get the authentication policies settings.
//RequestType: GET
//Input:
func (s *AuthenticationPoliciesService) GetSettings() (result *models.AuthenticationPoliciesSettings, resp *http.Response, err error) {
	path := "/authenticationPolicies/settings"
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

//UpdateSettings - Set the authentication policies settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *AuthenticationPoliciesService) UpdateSettings(input *UpdateSettingsInput) (result *models.AuthenticationPoliciesSettings, resp *http.Response, err error) {
	path := "/authenticationPolicies/settings"
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

//GetDefaultAuthenticationPolicy - Get the default configured authentication policy.
//RequestType: GET
//Input:
func (s *AuthenticationPoliciesService) GetDefaultAuthenticationPolicy() (result *models.AuthenticationPolicy, resp *http.Response, err error) {
	path := "/authenticationPolicies/default"
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

//UpdateDefaultAuthenticationPolicy - Set the default authentication policy.
//RequestType: PUT
//Input: input *UpdateDefaultAuthenticationPolicyInput
func (s *AuthenticationPoliciesService) UpdateDefaultAuthenticationPolicy(input *UpdateDefaultAuthenticationPolicyInput) (result *models.AuthenticationPolicy, resp *http.Response, err error) {
	path := "/authenticationPolicies/default"
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

type UpdateDefaultAuthenticationPolicyInput struct {
	Body models.AuthenticationPolicy

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.AuthenticationPoliciesSettings
}
