package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type DataStoresService service

//GetCustomDataStoreDescriptors - Get the list of available custom data store descriptors.
//RequestType: GET
//Input:
func (s *DataStoresService) GetCustomDataStoreDescriptors() (result *CustomDataStoreDescriptors, resp *http.Response, err error) {
	path := "/dataStores/descriptors"
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

//GetCustomDataStoreDescriptor - Get the description of a custom data store plugin by ID.
//RequestType: GET
//Input: input *GetCustomDataStoreDescriptorInput
func (s *DataStoresService) GetCustomDataStoreDescriptor(input *GetCustomDataStoreDescriptorInput) (result *CustomDataStoreDescriptor, resp *http.Response, err error) {
	path := "/dataStores/descriptors/{id}"
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

//GetDataStores - Get list of data stores.
//RequestType: GET
//Input:
func (s *DataStoresService) GetDataStores() (result *DataStores, resp *http.Response, err error) {
	path := "/dataStores"
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

//CreateJdbcDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateJdbcDataStoreInput
func (s *DataStoresService) CreateJdbcDataStore(input *CreateJdbcDataStoreInput) (result *JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores"
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

//CreateLdapDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateLdapDataStoreInput
func (s *DataStoresService) CreateLdapDataStore(input *CreateLdapDataStoreInput) (result *LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores"
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

//CreateCustomDataStore - Create a new data store.
//RequestType: POST
//Input: input *CreateCustomDataStoreInput
func (s *DataStoresService) CreateCustomDataStore(input *CreateCustomDataStoreInput) (result *CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores"
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

//DeleteDataStore - Delete a data store.
//RequestType: DELETE
//Input: input *DeleteDataStoreInput
func (s *DataStoresService) DeleteDataStore(input *DeleteDataStoreInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetDataStoreInput
func (s *DataStoresService) GetDataStore(input *GetDataStoreInput) (result *DataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetJdbcDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetJdbcDataStoreInput
func (s *DataStoresService) GetJdbcDataStore(input *GetJdbcDataStoreInput) (result *JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetLdapDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetLdapDataStoreInput
func (s *DataStoresService) GetLdapDataStore(input *GetLdapDataStoreInput) (result *LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetCustomDataStore - Find data store by ID.
//RequestType: GET
//Input: input *GetCustomDataStoreInput
func (s *DataStoresService) GetCustomDataStore(input *GetCustomDataStoreInput) (result *CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//UpdateJdbcDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateJdbcDataStoreInput
func (s *DataStoresService) UpdateJdbcDataStore(input *UpdateJdbcDataStoreInput) (result *JdbcDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//UpdateLdapDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateLdapDataStoreInput
func (s *DataStoresService) UpdateLdapDataStore(input *UpdateLdapDataStoreInput) (result *LdapDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//UpdateCustomDataStore - Update a data store.
//RequestType: PUT
//Input: input *UpdateCustomDataStoreInput
func (s *DataStoresService) UpdateCustomDataStore(input *UpdateCustomDataStoreInput) (result *CustomDataStore, resp *http.Response, err error) {
	path := "/dataStores/{id}"
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

//GetActions - List the actions for a data store instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *DataStoresService) GetActions(input *GetActionsInput) (result *Actions, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions"
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

//GetAction - Find a data store instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *DataStoresService) GetAction(input *GetActionInput) (result *Action, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions/{actionId}"
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

//InvokeAction - Invokes an action for a data source instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *DataStoresService) InvokeAction(input *InvokeActionInput) (result *ActionResult, resp *http.Response, err error) {
	path := "/dataStores/{id}/actions/{actionId}/invokeAction"
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
