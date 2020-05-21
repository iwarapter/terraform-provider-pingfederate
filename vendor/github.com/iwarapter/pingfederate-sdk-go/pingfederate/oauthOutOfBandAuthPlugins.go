package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthOutOfBandAuthPluginsService service

//GetOOBAuthPluginDescriptors - Get the list of available Out of Band authenticator plugin descriptors.
//RequestType: GET
//Input:
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthPluginDescriptors() (result *OutOfBandAuthPluginDescriptors, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/descriptors"
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

//GetOOBAuthPluginDescriptor - Get the descriptor of an Out of Band authenticator plugin.
//RequestType: GET
//Input: input *GetOOBAuthPluginDescriptorInput
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthPluginDescriptor(input *GetOOBAuthPluginDescriptorInput) (result *OutOfBandAuthPluginDescriptor, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/descriptors/{id}"
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

//GetOOBAuthenticators - Get a list of Out of Band authenticator plugin instances.
//RequestType: GET
//Input:
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthenticators() (result *OutOfBandAuthenticators, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins"
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

//CreateOOBAuthenticator - Create an Out of Band authenticator plugin instance.
//RequestType: POST
//Input: input *CreateOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) CreateOOBAuthenticator(input *CreateOOBAuthenticatorInput) (result *OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins"
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

//GetOOBAuthenticator - Get a specific Out of Band authenticator plugin instance.
//RequestType: GET
//Input: input *GetOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) GetOOBAuthenticator(input *GetOOBAuthenticatorInput) (result *OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
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

//UpdateOOBAuthenticator - Update an Out of Band authenticator plugin instance.
//RequestType: PUT
//Input: input *UpdateOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) UpdateOOBAuthenticator(input *UpdateOOBAuthenticatorInput) (result *OutOfBandAuthenticator, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
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

//DeleteOOBAuthenticator - Delete an Out of Band authenticator plugin instance.
//RequestType: DELETE
//Input: input *DeleteOOBAuthenticatorInput
func (s *OauthOutOfBandAuthPluginsService) DeleteOOBAuthenticator(input *DeleteOOBAuthenticatorInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}"
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

//GetActions - List of actions for an Out of Band authenticator plugin instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *OauthOutOfBandAuthPluginsService) GetActions(input *GetActionsInput) (result *Actions, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions"
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

//GetAction - Find an Out of Band authenticator plugin instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *OauthOutOfBandAuthPluginsService) GetAction(input *GetActionInput) (result *Action, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions/{actionId}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

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

//InvokeAction - Invokes an action for Out of Band authenticator plugin instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *OauthOutOfBandAuthPluginsService) InvokeAction(input *InvokeActionInput) (result *ActionResult, resp *http.Response, err error) {
	path := "/oauth/outOfBandAuthPlugins/{id}/actions/{actionId}/invokeAction"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
