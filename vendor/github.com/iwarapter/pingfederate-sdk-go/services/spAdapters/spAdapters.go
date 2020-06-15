package spAdapters

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpAdaptersService struct {
	Client *client.PfClient
}

// New creates a new instance of the SpAdaptersService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *SpAdaptersService {

	return &SpAdaptersService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSpAdapterDescriptors - Get the list of available SP adapter descriptors.
//RequestType: GET
//Input:
func (s *SpAdaptersService) GetSpAdapterDescriptors() (result *models.SpAdapterDescriptors, resp *http.Response, err error) {
	path := "/sp/adapters/descriptors"
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

//GetSpAdapterDescriptorsById - Get the description of an SP adapter plugin by ID.
//RequestType: GET
//Input: input *GetSpAdapterDescriptorsByIdInput
func (s *SpAdaptersService) GetSpAdapterDescriptorsById(input *GetSpAdapterDescriptorsByIdInput) (result *models.SpAdapterDescriptor, resp *http.Response, err error) {
	path := "/sp/adapters/descriptors/{id}"
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

//GetSpAdapters - Get the list of configured SP adapter instances.
//RequestType: GET
//Input: input *GetSpAdaptersInput
func (s *SpAdaptersService) GetSpAdapters(input *GetSpAdaptersInput) (result *models.SpAdapters, resp *http.Response, err error) {
	path := "/sp/adapters"
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

//CreateSpAdapter - Create a new SP adapter instance.
//RequestType: POST
//Input: input *CreateSpAdapterInput
func (s *SpAdaptersService) CreateSpAdapter(input *CreateSpAdapterInput) (result *models.SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters"
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

//GetSpAdapter - Find an SP adapter instance by ID.
//RequestType: GET
//Input: input *GetSpAdapterInput
func (s *SpAdaptersService) GetSpAdapter(input *GetSpAdapterInput) (result *models.SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
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

//UpdateSpAdapter - Update an SP adapter instance.
//RequestType: PUT
//Input: input *UpdateSpAdapterInput
func (s *SpAdaptersService) UpdateSpAdapter(input *UpdateSpAdapterInput) (result *models.SpAdapter, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
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

//DeleteSpAdapter - Delete an SP adapter instance.
//RequestType: DELETE
//Input: input *DeleteSpAdapterInput
func (s *SpAdaptersService) DeleteSpAdapter(input *DeleteSpAdapterInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/sp/adapters/{id}"
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

//GetActions - List the actions for an SP adapter instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *SpAdaptersService) GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions"
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

//GetAction - Find an SP adapter instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *SpAdaptersService) GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions/{actionId}"
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

//InvokeAction - Invokes an action for an SP adapter instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *SpAdaptersService) InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error) {
	path := "/sp/adapters/{id}/actions/{actionId}/invokeAction"
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

//GetUrlMappings - (Deprecated) List the mappings between URLs and adapter instances.
//RequestType: GET
//Input:
func (s *SpAdaptersService) GetUrlMappings() (result *models.SpAdapterUrlMappings, resp *http.Response, err error) {
	path := "/sp/adapters/urlMappings"
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

//UpdateUrlMappings - (Deprecated) Update the mappings between URLs and adapters instances.
//RequestType: PUT
//Input: input *UpdateUrlMappingsInput
func (s *SpAdaptersService) UpdateUrlMappings(input *UpdateUrlMappingsInput) (result *models.SpAdapterUrlMappings, resp *http.Response, err error) {
	path := "/sp/adapters/urlMappings"
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

type CreateSpAdapterInput struct {
	Body models.SpAdapter
}

type DeleteSpAdapterInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetSpAdapterInput struct {
	Id string
}

type GetSpAdapterDescriptorsByIdInput struct {
	Id string
}

type GetSpAdaptersInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateSpAdapterInput struct {
	Body models.SpAdapter
	Id   string
}

type UpdateUrlMappingsInput struct {
	Body models.SpAdapterUrlMappings
}
