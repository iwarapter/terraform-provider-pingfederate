package oauthClientRegistrationPolicies

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientRegistrationPoliciesService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthClientRegistrationPoliciesService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthClientRegistrationPoliciesService {

	return &OauthClientRegistrationPoliciesService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetDynamicClientRegistrationDescriptors - Get the list of available client registration policy plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationDescriptors() (result *models.ClientRegistrationPolicyDescriptors, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/descriptors"
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

//GetDynamicClientRegistrationDescriptor - Get the description of a client registration policy plugin descriptor.
//RequestType: GET
//Input: input *GetDynamicClientRegistrationDescriptorInput
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationDescriptor(input *GetDynamicClientRegistrationDescriptorInput) (result *models.ClientRegistrationPolicyDescriptor, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/descriptors/{id}"
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

//GetDynamicClientRegistrationPolicies - Get a list of client registration policy plugin instances.
//RequestType: GET
//Input:
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationPolicies() (result *models.ClientRegistrationPolicies, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies"
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

//CreateDynamicClientRegistrationPolicy - Create a client registration policy plugin instance.
//RequestType: POST
//Input: input *CreateDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) CreateDynamicClientRegistrationPolicy(input *CreateDynamicClientRegistrationPolicyInput) (result *models.ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies"
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

//GetDynamicClientRegistrationPolicy - Get a specific client registration policy plugin instance.
//RequestType: GET
//Input: input *GetDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) GetDynamicClientRegistrationPolicy(input *GetDynamicClientRegistrationPolicyInput) (result *models.ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
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

//UpdateDynamicClientRegistrationPolicy - Update a client registration policy plugin instance.
//RequestType: PUT
//Input: input *UpdateDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) UpdateDynamicClientRegistrationPolicy(input *UpdateDynamicClientRegistrationPolicyInput) (result *models.ClientRegistrationPolicy, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
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

//DeleteDynamicClientRegistrationPolicy - Delete a client registration policy plugin instance.
//RequestType: DELETE
//Input: input *DeleteDynamicClientRegistrationPolicyInput
func (s *OauthClientRegistrationPoliciesService) DeleteDynamicClientRegistrationPolicy(input *DeleteDynamicClientRegistrationPolicyInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/clientRegistrationPolicies/{id}"
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

type CreateDynamicClientRegistrationPolicyInput struct {
	Body models.ClientRegistrationPolicy
}

type DeleteDynamicClientRegistrationPolicyInput struct {
	Id string
}

type GetDynamicClientRegistrationDescriptorInput struct {
	Id string
}

type GetDynamicClientRegistrationPolicyInput struct {
	Id string
}

type UpdateDynamicClientRegistrationPolicyInput struct {
	Body models.ClientRegistrationPolicy
	Id   string
}
