package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AuthenticationPolicyContractsService service

//GetAuthenticationPolicyContracts - Gets the Authentication Policy Contracts.
//RequestType: GET
//Input: input *GetAuthenticationPolicyContractsInput
func (s *AuthenticationPolicyContractsService) GetAuthenticationPolicyContracts(input *GetAuthenticationPolicyContractsInput) (result *AuthenticationPolicyContracts, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts"
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

//CreateAuthenticationPolicyContract - Create a new Authentication Policy Contract.
//RequestType: POST
//Input: input *CreateAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) CreateAuthenticationPolicyContract(input *CreateAuthenticationPolicyContractInput) (result *AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts"
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

//GetAuthenticationPolicyContract - Gets the Authentication Policy Contract by ID.
//RequestType: GET
//Input: input *GetAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) GetAuthenticationPolicyContract(input *GetAuthenticationPolicyContractInput) (result *AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
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

//UpdateAuthenticationPolicyContract - Update an Authentication Policy Contract by ID.
//RequestType: PUT
//Input: input *UpdateAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) UpdateAuthenticationPolicyContract(input *UpdateAuthenticationPolicyContractInput) (result *AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
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

//DeleteAuthenticationPolicyContract - Delete an Authentication Policy Contract.
//RequestType: DELETE
//Input: input *DeleteAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) DeleteAuthenticationPolicyContract(input *DeleteAuthenticationPolicyContractInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
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
