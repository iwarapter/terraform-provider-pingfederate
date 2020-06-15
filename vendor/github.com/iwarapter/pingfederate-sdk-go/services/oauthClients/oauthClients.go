package oauthClients

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientsService struct {
	Client *client.PfClient
}

// New creates a new instance of the OauthClientsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *OauthClientsService {

	return &OauthClientsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetClients - Get the list of OAuth clients.
//RequestType: GET
//Input: input *GetClientsInput
func (s *OauthClientsService) GetClients(input *GetClientsInput) (result *models.Clients, resp *http.Response, err error) {
	path := "/oauth/clients"
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

//CreateClient - Create a new OAuth client.
//RequestType: POST
//Input: input *CreateClientInput
func (s *OauthClientsService) CreateClient(input *CreateClientInput) (result *models.Client, resp *http.Response, err error) {
	path := "/oauth/clients"
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

//GetClient - Find the OAuth client by ID.
//RequestType: GET
//Input: input *GetClientInput
func (s *OauthClientsService) GetClient(input *GetClientInput) (result *models.Client, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
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

//UpdateClient - Updates the OAuth client.
//RequestType: PUT
//Input: input *UpdateClientInput
func (s *OauthClientsService) UpdateClient(input *UpdateClientInput) (result *models.Client, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
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

//DeleteClient - Delete an OAuth client.
//RequestType: DELETE
//Input: input *DeleteClientInput
func (s *OauthClientsService) DeleteClient(input *DeleteClientInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/oauth/clients/{id}"
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

//GetClientSecret - Get the client secret of an existing OAuth client.
//RequestType: GET
//Input: input *GetClientSecretInput
func (s *OauthClientsService) GetClientSecret(input *GetClientSecretInput) (result *models.ClientSecret, resp *http.Response, err error) {
	path := "/oauth/clients/{id}/clientAuth/clientSecret"
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

//UpdateClientSecret - Update the client secret of an existing OAuth client.
//RequestType: PUT
//Input: input *UpdateClientSecretInput
func (s *OauthClientsService) UpdateClientSecret(input *UpdateClientSecretInput) (result *models.ClientSecret, resp *http.Response, err error) {
	path := "/oauth/clients/{id}/clientAuth/clientSecret"
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

type CreateClientInput struct {
	Body models.Client
}

type DeleteClientInput struct {
	Id string
}

type GetClientInput struct {
	Id string
}

type GetClientSecretInput struct {
	Id string
}

type GetClientsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type UpdateClientInput struct {
	Body models.Client
	Id   string
}

type UpdateClientSecretInput struct {
	Body models.ClientSecret
	Id   string
}
