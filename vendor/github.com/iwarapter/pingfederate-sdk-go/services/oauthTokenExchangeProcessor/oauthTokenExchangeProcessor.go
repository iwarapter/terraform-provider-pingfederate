package oauthTokenExchangeProcessor

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthTokenExchangeProcessorService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthTokenExchangeProcessorService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthTokenExchangeProcessorService {

	return &OauthTokenExchangeProcessorService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSettings - Get general OAuth 2.0 Token Exchange Processor settings.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeProcessorService) GetSettings() (result *models.TokenExchangeProcessorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/settings"
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

//UpdateSettings - Update general OAuth 2.0 Token Exchange Processor settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthTokenExchangeProcessorService) UpdateSettings(input *UpdateSettingsInput) (result *models.TokenExchangeProcessorSettings, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPolicies - Get list of OAuth 2.0 Token Exchange Processor policies.
//RequestType: GET
//Input:
func (s *OauthTokenExchangeProcessorService) GetPolicies() (result *models.TokenExchangeProcessorPolicies, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies"
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

//CreatePolicy - Create a new OAuth 2.0 Token Exchange Processor policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthTokenExchangeProcessorService) CreatePolicy(input *CreatePolicyInput) (result *models.TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPolicy - Find an OAuth 2.0 Token Exchange Processor policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthTokenExchangeProcessorService) GetPolicy(input *GetPolicyInput) (result *models.TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
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

//UpdatePolicy - Update an OAuth 2.0 Token Exchange Processor policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthTokenExchangeProcessorService) UpdatePolicy(input *UpdatePolicyInput) (result *models.TokenExchangeProcessorPolicy, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("bypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeletePolicy - Delete an OAuth 2.0 Token Exchange Processor policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthTokenExchangeProcessorService) DeletePolicy(input *DeletePolicyInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/tokenExchange/processor/policies/{id}"
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
	Body models.TokenExchangeProcessorPolicy

	BypassExternalValidation *bool
}

type DeletePolicyInput struct {
	Id string
}

type GetPolicyInput struct {
	Id string
}

type UpdatePolicyInput struct {
	Body models.TokenExchangeProcessorPolicy
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.TokenExchangeProcessorSettings

	BypassExternalValidation *bool
}
