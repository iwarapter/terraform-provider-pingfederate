package oauthCibaServerPolicy

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthCibaServerPolicyService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthCibaServerPolicyService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthCibaServerPolicyService {

	return &OauthCibaServerPolicyService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSettings - Get general ciba server request policy settings.
//RequestType: GET
//Input:
func (s *OauthCibaServerPolicyService) GetSettings() (result *models.CibaServerPolicySettings, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/settings"
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

//UpdateSettings - Update general ciba server request policy settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *OauthCibaServerPolicyService) UpdateSettings(input *UpdateSettingsInput) (result *models.CibaServerPolicySettings, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/settings"
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

//GetPolicies - Get list of request policies.
//RequestType: GET
//Input:
func (s *OauthCibaServerPolicyService) GetPolicies() (result *models.RequestPolicies, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies"
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

//CreatePolicy - Create a new request policy.
//RequestType: POST
//Input: input *CreatePolicyInput
func (s *OauthCibaServerPolicyService) CreatePolicy(input *CreatePolicyInput) (result *models.RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies"
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

//GetPolicy - Find request policy by ID.
//RequestType: GET
//Input: input *GetPolicyInput
func (s *OauthCibaServerPolicyService) GetPolicy(input *GetPolicyInput) (result *models.RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
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

//UpdatePolicy - Update a request policy.
//RequestType: PUT
//Input: input *UpdatePolicyInput
func (s *OauthCibaServerPolicyService) UpdatePolicy(input *UpdatePolicyInput) (result *models.RequestPolicy, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
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

//DeletePolicy - Delete a request policy.
//RequestType: DELETE
//Input: input *DeletePolicyInput
func (s *OauthCibaServerPolicyService) DeletePolicy(input *DeletePolicyInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/cibaServerPolicy/requestPolicies/{id}"
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
	Body models.RequestPolicy

	BypassExternalValidation *bool
}

type DeletePolicyInput struct {
	Id string
}

type GetPolicyInput struct {
	Id string
}

type UpdatePolicyInput struct {
	Body models.RequestPolicy
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.CibaServerPolicySettings

	BypassExternalValidation *bool
}
