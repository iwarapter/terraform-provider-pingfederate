package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OauthClientsService service

//GetClients - Get the list of OAuth clients.
//RequestType: GET
//Input: input *GetClientsInput
func (s *OauthClientsService) GetClients(input *GetClientsInput) (result *Clients, resp *http.Response, err error) {
	path := "/oauth/clients"
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

//CreateClient - Create a new OAuth client.
//RequestType: POST
//Input: input *CreateClientInput
func (s *OauthClientsService) CreateClient(input *CreateClientInput) (result *Client, resp *http.Response, err error) {
	path := "/oauth/clients"
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

//GetClient - Find the OAuth client by ID.
//RequestType: GET
//Input: input *GetClientInput
func (s *OauthClientsService) GetClient(input *GetClientInput) (result *Client, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
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

//UpdateClient - Updates the OAuth client.
//RequestType: PUT
//Input: input *UpdateClientInput
func (s *OauthClientsService) UpdateClient(input *UpdateClientInput) (result *Client, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
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

//DeleteClient - Delete an OAuth client.
//RequestType: DELETE
//Input: input *DeleteClientInput
func (s *OauthClientsService) DeleteClient(input *DeleteClientInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
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

//GetClientSecret - Get the client secret of an existing OAuth client.
//RequestType: GET
//Input: input *GetClientSecretInput
func (s *OauthClientsService) GetClientSecret(input *GetClientSecretInput) (result *ClientSecret, resp *http.Response, err error) {
	path := "/oauth/clients/{id}/clientAuth/clientSecret"
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

//UpdateClientSecret - Update the client secret of an existing OAuth client.
//RequestType: PUT
//Input: input *UpdateClientSecretInput
func (s *OauthClientsService) UpdateClientSecret(input *UpdateClientSecretInput) (result *ClientSecret, resp *http.Response, err error) {
	path := "/oauth/clients/{id}/clientAuth/clientSecret"
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
