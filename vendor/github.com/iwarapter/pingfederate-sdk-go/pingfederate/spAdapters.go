package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SpAdaptersService service

//GetSpAdapterDescriptors - Get the list of available SP adapter descriptors.
//RequestType: GET
//Input:
func (s *SpAdaptersService) GetSpAdapterDescriptors() (result *SpAdapterDescriptors, resp *http.Response, err error) {
	path := "/sp/adapters/descriptors"
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

//GetSpAdapterDescriptorsById - Get the description of an SP adapter plugin by ID.
//RequestType: GET
//Input: input *GetSpAdapterDescriptorsByIdInput
func (s *SpAdaptersService) GetSpAdapterDescriptorsById(input *GetSpAdapterDescriptorsByIdInput) (result *SpAdapterDescriptor, resp *http.Response, err error) {
	path := "/sp/adapters/descriptors/{id}"
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

//GetSpAdapters - Get the list of configured SP adapter instances.
//RequestType: GET
//Input: input *GetSpAdaptersInput
func (s *SpAdaptersService) GetSpAdapters(input *GetSpAdaptersInput) (result *SpAdapters, resp *http.Response, err error) {
	path := "/sp/adapters"
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

//CreateSpAdapter - Create a new SP adapter instance.
//RequestType: POST
//Input: input *CreateSpAdapterInput
func (s *SpAdaptersService) CreateSpAdapter(input *CreateSpAdapterInput) (result *SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters"
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

//GetSpAdapter - Find an SP adapter instance by ID.
//RequestType: GET
//Input: input *GetSpAdapterInput
func (s *SpAdaptersService) GetSpAdapter(input *GetSpAdapterInput) (result *SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
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

//UpdateSpAdapter - Update an SP adapter instance.
//RequestType: PUT
//Input: input *UpdateSpAdapterInput
func (s *SpAdaptersService) UpdateSpAdapter(input *UpdateSpAdapterInput) (result *SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
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

//DeleteSpAdapter - Delete an SP adapter instance.
//RequestType: DELETE
//Input: input *DeleteSpAdapterInput
func (s *SpAdaptersService) DeleteSpAdapter(input *DeleteSpAdapterInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
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

//GetActions - List the actions for an SP adapter instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *SpAdaptersService) GetActions(input *GetActionsInput) (result *Actions, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions"
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

//GetAction - Find an SP adapter instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *SpAdaptersService) GetAction(input *GetActionInput) (result *Action, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions/{actionId}"
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

//InvokeAction - Invokes an action for an SP adapter instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *SpAdaptersService) InvokeAction(input *InvokeActionInput) (result *ActionResult, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions/{actionId}/invokeAction"
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

//GetUrlMappings - (Deprecated) List the mappings between URLs and adapter instances.
//RequestType: GET
//Input:
func (s *SpAdaptersService) GetUrlMappings() (result *SpAdapterUrlMappings, resp *http.Response, err error) {
	path := "/sp/adapters/urlMappings"
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

//UpdateUrlMappings - (Deprecated) Update the mappings between URLs and adapters instances.
//RequestType: PUT
//Input: input *UpdateUrlMappingsInput
func (s *SpAdaptersService) UpdateUrlMappings(input *UpdateUrlMappingsInput) (result *SpAdapterUrlMappings, resp *http.Response, err error) {
	path := "/sp/adapters/urlMappings"
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
