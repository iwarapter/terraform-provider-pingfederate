package session

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SessionService struct {
	Client *client.PfClient
}

// New creates a new instance of the SessionService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *SessionService {

	return &SessionService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSessionSettings - Get general session management settings.
//RequestType: GET
//Input:
func (s *SessionService) GetSessionSettings() (result *models.SessionSettings, resp *http.Response, err error) {
	path := "/session/settings"
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

//UpdateSessionSettings - Update general session management settings.
//RequestType: PUT
//Input: input *UpdateSessionSettingsInput
func (s *SessionService) UpdateSessionSettings(input *UpdateSessionSettingsInput) (result *models.SessionSettings, resp *http.Response, err error) {
	path := "/session/settings"
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

//GetGlobalPolicy - Get the global authentication session policy.
//RequestType: GET
//Input:
func (s *SessionService) GetGlobalPolicy() (result *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/global"
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

//UpdateGlobalPolicy - Update the global authentication session policy.
//RequestType: PUT
//Input: input *UpdateGlobalPolicyInput
func (s *SessionService) UpdateGlobalPolicy(input *UpdateGlobalPolicyInput) (result *models.GlobalAuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/global"
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

//GetApplicationPolicy - Get the application session policy.
//RequestType: GET
//Input:
func (s *SessionService) GetApplicationPolicy() (result *models.ApplicationSessionPolicy, resp *http.Response, err error) {
	path := "/session/applicationSessionPolicy"
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

//UpdateApplicationPolicy - Update the application session policy.
//RequestType: PUT
//Input: input *UpdateApplicationPolicyInput
func (s *SessionService) UpdateApplicationPolicy(input *UpdateApplicationPolicyInput) (result *models.ApplicationSessionPolicy, resp *http.Response, err error) {
	path := "/session/applicationSessionPolicy"
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

//GetSourcePolicies - Get list of session policies.
//RequestType: GET
//Input:
func (s *SessionService) GetSourcePolicies() (result *models.AuthenticationSessionPolicies, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies"
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

//CreateSourcePolicy - Create a new session policy.
//RequestType: POST
//Input: input *CreateSourcePolicyInput
func (s *SessionService) CreateSourcePolicy(input *CreateSourcePolicyInput) (result *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetSourcePolicy - Find session policy by ID.
//RequestType: GET
//Input: input *GetSourcePolicyInput
func (s *SessionService) GetSourcePolicy(input *GetSourcePolicyInput) (result *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
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

//UpdateSourcePolicy - Update a session policy.
//RequestType: PUT
//Input: input *UpdateSourcePolicyInput
func (s *SessionService) UpdateSourcePolicy(input *UpdateSourcePolicyInput) (result *models.AuthenticationSessionPolicy, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//DeleteSourcePolicy - Delete a session policy.
//RequestType: DELETE
//Input: input *DeleteSourcePolicyInput
func (s *SessionService) DeleteSourcePolicy(input *DeleteSourcePolicyInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/session/authenticationSessionPolicies/{id}"
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

type CreateSourcePolicyInput struct {
	Body models.AuthenticationSessionPolicy
}

type DeleteSourcePolicyInput struct {
	Id string
}

type GetSourcePolicyInput struct {
	Id string
}

type UpdateApplicationPolicyInput struct {
	Body models.ApplicationSessionPolicy
}

type UpdateGlobalPolicyInput struct {
	Body models.GlobalAuthenticationSessionPolicy
}

type UpdateSessionSettingsInput struct {
	Body models.SessionSettings
}

type UpdateSourcePolicyInput struct {
	Body models.AuthenticationSessionPolicy
	Id   string
}
