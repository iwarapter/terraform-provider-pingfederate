package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type IdpStsRequestParametersContractsService service

//GetStsRequestParamContracts - Get the list of STS Request Parameters Contracts.
//RequestType: GET
//Input:
func (s *IdpStsRequestParametersContractsService) GetStsRequestParamContracts() (result *StsRequestParametersContracts, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts"
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

//CreateStsRequestParamContract - Create a new STS Request Parameters Contract.
//RequestType: POST
//Input: input *CreateStsRequestParamContractInput
func (s *IdpStsRequestParametersContractsService) CreateStsRequestParamContract(input *CreateStsRequestParamContractInput) (result *StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts"
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

//GetStsRequestParamContractById - Get a STS Request Parameters Contract.
//RequestType: GET
//Input: input *GetStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) GetStsRequestParamContractById(input *GetStsRequestParamContractByIdInput) (result *StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
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

//UpdateStsRequestParamContractById - Update a STS Request Parameters Contract.
//RequestType: PUT
//Input: input *UpdateStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) UpdateStsRequestParamContractById(input *UpdateStsRequestParamContractByIdInput) (result *StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
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

//DeleteStsRequestParamContractById - Delete a STS Request Parameters Contract.
//RequestType: DELETE
//Input: input *DeleteStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) DeleteStsRequestParamContractById(input *DeleteStsRequestParamContractByIdInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
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
