package authenticationSelectors

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationSelectorsService struct {
	Client *client.PfClient
}

// New creates a new instance of the AuthenticationSelectorsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *AuthenticationSelectorsService {

	return &AuthenticationSelectorsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetAuthenticationSelectorDescriptors - Get the list of available Authentication Selector descriptors.
//RequestType: GET
//Input:
func (s *AuthenticationSelectorsService) GetAuthenticationSelectorDescriptors() (result *models.AuthenticationSelectorDescriptors, resp *http.Response, err error) {
	path := "/authenticationSelectors/descriptors"
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

//GetAuthenticationSelectorDescriptorsById - Get the description of an Authentication Selector plugin by ID.
//RequestType: GET
//Input: input *GetAuthenticationSelectorDescriptorsByIdInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelectorDescriptorsById(input *GetAuthenticationSelectorDescriptorsByIdInput) (result *models.AuthenticationSelectorDescriptor, resp *http.Response, err error) {
	path := "/authenticationSelectors/descriptors/{id}"
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

//GetAuthenticationSelectors - Get the list of configured Authentication Selector instances.
//RequestType: GET
//Input: input *GetAuthenticationSelectorsInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelectors(input *GetAuthenticationSelectorsInput) (result *models.AuthenticationSelectors, resp *http.Response, err error) {
	path := "/authenticationSelectors"
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

//CreateAuthenticationSelector - Create a new authentication selector instance.
//RequestType: POST
//Input: input *CreateAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) CreateAuthenticationSelector(input *CreateAuthenticationSelectorInput) (result *models.AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors"
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

//GetAuthenticationSelector - Get an Authentication Selector instance by ID.
//RequestType: GET
//Input: input *GetAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) GetAuthenticationSelector(input *GetAuthenticationSelectorInput) (result *models.AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
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

//UpdateAuthenticationSelector - Update an authentication selector instance.
//RequestType: PUT
//Input: input *UpdateAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) UpdateAuthenticationSelector(input *UpdateAuthenticationSelectorInput) (result *models.AuthenticationSelector, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
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

//DeleteAuthenticationSelector - Delete an Authentication Selector instance.
//RequestType: DELETE
//Input: input *DeleteAuthenticationSelectorInput
func (s *AuthenticationSelectorsService) DeleteAuthenticationSelector(input *DeleteAuthenticationSelectorInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/authenticationSelectors/{id}"
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

type CreateAuthenticationSelectorInput struct {
	Body models.AuthenticationSelector
}

type DeleteAuthenticationSelectorInput struct {
	Id string
}

type GetAuthenticationSelectorInput struct {
	Id string
}

type GetAuthenticationSelectorDescriptorsByIdInput struct {
	Id string
}

type GetAuthenticationSelectorsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type UpdateAuthenticationSelectorInput struct {
	Body models.AuthenticationSelector
	Id   string
}
