package idpStsRequestParametersContracts

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpStsRequestParametersContractsService struct {
	Client *client.PfClient
}

// New creates a new instance of the IdpStsRequestParametersContractsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *IdpStsRequestParametersContractsService {

	return &IdpStsRequestParametersContractsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetStsRequestParamContracts - Get the list of STS Request Parameters Contracts.
//RequestType: GET
//Input:
func (s *IdpStsRequestParametersContractsService) GetStsRequestParamContracts() (result *models.StsRequestParametersContracts, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts"
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

//CreateStsRequestParamContract - Create a new STS Request Parameters Contract.
//RequestType: POST
//Input: input *CreateStsRequestParamContractInput
func (s *IdpStsRequestParametersContractsService) CreateStsRequestParamContract(input *CreateStsRequestParamContractInput) (result *models.StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts"
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

//GetStsRequestParamContractById - Get a STS Request Parameters Contract.
//RequestType: GET
//Input: input *GetStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) GetStsRequestParamContractById(input *GetStsRequestParamContractByIdInput) (result *models.StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
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

//UpdateStsRequestParamContractById - Update a STS Request Parameters Contract.
//RequestType: PUT
//Input: input *UpdateStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) UpdateStsRequestParamContractById(input *UpdateStsRequestParamContractByIdInput) (result *models.StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
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

//DeleteStsRequestParamContractById - Delete a STS Request Parameters Contract.
//RequestType: DELETE
//Input: input *DeleteStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) DeleteStsRequestParamContractById(input *DeleteStsRequestParamContractByIdInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
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

type CreateStsRequestParamContractInput struct {
	Body models.StsRequestParametersContract
}

type DeleteStsRequestParamContractByIdInput struct {
	Id string
}

type GetStsRequestParamContractByIdInput struct {
	Id string
}

type UpdateStsRequestParamContractByIdInput struct {
	Body models.StsRequestParametersContract
	Id   string
}
