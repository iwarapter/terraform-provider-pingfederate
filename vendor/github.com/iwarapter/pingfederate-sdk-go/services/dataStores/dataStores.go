package dataStores

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type DataStoresService struct {
	Client *client.PfClient
}

// New creates a new instance of the DataStoresService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *DataStoresService {

	return &DataStoresService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetCustomDataStoreDescriptors - Get the list of available custom data store descriptors.
//RequestType: GET
//Input:
func (s *DataStoresService) GetCustomDataStoreDescriptors() (result *models.CustomDataStoreDescriptors, resp *http.Response, err error) {
	path := "/dataStores/descriptors"
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

//GetCustomDataStoreDescriptor - Get the description of a custom data store plugin by ID.
//RequestType: GET
//Input: input *GetCustomDataStoreDescriptorInput
func (s *DataStoresService) GetCustomDataStoreDescriptor(input *GetCustomDataStoreDescriptorInput) (result *models.CustomDataStoreDescriptor, resp *http.Response, err error) {
	path := "/dataStores/descriptors/{id}"
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

//GetDataStores - Get list of data stores.
//RequestType: GET
//Input:
func (s *DataStoresService) GetDataStores() (result *models.DataStores, resp *http.Response, err error) {
	path := "/dataStores"
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

//CreateJdbcDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateJdbcDataStoreInput
func (s *DataStoresService) CreateJdbcDataStore(input *CreateJdbcDataStoreInput) (result *models.JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores"
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

//CreateLdapDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateLdapDataStoreInput
func (s *DataStoresService) CreateLdapDataStore(input *CreateLdapDataStoreInput) (result *models.LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores"
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

//CreateCustomDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateCustomDataStoreInput
func (s *DataStoresService) CreateCustomDataStore(input *CreateCustomDataStoreInput) (result *models.CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores"
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

//GetDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetDataStoreInput
func (s *DataStoresService) GetDataStore(input *GetDataStoreInput) (result *models.DataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//DeleteDataStore - Delete a data store.
//RequestType: DELETE
//Input: input *DeleteDataStoreInput
func (s *DataStoresService) DeleteDataStore(input *DeleteDataStoreInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetJdbcDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetJdbcDataStoreInput
func (s *DataStoresService) GetJdbcDataStore(input *GetJdbcDataStoreInput) (result *models.JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetLdapDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetLdapDataStoreInput
func (s *DataStoresService) GetLdapDataStore(input *GetLdapDataStoreInput) (result *models.LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetCustomDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetCustomDataStoreInput
func (s *DataStoresService) GetCustomDataStore(input *GetCustomDataStoreInput) (result *models.CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//UpdateJdbcDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateJdbcDataStoreInput
func (s *DataStoresService) UpdateJdbcDataStore(input *UpdateJdbcDataStoreInput) (result *models.JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//UpdateLdapDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateLdapDataStoreInput
func (s *DataStoresService) UpdateLdapDataStore(input *UpdateLdapDataStoreInput) (result *models.LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//UpdateCustomDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateCustomDataStoreInput
func (s *DataStoresService) UpdateCustomDataStore(input *UpdateCustomDataStoreInput) (result *models.CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetActions - List the actions for a data store instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *DataStoresService) GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions"
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

//GetAction - Find a data store instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *DataStoresService) GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions/{actionId}"
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

//InvokeAction - Invokes an action for a data source instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *DataStoresService) InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions/{actionId}/invokeAction"
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

type CreateCustomDataStoreInput struct {
	Body models.CustomDataStore

	BypassExternalValidation *bool
}

type CreateJdbcDataStoreInput struct {
	Body models.JdbcDataStore

	BypassExternalValidation *bool
}

type CreateLdapDataStoreInput struct {
	Body models.LdapDataStore

	BypassExternalValidation *bool
}

type DeleteDataStoreInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetCustomDataStoreInput struct {
	Id string
}

type GetCustomDataStoreDescriptorInput struct {
	Id string
}

type GetDataStoreInput struct {
	Id string
}

type GetJdbcDataStoreInput struct {
	Id string
}

type GetLdapDataStoreInput struct {
	Id string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateCustomDataStoreInput struct {
	Body models.CustomDataStore
	Id   string

	BypassExternalValidation *bool
}

type UpdateJdbcDataStoreInput struct {
	Body models.JdbcDataStore
	Id   string

	BypassExternalValidation *bool
}

type UpdateLdapDataStoreInput struct {
	Body models.LdapDataStore
	Id   string

	BypassExternalValidation *bool
}
