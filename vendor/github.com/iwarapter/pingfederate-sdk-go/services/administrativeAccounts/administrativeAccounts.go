package administrativeAccounts

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AdministrativeAccountsService struct {
	Client *client.PfClient
}

// New creates a new instance of the AdministrativeAccountsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *AdministrativeAccountsService {

	return &AdministrativeAccountsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetAccounts - Get all the PingFederate native Administrative Accounts.
//RequestType: GET
//Input:
func (s *AdministrativeAccountsService) GetAccounts() (result *models.AdministrativeAccounts, resp *http.Response, err error) {
	path := "/administrativeAccounts"
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

//AddAccount - Add a new PingFederate native Administrative Account.
//RequestType: POST
//Input: input *AddAccountInput
func (s *AdministrativeAccountsService) AddAccount(input *AddAccountInput) (result *models.AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts"
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

//GetAccount - Get a PingFederate native Administrative Account.
//RequestType: GET
//Input: input *GetAccountInput
func (s *AdministrativeAccountsService) GetAccount(input *GetAccountInput) (result *models.AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

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

//UpdateAccount - Update the information for a native Administrative Account.
//RequestType: PUT
//Input: input *UpdateAccountInput
func (s *AdministrativeAccountsService) UpdateAccount(input *UpdateAccountInput) (result *models.AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

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

//DeleteAccount - Delete a PingFederate native Administrative Account information.
//RequestType: DELETE
//Input: input *DeleteAccountInput
func (s *AdministrativeAccountsService) DeleteAccount(input *DeleteAccountInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

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

//ResetPassword - Reset the Password of an existing PingFederate native Administrative Account.
//RequestType: POST
//Input: input *ResetPasswordInput
func (s *AdministrativeAccountsService) ResetPassword(input *ResetPasswordInput) (result *models.UserCredentials, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}/resetPassword"
	path = strings.Replace(path, "{username}", input.Username, -1)

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

//ChangePassword - Change the Password of current PingFederate native Account.
//RequestType: POST
//Input: input *ChangePasswordInput
func (s *AdministrativeAccountsService) ChangePassword(input *ChangePasswordInput) (result *models.UserCredentials, resp *http.Response, err error) {
	path := "/administrativeAccounts/changePassword"
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

type AddAccountInput struct {
	Body models.AdministrativeAccount
}

type ChangePasswordInput struct {
	Body models.UserCredentials
}

type DeleteAccountInput struct {
	Username string
}

type GetAccountInput struct {
	Username string
}

type ResetPasswordInput struct {
	Body     models.UserCredentials
	Username string
}

type UpdateAccountInput struct {
	Body     models.AdministrativeAccount
	Username string
}
