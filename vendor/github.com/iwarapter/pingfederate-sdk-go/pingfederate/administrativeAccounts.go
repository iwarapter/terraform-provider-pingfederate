package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AdministrativeAccountsService service

//GetAccounts - Get all the PingFederate native Administrative Accounts.
//RequestType: GET
//Input:
func (s *AdministrativeAccountsService) GetAccounts() (result *AdministrativeAccounts, resp *http.Response, err error) {
	path := "/administrativeAccounts"
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

//AddAccount - Add a new PingFederate native Administrative Account.
//RequestType: POST
//Input: input *AddAccountInput
func (s *AdministrativeAccountsService) AddAccount(input *AddAccountInput) (result *AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts"
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

//GetAccount - Get a PingFederate native Administrative Account.
//RequestType: GET
//Input: input *GetAccountInput
func (s *AdministrativeAccountsService) GetAccount(input *GetAccountInput) (result *AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

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

//UpdateAccount - Update the information for a native Administrative Account.
//RequestType: PUT
//Input: input *UpdateAccountInput
func (s *AdministrativeAccountsService) UpdateAccount(input *UpdateAccountInput) (result *AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

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

//DeleteAccount - Delete a PingFederate native Administrative Account information.
//RequestType: DELETE
//Input: input *DeleteAccountInput
func (s *AdministrativeAccountsService) DeleteAccount(input *DeleteAccountInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

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

//ResetPassword - Reset the Password of an existing PingFederate native Administrative Account.
//RequestType: POST
//Input: input *ResetPasswordInput
func (s *AdministrativeAccountsService) ResetPassword(input *ResetPasswordInput) (result *UserCredentials, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}/resetPassword"
	path = strings.Replace(path, "{username}", input.Username, -1)

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

//ChangePassword - Change the Password of current PingFederate native Account.
//RequestType: POST
//Input: input *ChangePasswordInput
func (s *AdministrativeAccountsService) ChangePassword(input *ChangePasswordInput) (result *UserCredentials, resp *http.Response, err error) {
	path := "/administrativeAccounts/changePassword"
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
