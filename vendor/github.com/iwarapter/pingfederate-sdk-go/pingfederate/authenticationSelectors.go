package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AuthenticationSelectorsService service

//GetAuthenticationSelectorDescriptors - Get the list of available Authentication Selector descriptors.
//RequestType: GET
//Input:
func (s *AuthenticationSelectorsService) GetAuthenticationSelectorDescriptors() (result *AuthenticationSelectorDescriptors, resp *http.Response, err error) {
	path := "/authenticationSelectors/descriptors"
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

//GetAuthenticationSelectorDescriptorsById - Get the description of an Authentication Selector plugin by ID.
//RequestType: GET
//Input: input *GetAuthenticationSelectorDescriptorsByIdInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelectorDescriptorsById(input *GetAuthenticationSelectorDescriptorsByIdInput) (result *AuthenticationSelectorDescriptor, resp *http.Response, err error) {
	path := "/authenticationSelectors/descriptors/{id}"
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

//GetAuthenticationSelectors - Get the list of configured Authentication Selector instances.
//RequestType: GET
//Input: input *GetAuthenticationSelectorsInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelectors(input *GetAuthenticationSelectorsInput) (result *AuthenticationSelectors, resp *http.Response, err error) {
	path := "/authenticationSelectors"
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

//CreateAuthenticationSelector - Create a new authentication selector instance.
//RequestType: POST
//Input: input *CreateAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) CreateAuthenticationSelector(input *CreateAuthenticationSelectorInput) (result *AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors"
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

//GetAuthenticationSelector - Get an Authentication Selector instance by ID.
//RequestType: GET
//Input: input *GetAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelector(input *GetAuthenticationSelectorInput) (result *AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
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

//UpdateAuthenticationSelector - Update an authentication selector instance.
//RequestType: PUT
//Input: input *UpdateAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) UpdateAuthenticationSelector(input *UpdateAuthenticationSelectorInput) (result *AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
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

//DeleteAuthenticationSelector - Delete an Authentication Selector instance.
//RequestType: DELETE
//Input: input *DeleteAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) DeleteAuthenticationSelector(input *DeleteAuthenticationSelectorInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
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
