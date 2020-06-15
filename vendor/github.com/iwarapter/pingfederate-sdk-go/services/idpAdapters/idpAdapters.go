package idpAdapters

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpAdaptersService struct {
	Client *client.PfClient
}

// New creates a new instance of the IdpAdaptersService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *IdpAdaptersService {

	return &IdpAdaptersService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetIdpAdapterDescriptors - Get the list of available IdP adapter descriptors.
//RequestType: GET
//Input:
func (s *IdpAdaptersService) GetIdpAdapterDescriptors() (result *models.IdpAdapterDescriptors, resp *http.Response, err error) {
	path := "/idp/adapters/descriptors"
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

//GetIdpAdapterDescriptorsById - Get the description of an IdP adapter plugin by ID.
//RequestType: GET
//Input: input *GetIdpAdapterDescriptorsByIdInput
func (s *IdpAdaptersService) GetIdpAdapterDescriptorsById(input *GetIdpAdapterDescriptorsByIdInput) (result *models.IdpAdapterDescriptor, resp *http.Response, err error) {
	path := "/idp/adapters/descriptors/{id}"
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

//GetIdpAdapters - Get the list of configured IdP adapter instances.
//RequestType: GET
//Input: input *GetIdpAdaptersInput
func (s *IdpAdaptersService) GetIdpAdapters(input *GetIdpAdaptersInput) (result *models.IdpAdapters, resp *http.Response, err error) {
	path := "/idp/adapters"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
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

//CreateIdpAdapter - Create a new IdP adapter instance.
//RequestType: POST
//Input: input *CreateIdpAdapterInput
func (s *IdpAdaptersService) CreateIdpAdapter(input *CreateIdpAdapterInput) (result *models.IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters"
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

//GetIdpAdapter - Find an IdP adapter instance by ID.
//RequestType: GET
//Input: input *GetIdpAdapterInput
func (s *IdpAdaptersService) GetIdpAdapter(input *GetIdpAdapterInput) (result *models.IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
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

//UpdateIdpAdapter - Update an IdP adapter instance.
//RequestType: PUT
//Input: input *UpdateIdpAdapterInput
func (s *IdpAdaptersService) UpdateIdpAdapter(input *UpdateIdpAdapterInput) (result *models.IdpAdapter, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
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

//DeleteIdpAdapter - Delete an IdP adapter instance.
//RequestType: DELETE
//Input: input *DeleteIdpAdapterInput
func (s *IdpAdaptersService) DeleteIdpAdapter(input *DeleteIdpAdapterInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/idp/adapters/{id}"
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

//GetActions - List the actions for an IdP adapter instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *IdpAdaptersService) GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions"
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

//GetAction - Find an IdP adapter instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *IdpAdaptersService) GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions/{actionId}"
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

//InvokeAction - Invokes an action for an IdP adapter instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *IdpAdaptersService) InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error) {
	path := "/idp/adapters/{id}/actions/{actionId}/invokeAction"
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

type CreateIdpAdapterInput struct {
	Body models.IdpAdapter

	BypassExternalValidation *bool
}

type DeleteIdpAdapterInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetIdpAdapterInput struct {
	Id string
}

type GetIdpAdapterDescriptorsByIdInput struct {
	Id string
}

type GetIdpAdaptersInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateIdpAdapterInput struct {
	Body models.IdpAdapter
	Id   string

	BypassExternalValidation *bool
}
