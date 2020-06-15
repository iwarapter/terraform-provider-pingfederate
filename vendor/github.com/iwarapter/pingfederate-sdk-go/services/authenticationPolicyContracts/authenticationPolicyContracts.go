package authenticationPolicyContracts

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationPolicyContractsService struct {
	Client *client.PfClient
}

// New creates a new instance of the AuthenticationPolicyContractsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *AuthenticationPolicyContractsService {

	return &AuthenticationPolicyContractsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetAuthenticationPolicyContracts - Gets the Authentication Policy Contracts.
//RequestType: GET
//Input: input *GetAuthenticationPolicyContractsInput
func (s *AuthenticationPolicyContractsService) GetAuthenticationPolicyContracts(input *GetAuthenticationPolicyContractsInput) (result *models.AuthenticationPolicyContracts, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts"
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

//CreateAuthenticationPolicyContract - Create a new Authentication Policy Contract.
//RequestType: POST
//Input: input *CreateAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) CreateAuthenticationPolicyContract(input *CreateAuthenticationPolicyContractInput) (result *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts"
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

//GetAuthenticationPolicyContract - Gets the Authentication Policy Contract by ID.
//RequestType: GET
//Input: input *GetAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) GetAuthenticationPolicyContract(input *GetAuthenticationPolicyContractInput) (result *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
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

//UpdateAuthenticationPolicyContract - Update an Authentication Policy Contract by ID.
//RequestType: PUT
//Input: input *UpdateAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) UpdateAuthenticationPolicyContract(input *UpdateAuthenticationPolicyContractInput) (result *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
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

//DeleteAuthenticationPolicyContract - Delete an Authentication Policy Contract.
//RequestType: DELETE
//Input: input *DeleteAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) DeleteAuthenticationPolicyContract(input *DeleteAuthenticationPolicyContractInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
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

type CreateAuthenticationPolicyContractInput struct {
	Body models.AuthenticationPolicyContract
}

type DeleteAuthenticationPolicyContractInput struct {
	Id string
}

type GetAuthenticationPolicyContractInput struct {
	Id string
}

type GetAuthenticationPolicyContractsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type UpdateAuthenticationPolicyContractInput struct {
	Body models.AuthenticationPolicyContract
	Id   string
}
