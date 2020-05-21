package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type IdpAdaptersService service

//GetIdpAdapterDescriptors - Get the list of available IdP adapter descriptors.
//RequestType: GET
//Input:
func (s *IdpAdaptersService) GetIdpAdapterDescriptors() (result *IdpAdapterDescriptors, resp *http.Response, err error) {
	path := "/idp/adapters/descriptors"
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

//GetIdpAdapterDescriptorsById - Get the description of an IdP adapter plugin by ID.
//RequestType: GET
//Input: input *GetIdpAdapterDescriptorsByIdInput
func (s *IdpAdaptersService) GetIdpAdapterDescriptorsById(input *GetIdpAdapterDescriptorsByIdInput) (result *IdpAdapterDescriptor, resp *http.Response, err error) {
	path := "/idp/adapters/descriptors/{id}"
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

//GetIdpAdapters - Get the list of configured IdP adapter instances.
//RequestType: GET
//Input: input *GetIdpAdaptersInput
func (s *IdpAdaptersService) GetIdpAdapters(input *GetIdpAdaptersInput) (result *IdpAdapters, resp *http.Response, err error) {
	path := "/idp/adapters"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.Filter != "" {
		q.Set("filter", input.Filter)
	}
	rel.RawQuery = q.Encode()
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

//CreateIdpAdapter - Create a new IdP adapter instance.
//RequestType: POST
//Input: input *CreateIdpAdapterInput
func (s *IdpAdaptersService) CreateIdpAdapter(input *CreateIdpAdapterInput) (result *IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetIdpAdapter - Find an IdP adapter instance by ID.
//RequestType: GET
//Input: input *GetIdpAdapterInput
func (s *IdpAdaptersService) GetIdpAdapter(input *GetIdpAdapterInput) (result *IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
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

//UpdateIdpAdapter - Update an IdP adapter instance.
//RequestType: PUT
//Input: input *UpdateIdpAdapterInput
func (s *IdpAdaptersService) UpdateIdpAdapter(input *UpdateIdpAdapterInput) (result *IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteIdpAdapter - Delete an IdP adapter instance.
//RequestType: DELETE
//Input: input *DeleteIdpAdapterInput
func (s *IdpAdaptersService) DeleteIdpAdapter(input *DeleteIdpAdapterInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
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

//GetActions - List the actions for an IdP adapter instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *IdpAdaptersService) GetActions(input *GetActionsInput) (result *Actions, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions"
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

//GetAction - Find an IdP adapter instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *IdpAdaptersService) GetAction(input *GetActionInput) (result *Action, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions/{actionId}"
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

//InvokeAction - Invokes an action for an IdP adapter instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *IdpAdaptersService) InvokeAction(input *InvokeActionInput) (result *ActionResult, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions/{actionId}/invokeAction"
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
