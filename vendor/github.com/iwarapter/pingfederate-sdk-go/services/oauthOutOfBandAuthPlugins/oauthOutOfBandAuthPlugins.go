package oauthOutOfBandAuthPlugins

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthOutOfBandAuthPluginsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthOutOfBandAuthPluginsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthOutOfBandAuthPluginsService {

	return &OauthOutOfBandAuthPluginsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetOOBAuthPluginDescriptors - Get the list of available Out of Band authenticator plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthPluginDescriptors() (result *models.OutOfBandAuthPluginDescriptors, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/descriptors"
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

//GetOOBAuthPluginDescriptor - Get the descriptor of an Out of Band authenticator plugin.
//RequestType: GET
//Input: input *GetOOBAuthPluginDescriptorInput
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthPluginDescriptor(input *GetOOBAuthPluginDescriptorInput) (result *models.OutOfBandAuthPluginDescriptor, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/descriptors/{id}"
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

//GetOOBAuthenticators - Get a list of Out of Band authenticator plugin instances.
//RequestType: GET
//Input:
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthenticators() (result *models.OutOfBandAuthenticators, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins"
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

//CreateOOBAuthenticator - Create an Out of Band authenticator plugin instance.
//RequestType: POST
//Input: input *CreateOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) CreateOOBAuthenticator(input *CreateOOBAuthenticatorInput) (result *models.OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins"
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

//GetOOBAuthenticator - Get a specific Out of Band authenticator plugin instance.
//RequestType: GET
//Input: input *GetOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthenticator(input *GetOOBAuthenticatorInput) (result *models.OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
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

//UpdateOOBAuthenticator - Update an Out of Band authenticator plugin instance.
//RequestType: PUT
//Input: input *UpdateOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) UpdateOOBAuthenticator(input *UpdateOOBAuthenticatorInput) (result *models.OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
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

//DeleteOOBAuthenticator - Delete an Out of Band authenticator plugin instance.
//RequestType: DELETE
//Input: input *DeleteOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) DeleteOOBAuthenticator(input *DeleteOOBAuthenticatorInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
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

//GetActions - List of actions for an Out of Band authenticator plugin instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *OauthOutOfBandAuthPluginsService) GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions"
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

//GetAction - Find an Out of Band authenticator plugin instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *OauthOutOfBandAuthPluginsService) GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions/{actionId}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

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

//InvokeAction - Invokes an action for Out of Band authenticator plugin instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *OauthOutOfBandAuthPluginsService) InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions/{actionId}/invokeAction"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type CreateOOBAuthenticatorInput struct {
	Body models.OutOfBandAuthenticator
}

type DeleteOOBAuthenticatorInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetOOBAuthPluginDescriptorInput struct {
	Id string
}

type GetOOBAuthenticatorInput struct {
	Id string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateOOBAuthenticatorInput struct {
	Body models.OutOfBandAuthenticator
	Id   string
}
